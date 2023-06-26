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
            <div className='w-[95%] h-full bg-foreground rounded-b-xl flex justify-center items-center'>

                <div className="w-[98%] h-[95%]">
                    {!oStatus || !oStatus.stream ?
                        <button onClick={() => handleStartStream()} className="text-white">
                            Start Stream
                        </button>
                        :
                        <button onClick={() => handleStopStream()} className="text-white">
                            Stop Stream
                        </button>
                    }

                    {!oStatus || !oStatus.stream &&
                        <>
                            {!oStatus || !oStatus.output ?
                                <button onClick={() => handleStartPreview()} className="relative inline-flex items-center justify-center p-0.5 mb-2 mr-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-cyan-500 to-blue-500 group-hover:from-cyan-500 group-hover:to-blue-500 hover:text-white dark:text-white">
                                    <span className="relative px-2.5 py-1.5 transition-all ease-in duration-75 bg-white dark:bg-foreground rounded-md group-hover:bg-opacity-0">
                                        Start Preview
                                    </span>
                                </button>
                                :
                                <button onClick={() => handleStopPreview()} className="relative inline-flex items-center justify-center p-0.5 mb-2 mr-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-pink-500 to-orange-400 group-hover:from-pink-500 group-hover:to-orange-400 hover:text-white dark:text-white">
                                    <span className="relative px-2.5 py-1.5 transition-all ease-in duration-75 bg-white dark:bg-foreground rounded-md group-hover:bg-opacity-0">
                                        Stop Preview
                                    </span>
                                </button>
                            }
                        </>
                    }
                </div>

            </div>
        </div>
    )
}

export default Controls