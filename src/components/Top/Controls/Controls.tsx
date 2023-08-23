'use client';

import { checkStatus, createNewPlaylist, enablePreview, getCurrentPlaylist, startPreview, startStream, stopPreview, stopStream } from "@/actions";
import { debounce } from "@/components/Utils/debounce";
import { usePC } from "@/context/pcContext";
import { StatusResponse__Output } from "@/pb/service/StatusResponse";
import { Dispatch, SetStateAction, useMemo, useState } from "react";

type Props = {
    status: StatusResponse__Output | undefined
    addVideoElement: Function
    removeVideoElement: Function
    streamVolume: number | undefined
}

const Controls = ({ status, addVideoElement, removeVideoElement, streamVolume }: Props) => {
    const { pc, newPc, sendMsg } = usePC();
    const [oStatus, setOStatus] = useState<StatusResponse__Output | undefined>(status)

    const connectWebRTC = async () => {
        const currentPlaylist = await getCurrentPlaylist()
        if (!currentPlaylist?.songs) {
            await createNewPlaylist()
        }

        const peer: RTCPeerConnection | null = await newPc()

        if (peer) {

            peer.oniceconnectionstatechange = e => console.log(peer.iceConnectionState)

            peer.ontrack = async function (event) {
                addVideoElement(event)
                const status: StatusResponse__Output | undefined = await checkStatus()
                setOStatus(status)
            }

            peer.onicecandidate = async event => {
                if (event.candidate === null) {
                    const sdp = btoa(JSON.stringify(peer.localDescription))
                    const serverSdp = await enablePreview(sdp)

                    try {
                        peer.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(serverSdp))))
                    } catch (e) {
                        console.log(e)
                    }
                }
            }

            // Offer to receive 1 audio, and 1 video track
            peer.addTransceiver('video', {
                'direction': 'sendrecv'
            })
            peer.addTransceiver('audio', {
                'direction': 'sendrecv'
            })

            peer.createOffer().then(d => peer.setLocalDescription(d)).catch(err => console.log(err))
        }

    }

    return (
        <div className='w-1/4 flex flex-col items-end'>
            <div className='w-full h-full bg-foreground rounded-b-xl flex justify-center items-center'>
                <div className="w-[98%] h-[98%] text-white">
                    <h2 className="w-full text-center opacity-85 text-xl font-semibold uppercase tracking-wider">
                        Controls
                    </h2>
                    <div className="w-full h-1 mx-auto my-1 bg-primary"></div>
                    {oStatus &&
                        <div className="w-full h-full flex flex-col gap-5">
                            <StreamControls connectWebRTC={connectWebRTC} pc={pc} sendMsg={sendMsg} oStatus={oStatus} setOStatus={setOStatus} volume={streamVolume} removeVideoElement={removeVideoElement} />
                            <PreviewControls connectWebRTC={connectWebRTC} pc={pc} oStatus={oStatus} setOStatus={setOStatus} removeVideoElement={removeVideoElement} />
                        </div>
                    }
                </div>
            </div>
        </div>
    )
}

export default Controls

type StreamControlsProps = {
    connectWebRTC: () => Promise<void>
    pc: RTCPeerConnection | null
    sendMsg: (msg: string) => void
    oStatus: StatusResponse__Output
    setOStatus: Dispatch<SetStateAction<StatusResponse__Output | undefined>>
    volume: number | undefined
    removeVideoElement: Function
}

