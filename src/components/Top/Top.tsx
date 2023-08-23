'use client';

import Preview from "./Preview/Preview";
import Controls from "./Controls/Controls";
import { useRef, useState } from "react";
import { StatusResponse__Output } from "@/pb/service/StatusResponse";
import Settings from "./Settings/Settings";
import { TwitchStreamKey__Output } from "@/pb/service/TwitchStreamKey";
import { DevCredentials__Output } from "@/pb/service/DevCredentials";
import { checkStatus } from "@/actions";
import { usePC } from "@/context/pcContext";
import { StreamParametersResponse__Output } from "@/pb/service/StreamParametersResponse";

type Props = {
    status: StatusResponse__Output | undefined
    statusStreamKey: TwitchStreamKey__Output | undefined
    twitchCredentials: DevCredentials__Output | undefined
    streamParams: StreamParametersResponse__Output | undefined
}

const Top = ({ status, statusStreamKey, twitchCredentials, streamParams }: Props) => {
    const { setIsPreviewLoaded, setVideoElementSize } = usePC();
    const vRef = useRef<HTMLDivElement>(null)
    const [streamStatus, setStreamStatus] = useState<StatusResponse__Output | undefined>(status)
    const [muted, setMuted] = useState<boolean>(typeof window !== 'undefined' && localStorage.getItem("volume") != null && Number(localStorage.getItem("volume")) == 0 ? true : false)
    const [volume, setVolume] = useState<number>(typeof window !== 'undefined' && localStorage.getItem("volume") != null ? Number(localStorage.getItem("volume")) : 10)
    const [prevVolume, setPrevVolume] = useState<number>(typeof window !== 'undefined' && localStorage.getItem("volume") != null ? Number(localStorage.getItem("volume")) : 10)

    const addVideoElement = (event: RTCTrackEvent) => {
        if (event.track.kind === "video" && vRef.current) {
            const parent: HTMLDivElement | null = vRef.current

            const video: HTMLVideoElement | null = document.createElement("video")
            video.srcObject = event.streams[0]
            video.autoplay = true
            video.muted = true
            setMuted(true)
            setVolume(0)

            video.onloadeddata = function () {
                const boundingClient = video.getBoundingClientRect()
                setVideoElementSize({ width: boundingClient.width, height: boundingClient.height })
                setIsPreviewLoaded(true)
            }

            video.onerror = function () {
                setIsPreviewLoaded(false)
                console.log("couldn't load video. Something went wrong")
            }

            parent.appendChild(video)
        }
    }

    const removeVideoElement = async () => {
        if (vRef.current) {
            vRef.current.children[1].remove()
            const status = await checkStatus();
            setStreamStatus(status)
            setIsPreviewLoaded(false)
        }
    }

    const handleSoundMuting = () => {
        if (vRef.current && vRef.current.children[1]) {
            const videoElement: HTMLVideoElement = vRef.current.children[1] as HTMLVideoElement

            if (muted) {
                videoElement.volume = prevVolume / 100
                setVolume(prevVolume)
                setPrevVolume(volume)

                localStorage.setItem("volume", prevVolume.toString())
            }

            if (!muted) {
                setPrevVolume(volume)
                videoElement.volume = 0
                setVolume(0)

                localStorage.setItem("volume", "0")
            }


            videoElement.muted = !muted
            setMuted(!muted)
        }
    }

    const handleVolume = (value: number) => {
        if (vRef.current && vRef.current.children[1]) {
            const videoElement: HTMLVideoElement | null = vRef.current.children[1] as HTMLVideoElement

            videoElement.volume = value / 100
            setVolume(value)

            localStorage.setItem("volume", value.toString())

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
            <div className="w-[99%] h-full flex mx-auto gap-2">
                <Settings statusStreamKey={statusStreamKey} twitchCredentials={twitchCredentials} />
                <Preview
                    status={streamStatus}
                    addVideoElement={addVideoElement}
                    handleSoundMuting={handleSoundMuting}
                    handleVolume={handleVolume}
                    muted={muted}
                    volume={volume}
                    streamParams={streamParams}
                    ref={vRef}
                />
                <Controls
                    status={streamStatus}
                    addVideoElement={addVideoElement}
                    removeVideoElement={removeVideoElement}
                    streamVolume={streamParams?.volume}
                />
            </div>
        </>
    )
}

export default Top