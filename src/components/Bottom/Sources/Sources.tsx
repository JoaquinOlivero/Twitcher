'use client';

import { deleteBgVideo, getBgVideos, swapBgVideo, uploadVideoFile } from "@/actions";
import Modal from "@/components/Modal/Modal";
import ActionButton from "@/components/Utils/ActionButton";
import { debounce } from "@/components/Utils/debounce";
import { delay } from "@/components/Utils/delay";
import { usePC } from "@/context/pcContext";
import { BackgroundVideo__Output } from "@/pb/service/BackgroundVideo";
import { BackgroundVideosResponse__Output } from "@/pb/service/BackgroundVideosResponse";
import { BaseFabricObject } from "fabric/*";
import { useEffect, useMemo, useRef, useState } from "react";

type Overlay = {
    id: string
    type: string
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
    textAlign: string
}


type Props = {
    bgVideos: BackgroundVideosResponse__Output['videos'] | undefined
}

const Sources = ({ bgVideos }: Props) => {
    const { Overlays } = usePC();

    const handleMenuClick = (element: EventTarget & HTMLDivElement) => {
        const chevron = element.lastChild as HTMLDivElement
        const content = element.nextSibling as HTMLDivElement

        if (content.classList.contains("hidden")) {
            // change styling of sub menu header.
            element.classList.add("bg-primary/20")
            element.classList.remove("hover:bg-background")

            // change chevron
            chevron.innerHTML = `
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 15.75l7.5-7.5 7.5 7.5" />
            </svg>
          `
            // show the content of the sub menu by removing the "hidden" className.
            content.classList.remove("hidden")
        } else {
            element.classList.remove("bg-primary/20")
            element.classList.add("hover:bg-background")

            chevron.innerHTML = `
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
            </svg>
          `
            content.classList.add("hidden")
        }
    }

    return (
        <div className='w-1/4 h-full bg-foreground rounded-t-xl flex flex-col gap-2 overflow-hidden'>
            <div className="text-[#fff] w-full relative flex flex-col">
                <div className="text-center opacity-85 text-xl font-semibold uppercase tracking-wider my-1">
                    Sources
                </div>
                <div className="w-[98%] h-1 mx-auto bg-primary"></div>
            </div>

            <div className="w-[98%] h-auto mx-auto select-none">
                <div className="text-white w-[98%] mx-auto">
                    <div className="flex justify-between cursor-pointer transition hover:bg-background" onClick={(e) => handleMenuClick(e.currentTarget)}>
                        <span className="capitalize font-semibold tracking-wider">Song overlay</span>
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
                        </svg>
                    </div>

                    <div className="transition hidden bg-background p-1 font-semibold text-sm tracking-wider">
                        {Overlays &&
                            Overlays.map((o) => {
                                return <OverlayObject object={o} key={o.id} />
                            })
                        }
                    </div>
                </div>

                <div className="text-white w-[98%] mx-auto">
                    <div className="flex justify-between cursor-pointer transition hover:bg-background" onClick={(e) => handleMenuClick(e.currentTarget)}>
                        <span className="capitalize font-semibold tracking-wider">Background Video</span>
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
                        </svg>
                    </div>
                    <div className="transition hidden bg-background p-1 font-semibold text-sm tracking-wider">
                        <BackgroundVideos bgVideos={bgVideos} />
                    </div>
                </div>
            </div>

        </div>
    )
}

export default Sources