const StreamControls = ({ connectWebRTC, pc, sendMsg, oStatus, setOStatus, volume, removeVideoElement }: StreamControlsProps) => {
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const [streamVolume, setStreamVolume] = useState<number | undefined>(volume)

    const handleStartStream = async () => {
        setIsLoading(true)
        await connectWebRTC()
        const res = await startStream()
        setIsLoading(false)

        if (res) {
            if (res.volume) {
                setStreamVolume(res.volume * 100)
            }

            if (res.status) {
                setOStatus(res.status)
            }
        }
    }

    const handleStopStream = async () => {
        const status = await stopStream()

        if (status) {
            setOStatus(status)
        }

        removeVideoElement()
        pc?.close()
    }


    const handleVolume = async (v: number) => {
        const msg = {
            "type": "volume",
            "volume": `${v / 100}`
        }
        setStreamVolume(v)
        sendMsg(JSON.stringify(msg))
        debouncedTrigger(v)
    }

    const debouncedTrigger = useMemo(() =>
        debounce((v: number) => {
            const msg = {
                "type": "volumeDb",
                "volume": `${v / 100}`
            }
            sendMsg(JSON.stringify(msg))
        }, 250),
        [sendMsg]
    )

    return (
        <div className="w-full flex flex-col gap-3 my-5">
            <div>
                <h2 className="opacity-85 text-md text-center font-semibold uppercase tracking-wider">
                    Stream
                </h2>
                <div className="w-1/3 h-0.5 mx-auto bg-primary"></div>
            </div>

            <div className="w-full flex items-center">
                <div className="flex w-1/2 justify-start gap-2">
                    <button
                        className={"bg-green-500 w-7 h-7 flex items-center justify-center rounded cursor-pointer outline-none transition hover:bg-green-400 " + (oStatus.stream || oStatus.preview ? "opacity-60 pointer-events-none" : "")}
                        disabled={oStatus.stream || oStatus.preview}
                        onClick={() => handleStartStream()}>
                        {!isLoading ?
                            <svg xmlns="http://www.w3.org/2000/svg" fill="white" viewBox="0 0 24 24" strokeWidth={1.5} stroke="white" className="w-5 h-5">
                                <path strokeLinecap="round" strokeLinejoin="round" d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.348a1.125 1.125 0 010 1.971l-11.54 6.347a1.125 1.125 0 01-1.667-.985V5.653z" />
                            </svg>
                            :
                            <svg className="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                            </svg>
                        }
                    </button>

                    <button
                        className={"bg-red-500 w-7 h-7 flex items-center justify-center rounded cursor-pointer transition hover:bg-red-400 " + (!oStatus.stream ? "opacity-60 pointer-events-none" : "")}
                        disabled={!oStatus.stream}
                        onClick={() => handleStopStream()}>
                        <svg xmlns="http://www.w3.org/2000/svg" fill="white" viewBox="0 0 24 24" strokeWidth={1.5} stroke="white" className="w-5 h-5">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M5.25 7.5A2.25 2.25 0 017.5 5.25h9a2.25 2.25 0 012.25 2.25v9a2.25 2.25 0 01-2.25 2.25h-9a2.25 2.25 0 01-2.25-2.25v-9z" />
                        </svg>
                    </button>
                </div>
                {oStatus.stream &&
                    <div className="w-1/2 text-center flex items-center justify-start gap-2">
                        <span className="uppercase font-semibold tracking-wider text-sm opacity-80">volume</span>
                        <input type="range" min={0} max={100} step={1} onChange={(e) => handleVolume(parseInt(e.currentTarget.value))} value={streamVolume}
                            className="appearance-none h-3 rounded-lg bg-white/25 cursor-ew-resize
                            [&::-webkit-slider-thumb]:appearance-none [&::-webkit-slider-thumb]:h-5 [&::-webkit-slider-thumb]:w-3 [&::-webkit-slider-thumb]:rounded-sm [&::-webkit-slider-thumb]:bg-primary"
                        />
                        <span className="font-semibold tracking-wider text-sm opacity-80">{streamVolume}</span>
                    </div>
                }
            </div>
        </div>
    )
}

type PreviewControlsProps = {
    connectWebRTC: () => Promise<void>
    pc: RTCPeerConnection | null
    oStatus: StatusResponse__Output
    setOStatus: Dispatch<SetStateAction<StatusResponse__Output | undefined>>
    removeVideoElement: Function
}

const PreviewControls = ({ connectWebRTC, pc, oStatus, setOStatus, removeVideoElement }: PreviewControlsProps) => {
    const [isLoading, setIsLoading] = useState<boolean>(false)

    const handleStartPreview = async () => {
        setIsLoading(true)
        await connectWebRTC()
        const status = await startPreview()
        setIsLoading(false)

        if (status) {
            setOStatus(status)
        }
    }

    const handleStopPreview = async () => {
        const status = await stopPreview()

        if (status) {
            setOStatus(status)
        }

        removeVideoElement()
        pc?.close()
    }

    return (
        <div className="w-full flex flex-col gap-3">
            <div>
                <h2 className="opacity-85 text-md text-center font-semibold uppercase tracking-wider">
                    Preview
                </h2>
                <div className="w-1/3 h-0.5 mx-auto bg-primary"></div>
            </div>

            <div className="w-full">
                <div className="flex items-center justify-center w-full gap-2">
                    <button className={"bg-green-500 w-7 h-7 flex items-center justify-center rounded cursor-pointer transition hover:bg-green-400 " + (oStatus.stream || oStatus.preview ? "opacity-60 pointer-events-none" : "")}
                        disabled={oStatus.stream || oStatus.preview}
                        onClick={() => handleStartPreview()}>
                        {!isLoading ?
                            <svg xmlns="http://www.w3.org/2000/svg" fill="white" viewBox="0 0 24 24" strokeWidth={1.5} stroke="white" className="w-5 h-5">
                                <path strokeLinecap="round" strokeLinejoin="round" d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.348a1.125 1.125 0 010 1.971l-11.54 6.347a1.125 1.125 0 01-1.667-.985V5.653z" />
                            </svg>
                            :
                            <svg className="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                            </svg>
                        }
                    </button>

                    <button className={"bg-red-500 w-7 h-7 flex items-center justify-center rounded cursor-pointer transition hover:bg-red-400 " + (oStatus.stream || !oStatus.preview ? "opacity-60 pointer-events-none" : "")}
                        disabled={!oStatus.preview}
                        onClick={() => handleStopPreview()}>
                        <svg xmlns="http://www.w3.org/2000/svg" fill="white" viewBox="0 0 24 24" strokeWidth={1.5} stroke="white" className="w-5 h-5">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M5.25 7.5A2.25 2.25 0 017.5 5.25h9a2.25 2.25 0 012.25 2.25v9a2.25 2.25 0 01-2.25 2.25h-9a2.25 2.25 0 01-2.25-2.25v-9z" />
                        </svg>
                    </button>
                </div>
            </div>
        </div>
    )
}