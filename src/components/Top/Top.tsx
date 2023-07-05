'use client';

import Preview from "../Preview/Preview";
import Controls from "../Controls/Controls";
import { useRef, useState } from "react";
import { StatusResponse__Output } from "@/pb/service/StatusResponse";
import Settings from "./Settings/Settings";

type Props = {
    status: StatusResponse__Output | undefined
}

const Top = ({ status }: Props) => {
    const vRef = useRef<HTMLDivElement>(null)
    const [muted, setMuted] = useState<boolean>(false)
    const [volume, setVolume] = useState<number>(10)
    const [prevVolume, setPrevVolume] = useState<number>(10)
    const [isLoaded, setIsLoaded] = useState<boolean>(false)
    const [isModalOpen, setIsModalOpen] = useState<boolean>(false)

    const addVideoElement = (event: RTCTrackEvent) => {
        if (event.track.kind === "video" && vRef.current) {
            const parent: HTMLDivElement | null = vRef.current

            const video: HTMLVideoElement | null = document.createElement("video")
            video.srcObject = event.streams[0]
            video.autoplay = true
            video.volume = volume / 100

            parent.appendChild(video)

            setIsLoaded(true)
        }
    }

    const removeVideoElement = () => {
        if (vRef.current) {
            vRef.current.children[0].remove()
            setIsLoaded(false)
        }
    }

    const handleSoundMuting = () => {
        if (vRef.current && vRef.current.children[0]) {
            const videoElement: HTMLVideoElement = vRef.current.children[0] as HTMLVideoElement

            if (muted) {
                videoElement.volume = prevVolume / 100
                setVolume(prevVolume)
                setPrevVolume(volume)
            }

            if (!muted) {
                setPrevVolume(volume)
                videoElement.volume = 0
                setVolume(0)
            }


            videoElement.muted = !muted
            setMuted(!muted)


        }
    }

    const handleVolume = (value: number) => {
        if (vRef.current && vRef.current.children[0]) {
            const videoElement: HTMLVideoElement | null = vRef.current.children[0] as HTMLVideoElement
            videoElement.volume = value / 100
            setVolume(value)
            if (muted) {
                videoElement.muted = false
                setMuted(false)
            }

            if (value === 0) {
                videoElement.muted = true
                setMuted(true)
            }
        }
    }

    return (
        <>

            <div className="w-[99%] flex mx-auto">
                <Settings />
                <Preview
                    status={status}
                    addVideoElement={addVideoElement}
                    handleSoundMuting={handleSoundMuting}
                    handleVolume={handleVolume}
                    muted={muted}
                    isLoaded={isLoaded}
                    volume={volume}
                    ref={vRef}
                />
                <Controls
                    status={status}
                    addVideoElement={addVideoElement}
                    removeVideoElement={removeVideoElement}
                />
            </div>
        </>
    )
}

export default Top

// type SettingsButtonProps = {
//     text: string
// }

// const SettingsButton = ({ text }: SettingsButtonProps) => {

//     return (
//         <div className="py-1 my-1 transition text-sm capitalize font-semibold tracking-wider opacity-80 hover:bg-background hover:opacity-100 cursor-pointer">
//             <span className="ml-2">{text}</span>
//         </div>
//     )
// }

// type ModalProps = {
//     tWidth: string
//     tHeight: string
//     isOpen: boolean
//     setIsOpen: React.Dispatch<React.SetStateAction<boolean>>
// }

// const Modal = ({ tWidth, tHeight, isOpen, setIsOpen }: ModalProps) => {

//     return (
//         <>
//             {isOpen &&

//                 <div className="absolute left-0 top-0 w-full h-full z-20 flex justify-center items-center">

//                     {/* Transparent background */}
//                     <div className="absolute w-full h-full opacity-50 bg-black" onClick={() => setIsOpen(false)}></div>

//                     {/* Content */}
//                     <div className={`absolute bg-foreground z-30 rounded-lg ${tWidth} ${tHeight}`}>

//                     </div>

//                 </div>
//             }
//         </>
//     )
// }