const BackgroundVideos = ({ bgVideos }: Props) => {
    const [file, setFile] = useState<File | null>(null)
    const [videos, setVideos] = useState<BackgroundVideo__Output[] | undefined>(bgVideos)
    const [isUploading, setIsUploading] = useState<boolean>(false)
    const [uploadStatus, setUploadStatus] = useState<boolean | null>(null)
    const [errorMessage, setErrorMessage] = useState<string | null>(null)
    const [deleteErrorMessage, setDeleteErrorMessage] = useState<string | null>(null)
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const ref = useRef<HTMLInputElement>(null)

    const handleFileSubmit = async () => {
        if (file) {
            const formData = new FormData();
            formData.append('file', file);

            setIsUploading(true)
            const res = await uploadVideoFile(formData)
            setFile(null)
            setIsUploading(false)

            setUploadStatus(res.ok)
            setErrorMessage(res.msg)

            const bgVideos = await getBgVideos()

            if (bgVideos) {
                setVideos(bgVideos?.videos)
            }

            await delay(2500)
            setUploadStatus(null)
            if (ref.current) {
                ref.current.value = ""
            }

        }
    }

    const handleSelectBgVideo = async (video: BackgroundVideo__Output) => {
        await swapBgVideo(video)
        const bgVideos = await getBgVideos()

        if (bgVideos) {
            setVideos(bgVideos?.videos)
        }
    }

    const handleDeleteBgVideo = async (video: BackgroundVideo__Output) => {
        const res = await deleteBgVideo(video)
        if (res.ok) {
            const bgVideos = await getBgVideos()

            if (bgVideos) {
                setVideos(bgVideos?.videos)
            }
        } else {
            setDeleteErrorMessage(res.msg)
            await delay(2500)
            setDeleteErrorMessage(null)
        }
    }

    useEffect(() => {
    }, [videos])


    return (
        <div className="w-full flex flex-col gap-1">
            <div className="w-full flex items-center justify-end uppercase">
                <div className="w-auto flex gap-1 cursor-pointer text-xs opacity-70 transition hover:opacity-100 hover:text-primary" onClick={() => setIsOpen(true)}>
                    <span>Upload File</span>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-4 h-4">
                        <path strokeLinecap="round" strokeLinejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m3.75 9v6m3-3H9m1.5-12H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z" />
                    </svg>
                </div>
            </div>

            <Modal tWidth="w-1/4" tHeight="h-1/4" isOpen={isOpen} setIsOpen={setIsOpen}>
                <div className="relative w-full h-full flex items-center justify-center">
                    <div className="flex flex-col gap-2 text-base w-full">
                        <label className="text-md text-center" >Upload file</label>
                        <input type="file" ref={ref} onChange={(e) => setFile(e.target.files && e.target.files[0])}
                            className="bg-background w-full text-sm rounded cursor-pointer focus:outline-none dark:placeholder-gray-400" />
                    </div>
                    {file &&
                        <div className="absolute right-0 w-1/4 self-end">
                            <ActionButton
                                text="start upload"
                                waitingText="uploading"
                                width="full"
                                disabled={isUploading}
                                isWaiting={isUploading}
                                backgroundColor="bg-primary"
                                backgroundColorHover="bg-primary/70"
                                onClick={() => handleFileSubmit()}
                            />
                        </div>
                    }
                    {uploadStatus !== null &&
                        <div className="absolute left-0 w-full text-center self-end uppercase font-bold">
                            {uploadStatus && <span className="text-lime-400">successful upload</span>}
                            {!uploadStatus &&
                                <span className="text-red-600">upload failed. {errorMessage}</span>
                            }
                        </div>
                    }
                </div>
            </Modal>

            {videos && videos.map((video) => {
                return <div className="w-full flex items-center gap-1" key={video.id}>
                    <div className="w-3/4 truncate">{video.name}</div>
                    <div className="w-1/4 flex items-center justify-end">

                        {video.active ?
                            <span className="w-3/4 bg-primary border-2 border-primary rounded font-bold text-center transition opacity-80 pointer-events-none">selected</span>
                            :
                            <div className="flex w-full justify-end items-center gap-1">
                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" onClick={() => handleDeleteBgVideo(video)} className="w-[19px] h-[19px] cursor-pointer transition hover:text-red-600">
                                    <path strokeLinecap="round" strokeLinejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" />
                                </svg>
                                <span className="w-3/4 border-2 rounded cursor-pointer opacity-70 font-bold text-center transition hover:opacity-100 hover:border-primary" onClick={() => handleSelectBgVideo(video)}>select</span>
                            </div>
                        }
                    </div>

                    {deleteErrorMessage && <span className="absolute w-3/4 top-0 text-red-500 truncate hover:h-auto">{deleteErrorMessage}</span>}
                </div>
            })}

        </div>
    )
}


type OverlayObjectProps = {
    object: Overlay
}

