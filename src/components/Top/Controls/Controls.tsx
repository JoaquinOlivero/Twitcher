'use client';

import { checkStatus, createNewPlaylist, enablePreview, getCurrentPlaylist, startStream, stopOutput, stopStream } from "@/actions";
import { usePC } from "@/context/pcContext";
import { StatusResponse__Output } from "@/pb/service/StatusResponse";
import { useState } from "react";

type Props = {
    status: StatusResponse__Output | undefined
    addVideoElement: Function
    removeVideoElement: Function
}

const Controls = ({ status, addVideoElement, removeVideoElement }: Props) => {
    const { pc, newPc } = usePC();
    const [oStatus, setOStatus] = useState<StatusResponse__Output | undefined>(status)

    const handleStartPreview = async () => {

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

    const handleStopPreview = async () => {
        await stopOutput()

        removeVideoElement()

        const status: StatusResponse__Output | undefined = await checkStatus()
        setOStatus(status)
        pc?.close()
    }

    const handleStartStream = async () => {
        await handleStartPreview()
        await startStream()

        const status: StatusResponse__Output | undefined = await checkStatus()
        setOStatus(status)
    }

    const handleStopStream = async () => {
        await stopStream()

        await stopOutput()

        removeVideoElement()

        const status: StatusResponse__Output | undefined = await checkStatus()
        setOStatus(status)
        pc?.close()
    }

    return (
        <div className='w-1/4 flex flex-col items-end'>
            <div className='w-full h-full bg-foreground rounded-b-xl flex justify-center items-center'>
                <div className="w-[98%] h-[98%]">
                    {oStatus &&
                        <div className="flex gap-2">
                            <button className={"bg-green-500 w-7 h-7 flex items-center justify-center rounded cursor-pointer transition hover:bg-green-400 " + (oStatus.stream || oStatus.output ? "opacity-60 pointer-events-none" : "")} disabled={oStatus.stream || oStatus.output} onClick={() => handleStartStream()}>
                                <svg xmlns="http://www.w3.org/2000/svg" fill="white" viewBox="0 0 24 24" strokeWidth={1.5} stroke="white" className="w-5 h-5">
                                    <path strokeLinecap="round" strokeLinejoin="round" d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.348a1.125 1.125 0 010 1.971l-11.54 6.347a1.125 1.125 0 01-1.667-.985V5.653z" />
                                </svg>
                            </button>
                            <button className={"bg-red-500 w-7 h-7 flex items-center justify-center rounded cursor-pointer transition hover:bg-red-400 " + (!oStatus.stream ? "opacity-60 pointer-events-none" : "")} disabled={!oStatus.stream} onClick={() => handleStopStream()}>
                                <svg xmlns="http://www.w3.org/2000/svg" fill="white" viewBox="0 0 24 24" strokeWidth={1.5} stroke="white" className="w-5 h-5">
                                    <path strokeLinecap="round" strokeLinejoin="round" d="M5.25 7.5A2.25 2.25 0 017.5 5.25h9a2.25 2.25 0 012.25 2.25v9a2.25 2.25 0 01-2.25 2.25h-9a2.25 2.25 0 01-2.25-2.25v-9z" />
                                </svg>
                            </button>


                            {!oStatus.stream &&
                                <>
                                    {!oStatus.output ?
                                        <button className="text-white text-sm font-semibold border-2 rounded px-3.5 py-0.5 transition hover:border-lime-500 hover:bg-lime-500" onClick={() => handleStartPreview()}>
                                            Start Preview
                                        </button>
                                        :
                                        <button className="text-white text-sm font-semibold border-2 rounded px-3.5 py-0.5 transition hover:border-rose-500 hover:bg-rose-500" onClick={() => handleStopPreview()}>
                                            Stop Preview
                                        </button>
                                    }
                                </>
                            }
                        </div>
                    }
                </div>
            </div>
        </div>
    )
}

export default Controls