'use client';

import { OutputResponse__Output } from "@/pb/service/OutputResponse";
import Preview from "../Preview/Preview";
import Controls from "../Controls/Controls";
import { useRef, useState } from "react";

type Props = {
    outputStatus: OutputResponse__Output | undefined
}

const Top = ({ outputStatus }: Props) => {
    const vRef = useRef<HTMLVideoElement>(null)
    const [muted, setMuted] = useState<boolean>(false)
    const [volume, setVolume] = useState<number>(10)
    const [prevVolume, setPrevVolume] = useState<number>(10)
    const [isLoaded, setIsLoaded] = useState<boolean>(false)

    const addVideoElement = (event: RTCTrackEvent) => {
        if (event.track.kind === "video" && vRef.current) {
            var el: HTMLVideoElement | null = vRef.current
            el.srcObject = event.streams[0]
            el.autoplay = true
            el.volume = volume / 100
            setIsLoaded(true)
        }
    }

    const removeVideoElement = () => {
        if (vRef.current) {
            vRef.current.remove()
            setIsLoaded(false)
        }
    }

    const handleSoundMuting = () => {
        if (vRef.current) {
            const videoElement: HTMLVideoElement | null = vRef.current

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
        if (vRef.current) {
            const videoElement: HTMLVideoElement | null = vRef.current
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
            <Preview
                outputStatus={outputStatus}
                addVideoElement={addVideoElement}
                handleSoundMuting={handleSoundMuting}
                handleVolume={handleVolume}
                muted={muted}
                isLoaded={isLoaded}
                volume={volume}
                ref={vRef}
            />
            <Controls
                outputStatus={outputStatus}
                addVideoElement={addVideoElement}
                removeVideoElement={removeVideoElement}
            />
        </>
    )
}

export default Top