const OverlayObject = ({ object }: OverlayObjectProps) => {
    const [settings, setSettings] = useState<boolean>(false)
    const [pointX, setPointX] = useState<number | string>(object.pointX)
    const [pointY, setPointY] = useState<number | string>(object.pointY)
    const [width, setWidth] = useState<number | string>(object.width)
    const [fontSize, setFontSize] = useState<number | string>(object.fontSize)
    const [textColor, setTextColor] = useState<string>(object.textColor)
    const [textAlign, setTextAlign] = useState<string>(object.textAlign)
    const [showObject, setShowObject] = useState<boolean>(object.show)
    const settingsCogRef = useRef<SVGSVGElement | null>(null)
    const settingsMenuRef = useRef<HTMLDivElement | null>(null)
    const { fabricRef, videoElementSize, sendMsg } = usePC()

    const handleShowClick = (id: string) => {
        setShowObject(!showObject)
        if (fabricRef.current && videoElementSize) {
            const obj = fabricRef.current.getObjects().find(obj => (obj as any).id === id)
            if (obj) {
                obj.visible = !showObject
                fabricRef.current.renderAll()

                const pointX = obj.getCoords()[0].x
                const pointY = obj.getCoords()[0].y
                const actualPointX = Math.round((pointX / videoElementSize.width) * 1280)
                const actualPointY = Math.round((pointY / videoElementSize.height) * 720)

                object.pointX = Math.round(actualPointX)
                object.pointY = Math.round(actualPointY)

                const scaledWidth = obj.width * obj.scaleX
                const scaledHeight = obj.height * obj.scaleY
                const actualWidth = (scaledWidth / videoElementSize.width) * 1280
                const actualHeight = (scaledHeight / videoElementSize.height) * 720

                object.width = actualWidth
                object.height = actualHeight
                object.show = !showObject

                const msg = {
                    "type": "overlay",
                    "object": object
                }

                sendMsg(JSON.stringify(msg))
            }
        }
    }

    const handlePointX = (value: number) => {
        if (!value) {
            setPointX("")
            return
        }

        if (fabricRef.current) {
            setPointX(value)
            const obj = fabricRef.current.getObjects().find(obj => (obj as any).id === object.id)
            if (obj && videoElementSize) {
                obj.left = (value / 720) * videoElementSize.height
                obj.setCoords()
                fabricRef.current.renderAll()
                debouncedTrigger(obj, "modified")
            }
        }
    }

    const handlePointY = (value: number) => {
        if (!value) {
            setPointY("")
            return
        }

        if (fabricRef.current) {
            setPointY(value)
            const obj = fabricRef.current.getObjects().find(obj => (obj as any).id === object.id)
            if (obj && videoElementSize) {
                obj.top = (value / 1280) * videoElementSize.width
                obj.setCoords()
                fabricRef.current.renderAll()
                debouncedTrigger(obj, "modified")
            }
        }
    }

    const handleWidth = (value: number) => {
        if (!value) {
            setWidth("")
            return
        }

        if (fabricRef.current) {
            setWidth(value)

            const obj = fabricRef.current.getObjects().find(obj => (obj as any).id === object.id)
            if (obj && videoElementSize) {
                if (object.type === "img") {
                    obj.scaleToWidth((videoElementSize.width / 1280) * value)
                    fabricRef.current.renderAll()
                    debouncedTrigger(obj, "modified")
                } else {
                    obj.set("width", (videoElementSize.width / 1280) * value)
                    fabricRef.current.renderAll()
                    debouncedTrigger(obj, "resizing")
                }

            }
        }
    }

    const handleTextAlign = (value: string) => {
        if (value === textAlign) {
            return
        }

        if (fabricRef.current) {
            setTextAlign(value)

            const obj = fabricRef.current.getObjects().find(obj => (obj as any).id === object.id)
            if (obj) {
                obj.set("textAlign", value)
                fabricRef.current.renderAll()
                debouncedTrigger(obj, "modified")
            }
        }
    }

    const handleTextColor = (value: string) => {
        if (value === textColor) {
            return
        }

        if (fabricRef.current) {
            const rgb = hexToRgb(value)
            setTextColor(rgb)

            const obj = fabricRef.current.getObjects().find(obj => (obj as any).id === object.id)
            if (obj) {
                obj.set("fill", `rgb(${rgb})`)
                fabricRef.current.renderAll()
                debouncedTrigger(obj, "modified")
            }
        }
    }


    const handleFontSize = (value: number) => {
        if (!value) {
            setFontSize("")
            return
        }

        if (fabricRef.current && videoElementSize) {
            setFontSize(value)

            const obj = fabricRef.current.getObjects().find(obj => (obj as any).id === object.id)
            if (obj) {
                obj.set("fontSize", (videoElementSize.width / 1280) * value)
                fabricRef.current.renderAll()
                debouncedTrigger(obj, "modified")
            }

        }
    }

    const debouncedTrigger = useMemo(() =>
        debounce((obj: BaseFabricObject, action: string) => {
            // @ts-ignore
            obj.fire(action, { "target": obj })
        }, 300),
        []
    )

    useEffect(() => {
        if (settingsCogRef.current) {
            document.addEventListener("click", (e) => {
                if (settingsCogRef.current && settingsMenuRef.current) {
                    if (!settingsCogRef.current.contains(e.target as HTMLElement) && !settingsMenuRef.current.contains(e.target as HTMLElement)) {
                        setSettings(false)
                    }
                }
            }, true)
        }
    }, [settingsCogRef, settingsMenuRef])

    return (
        <div className="flex justify-between items-center">
            <span className="capitalize">{object.id.replace("song_", "")}</span>

            <div className="flex gap-2">
                <div onClick={() => handleShowClick(object.id)} className="cursor-pointer">
                    {showObject ?
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-5 h-5">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
                            <path strokeLinecap="round" strokeLinejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                        </svg>
                        :
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-5 h-5">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M3.98 8.223A10.477 10.477 0 001.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.45 10.45 0 0112 4.5c4.756 0 8.773 3.162 10.065 7.498a10.523 10.523 0 01-4.293 5.774M6.228 6.228L3 3m3.228 3.228l3.65 3.65m7.894 7.894L21 21m-3.228-3.228l-3.65-3.65m0 0a3 3 0 10-4.243-4.243m4.242 4.242L9.88 9.88" />
                        </svg>
                    }
                </div>

                <div className="relative">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" onClick={() => setSettings(!settings)} ref={settingsCogRef} className="w-5 h-5 cursor-pointer">
                        <path strokeLinecap="round" strokeLinejoin="round" d="M9.594 3.94c.09-.542.56-.94 1.11-.94h2.593c.55 0 1.02.398 1.11.94l.213 1.281c.063.374.313.686.645.87.074.04.147.083.22.127.324.196.72.257 1.075.124l1.217-.456a1.125 1.125 0 011.37.49l1.296 2.247a1.125 1.125 0 01-.26 1.431l-1.003.827c-.293.24-.438.613-.431.992a6.759 6.759 0 010 .255c-.007.378.138.75.43.99l1.005.828c.424.35.534.954.26 1.43l-1.298 2.247a1.125 1.125 0 01-1.369.491l-1.217-.456c-.355-.133-.75-.072-1.076.124a6.57 6.57 0 01-.22.128c-.331.183-.581.495-.644.869l-.213 1.28c-.09.543-.56.941-1.11.941h-2.594c-.55 0-1.02-.398-1.11-.94l-.213-1.281c-.062-.374-.312-.686-.644-.87a6.52 6.52 0 01-.22-.127c-.325-.196-.72-.257-1.076-.124l-1.217.456a1.125 1.125 0 01-1.369-.49l-1.297-2.247a1.125 1.125 0 01.26-1.431l1.004-.827c.292-.24.437-.613.43-.992a6.932 6.932 0 010-.255c.007-.378-.138-.75-.43-.99l-1.004-.828a1.125 1.125 0 01-.26-1.43l1.297-2.247a1.125 1.125 0 011.37-.491l1.216.456c.356.133.751.072 1.076-.124.072-.044.146-.087.22-.128.332-.183.582-.495.644-.869l.214-1.281z" />
                        <path strokeLinecap="round" strokeLinejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                    {settings && fabricRef.current &&
                        <div ref={settingsMenuRef} className="absolute w-40 h-auto bg-background -right-2 px-2 pb-1 rounded-t-xl flex flex-col gap-1 capitalize z-10">
                            <div className="w-full flex justify-between items-center">
                                <span>point x</span>
                                <input type="number" className="w-[42%] h-4 text-black outline-none px-1" value={pointX} onChange={(e) => handlePointX(parseInt(e.currentTarget.value))} />
                            </div>
                            <div className="w-full flex justify-between items-center">
                                <span>point y</span>
                                <input type="number" className="w-[42%] h-4 text-black outline-none px-1" value={pointY} onChange={(e) => handlePointY(parseInt(e.currentTarget.value))} />
                            </div>
                            <div className="w-full flex justify-between items-center">
                                <span>width</span>
                                <input type="number" className="w-[42%] h-4 text-black outline-none px-1" value={width} onChange={(e) => handleWidth(parseInt(e.currentTarget.value))} />
                            </div>
                            {object.type !== "textbox" &&
                                <div className="w-full flex justify-between items-center">
                                    <span>height</span>
                                    <input type="number" className="w-[42%] h-4 text-black outline-none px-1" value={width} onChange={(e) => handleWidth(parseInt(e.currentTarget.value))} />
                                </div>
                            }
                            {object.type === "textbox" &&
                                <>
                                    {/* <div className="w-full flex justify-between items-center">
                                        <span>font family</span>
                                    </div> */}
                                    <div className="w-full flex justify-between items-center">
                                        <span>font size</span>
                                        <input type="number" className="w-[42%] h-4 text-black outline-none px-1" value={fontSize} onChange={(e) => handleFontSize(parseInt(e.currentTarget.value))} />
                                    </div>
                                    <div className="w-full flex justify-between items-center">
                                        <span>text color</span>
                                        <input type="color" className="w-[42%] h-4 text-black outline-none border-none px-1 cursor-pointer" value={stringRGBToHex(textColor)} onChange={(e) => handleTextColor(e.currentTarget.value)} />
                                    </div>
                                    <div className="w-full flex justify-between items-center">
                                        <span>text align</span>
                                        <div className="flex">
                                            {/* left */}
                                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" onClick={() => handleTextAlign("left")} className={`w-5 h-5 cursor-pointer transition ${textAlign === "left" ? "bg-primary" : "opacity-20 hover:opacity-100"}`}>
                                                <path strokeLinecap="round" strokeLinejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25H12" />
                                            </svg>
                                            {/* center */}
                                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" onClick={() => handleTextAlign("center")} className={`w-5 h-5 cursor-pointer transition ${textAlign === "center" ? "bg-primary" : "opacity-20 hover:opacity-100"}`}>
                                                <path strokeLinecap="round" strokeLinejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
                                            </svg>
                                            {/* right */}
                                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" onClick={() => handleTextAlign("right")} className={`w-5 h-5 cursor-pointer transition ${textAlign === "right" ? "bg-primary" : "opacity-20 hover:opacity-100"}`}>
                                                <path strokeLinecap="round" strokeLinejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5M12 17.25h8.25" />
                                            </svg>
                                        </div>
                                    </div>
                                </>
                            }
                        </div>
                    }
                </div>
            </div>
        </div>
    )
}

function componentToHex(c: number) {
    var hex = c.toString(16);
    return hex.length == 1 ? "0" + hex : hex;
}

function stringRGBToHex(rgb: string) {
    const rgbArray = rgb.split(" ")
    return "#" + componentToHex(parseInt(rgbArray[0])) + componentToHex(parseInt(rgbArray[1])) + componentToHex(parseInt(rgbArray[2]));
}

function hexToRgb(hex: string) {
    var c: any;
    if (/^#([A-Fa-f0-9]{3}){1,2}$/.test(hex)) {
        c = hex.substring(1).split('');
        if (c.length == 3) {
            c = [c[0], c[0], c[1], c[1], c[2], c[2]];
        }
        c = '0x' + c.join('');
        return [(c >> 16) & 255] + ' ' + [(c >> 8) & 255] + ' ' + [c & 255];
    }
    throw new Error('Bad Hex');
}