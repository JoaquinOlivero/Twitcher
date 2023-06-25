'use client';
import { enablePreview } from "@/actions";
import { usePC } from "@/context/pcContext";
import { OutputResponse__Output } from "@/pb/service/OutputResponse";
import { forwardRef, useEffect } from "react";

type Props = {
    outputStatus: OutputResponse__Output | undefined
    addVideoElement: Function
    handleSoundMuting: Function
    handleVolume: (value: number) => void
    muted: boolean
    isLoaded: boolean
    volume: number
}

export type Ref = HTMLDivElement;


const Preview = forwardRef<Ref, Props>((props, vRef) => {
    const { newPc } = usePC();
    const { outputStatus, addVideoElement, muted, handleSoundMuting, isLoaded, volume, handleVolume } = props


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

    useEffect(() => {
        if (outputStatus && outputStatus.ready) {
            showPreview()
        }
    }, [])


    return (
        <div className="w-1/2 h-full mx-auto relative">
            <div className="bg-foreground w-full h-full z-0 rounded-b-xl"></div>

            {/* Controls */}
            <div className="absolute z-10 w-[98%] h-[98%] top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 text-white flex items-end">
                {isLoaded &&
                    <div className="flex items-center gap-x-1">
                        <button onClick={() => handleSoundMuting()} className="z-3">
                            {muted ?
                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-5 h-5">
                                    <path strokeLinecap="round" strokeLinejoin="round" d="M17.25 9.75L19.5 12m0 0l2.25 2.25M19.5 12l2.25-2.25M19.5 12l-2.25 2.25m-10.5-6l4.72-4.72a.75.75 0 011.28.531V19.94a.75.75 0 01-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.506-1.938-1.354A9.01 9.01 0 012.25 12c0-.83.112-1.633.322-2.395C2.806 8.757 3.63 8.25 4.51 8.25H6.75z" />
                                </svg>
                                :
                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-5 h-5">
                                    <path strokeLinecap="round" strokeLinejoin="round" d="M19.114 5.636a9 9 0 010 12.728M16.463 8.288a5.25 5.25 0 010 7.424M6.75 8.25l4.72-4.72a.75.75 0 011.28.53v15.88a.75.75 0 01-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.507-1.938-1.354A9.01 9.01 0 012.25 12c0-.83.112-1.633.322-2.396C2.806 8.756 3.63 8.25 4.51 8.25H6.75z" />
                                </svg>
                            }
                        </button>
                        <input id="default-range" type="range" value={volume} onChange={(e) => handleVolume(Number(e.currentTarget.value))}
                            className="w-full h-2.5 accent-slate-300 rounded-lg appearance-none cursor-pointer dark:bg-background"></input>
                    </div>
                }
            </div>

            <div className="absolute top-0 w-full h-full" ref={vRef}>
            </div>

        </div >
    )
})

// set display name
Preview.displayName = 'Preview';

export default Preview