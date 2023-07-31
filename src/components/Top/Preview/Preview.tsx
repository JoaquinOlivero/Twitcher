'use client';
import { enablePreview, getOverlays } from "@/actions";
import { usePC } from "@/context/pcContext";
import { StatusResponse__Output } from "@/pb/service/StatusResponse";
import { forwardRef, useEffect, useRef } from "react";
import * as fabric from 'fabric'; // v6
var FontFaceObserver = require('fontfaceobserver');

type Overlay = {
    id: string
    width: number
    height: number
    pointX: number
    pointY: number
    show: boolean
    coverId: string
    text: string
    fontFamily: string
    fontSize: number
    lineHeight: number
    textColor: string
}

type VideoElementSize = {
    width: number
    height: number
}

type Props = {
    status: StatusResponse__Output | undefined
    addVideoElement: Function
    videoElementSize: VideoElementSize | null
    handleSoundMuting: Function
    handleVolume: (value: number) => void
    muted: boolean
    volume: number
}

export type Ref = HTMLDivElement;

const Preview = forwardRef<Ref, Props>((props, vRef) => {
    const { newPc, Overlays, setOverlays, sendMsg, isPreviewLoaded, fabricRef } = usePC();
    const { status, addVideoElement, videoElementSize, muted, handleSoundMuting, volume, handleVolume } = props
    const canvasRef = useRef<HTMLCanvasElement>(null);

    const showPreview = async () => {
        const pc: RTCPeerConnection | null = await newPc()

        if (pc) {

            pc.oniceconnectionstatechange = e => console.log(pc.iceConnectionState)

            pc.ontrack = function (event) {
                addVideoElement(event)
            }

            pc.onicecandidate = async event => {
                if (event.candidate === null) {
                    const sdp = btoa(JSON.stringify(pc.localDescription))
                    const serverSdp = await enablePreview(sdp)

                    try {
                        pc.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(serverSdp))))
                    } catch (e) {
                        console.log(e)
                    }
                }
            }

            // Offer to receive 1 audio, and 1 video track
            pc.addTransceiver('video', {
                'direction': 'sendrecv'
            })
            pc.addTransceiver('audio', {
                'direction': 'sendrecv'
            })

            pc.createOffer().then(d => pc.setLocalDescription(d)).catch(err => console.log(err))
        }
    }

    const initOverlays = async () => {
        const data = await getOverlays()
        if (data && data.overlays) {
            setOverlays(data.overlays as Overlay[])
        }
    }


    const addTextbox = async (o: Overlay) => {
        var ffamily = o.fontFamily.substring(0, o.fontFamily.lastIndexOf(".")).toLowerCase()
        var myfont = new FontFaceObserver(ffamily)
        await myfont.load()

        if (isPreviewLoaded && videoElementSize) {
            const textbox = new fabric.Textbox(o.text, {
                id: o.id,
                width: (videoElementSize.width / 1280) * o.width,
                top: (o.pointY / 1280) * videoElementSize.width, // Stream width resolution = 1280px.
                left: (o.pointX / 720) * videoElementSize.height, // Stream height resolution = 720px.
                lockScalingFlip: true,
                lockRotation: true,
                fontSize: (videoElementSize.width / 1280) * o.fontSize,
                lineHeight: o.lineHeight,
                fill: o.textColor,
                editable: false
            })

            textbox.set("fontFamily", ffamily)

            textbox.setControlsVisibility({
                tl: false,
                tr: false,
                bl: false,
                br: false,
                mt: false,
                mb: false,
                mtr: false
            })

            textbox.on("modified", (options) => {
                const pointX = options.target.getCoords()[0].x
                const pointY = options.target.getCoords()[0].y
                const actualPointX = Math.round((pointX / videoElementSize.width) * 1280)
                const actualPointY = Math.round((pointY / videoElementSize.height) * 720)

                o.pointX = Math.round(actualPointX)
                o.pointY = Math.round(actualPointY)

                const scaledWidth = options.target.width * options.target.scaleX
                const scaledHeight = options.target.height * options.target.scaleY
                const actualWidth = (scaledWidth / videoElementSize.width) * 1280
                const actualHeight = (scaledHeight / videoElementSize.height) * 720

                o.width = actualWidth
                o.height = actualHeight

                const msg = {
                    "type": "overlay",
                    "object": o
                }

                sendMsg(JSON.stringify(msg))

            })

            fabricRef.current?.add(textbox)
        }
    }

    const addCoverImage = (o: Overlay) => {
        if (isPreviewLoaded && videoElementSize) {
            var imgObj = new Image();
            imgObj.src = `/api/covers/${o.coverId}`;
            imgObj.onload = () => {
                var image = new fabric.Image(imgObj);
                // Stream width resolution = 1280px
                image.scaleToWidth((videoElementSize.width / 1280) * o.width)
                // Stream height resolution = 720px.
                image.scaleToHeight((videoElementSize.height / 720) * o.height)
                image.set({
                    id: o.id,
                    top: (o.pointY / 1280) * videoElementSize.width, // Stream width resolution = 1280px.
                    left: (o.pointX / 720) * videoElementSize.height, // Stream height resolution = 720px.
                    lockScalingFlip: true,
                    lockRotation: true,
                    objectCaching: false,
                });

                image.setControlsVisibility({
                    mt: false,
                    mb: false,
                    mr: false,
                    ml: false,
                    mtr: false
                })

                image.on("modified", (options) => {
                    const pointX = options.target.getCoords()[0].x
                    const pointY = options.target.getCoords()[0].y
                    const actualPointX = Math.round((pointX / videoElementSize.width) * 1280)
                    const actualPointY = Math.round((pointY / videoElementSize.height) * 720)

                    o.pointX = Math.round(actualPointX)
                    o.pointY = Math.round(actualPointY)

                    const scaledWidth = options.target.width * options.target.scaleX
                    const scaledHeight = options.target.height * options.target.scaleY
                    const actualWidth = (scaledWidth / videoElementSize.width) * 1280
                    const actualHeight = (scaledHeight / videoElementSize.height) * 720

                    o.width = actualWidth
                    o.height = actualHeight

                    const msg = {
                        "type": "overlay",
                        "object": o
                    }

                    sendMsg(JSON.stringify(msg))

                })

                fabricRef.current!.add(image)
            }
        }
    };

    useEffect(() => {
        if (status && status.output) {
            showPreview()
        }
    }, [])

    useEffect(() => {
        if (isPreviewLoaded && videoElementSize) {
            const initFabric = () => {
                fabricRef.current = new fabric.Canvas(canvasRef.current!, {
                    width: videoElementSize.width,
                    height: videoElementSize.height,
                });

                fabricRef.current.selection = false; // disable group selection
            };

            initFabric();
        }

        return () => {
            if (isPreviewLoaded) fabricRef.current!.dispose();
        }
    }, [isPreviewLoaded])

    useEffect(() => {
        if (Overlays) {
            if (fabricRef.current && Overlays) {
                var i = 0
                while (i < Overlays.length) {
                    switch (Overlays[i].id) {
                        case "cover":
                            const img = fabricRef.current.getObjects().find(obj => (obj as any).id === "cover") as fabric.Image
                            if (img) {
                                fabricRef.current.remove(img);
                                addCoverImage(Overlays[i])
                            } else {
                                addCoverImage(Overlays[i])
                            }
                            break;
                        case "song_name":
                        case "song_author":
                        case "song_page":
                            const textbox = fabricRef.current.getObjects().find(obj => (obj as any).id === Overlays[i].id) as fabric.Textbox
                            if (textbox) {
                                textbox.set("text", Overlays[i].text)
                                fabricRef.current.requestRenderAll()
                            } else {
                                addTextbox(Overlays[i])
                            }
                            break;
                    }
                    i++
                }
            } else {
                initOverlays()
            }
        }
    }, [Overlays])


    return (
        <div className="w-1/2 h-full mx-auto relative">
            <div className="bg-foreground w-full h-full z-0 rounded-b-xl">
                <div className="flex justify-center items-center bg-foreground w-full h-full z-0 rounded-b-xl">
                    {!isPreviewLoaded && status && status.output &&
                        <div className="flex justify-center items-center gap-2">
                            <span className="font-semibold tracking-wider capitalize text-white">loading preview</span>
                            <svg className="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                            </svg>
                        </div>
                    }
                </div>
            </div>

            {/* Controls */}
            <div className="absolute bottom-1 left-0 z-10 w-[98%] left-1/2 transform -translate-x-1/2 text-white flex items-end">
                {isPreviewLoaded &&
                    <div className="flex items-center gap-x-1">
                        <button onClick={() => handleSoundMuting()} className="z-3">
                            {muted ?
                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor" className="w-4 h-4">
                                    <path strokeLinecap="round" strokeLinejoin="round" d="M17.25 9.75L19.5 12m0 0l2.25 2.25M19.5 12l2.25-2.25M19.5 12l-2.25 2.25m-10.5-6l4.72-4.72a.75.75 0 011.28.531V19.94a.75.75 0 01-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.506-1.938-1.354A9.01 9.01 0 012.25 12c0-.83.112-1.633.322-2.395C2.806 8.757 3.63 8.25 4.51 8.25H6.75z" />
                                </svg>
                                :
                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor" className="w-4 h-4">
                                    <path strokeLinecap="round" strokeLinejoin="round" d="M19.114 5.636a9 9 0 010 12.728M16.463 8.288a5.25 5.25 0 010 7.424M6.75 8.25l4.72-4.72a.75.75 0 011.28.53v15.88a.75.75 0 01-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.507-1.938-1.354A9.01 9.01 0 012.25 12c0-.83.112-1.633.322-2.396C2.806 8.756 3.63 8.25 4.51 8.25H6.75z" />
                                </svg>
                            }
                        </button>
                        <input id="default-range" type="range" value={volume} onChange={(e) => handleVolume(Number(e.currentTarget.value))}
                            className="w-full h-1 [&::-webkit-slider-thumb]:h-[5px] rounded-lg appearance-none cursor-pointer bg-white/80 accent-primary"></input>
                    </div>
                }
            </div>

            <div className="absolute top-0 w-full h-full" ref={vRef}>
                <div className="absolute top-0 z-10">
                    <canvas ref={canvasRef} />
                </div>
            </div>

        </div >
    )
})

// set display name
Preview.displayName = 'Preview';

export default Preview

const debounce = (callback: any, wait: any) => {
    let timeoutId: any = null;
    return (...args: any) => {
        window.clearTimeout(timeoutId);
        timeoutId = window.setTimeout(() => {
            callback.apply(null, args);
        }, wait);
    };
}