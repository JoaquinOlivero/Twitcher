'use client';

import { createNewPlaylist, enablePreview, getCurrentPlaylist, stopOutput } from "@/actions";
import { usePC } from "@/context/pcContext";
import { OutputResponse__Output } from "@/pb/service/OutputResponse";
import { useState } from "react";

type Props = {
    outputStatus: OutputResponse__Output | undefined
    addVideoElement: Function
    removeVideoElement: Function
}

const Controls = ({ outputStatus, addVideoElement, removeVideoElement }: Props) => {
    const { pc, newPc } = usePC();
    const [oStatus, setOStatus] = useState<OutputResponse__Output | undefined>(outputStatus)

    const handleStartPreview = async () => {

        const currentPlaylist = await getCurrentPlaylist()
        if (!currentPlaylist?.songs) {
            await createNewPlaylist()
        }

        const peer: RTCPeerConnection | null = await newPc()

        if (peer) {

            peer.oniceconnectionstatechange = e => console.log(peer.iceConnectionState)

            peer.ontrack = function (event) {
                addVideoElement(event)
                setOStatus({ ready: true })
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
        setOStatus({ ready: false })
        pc?.close()
    }

    return (
        <div className='w-1/4 flex flex-col items-end'>
            <div className='w-[95%] h-full bg-foreground rounded-b-xl flex justify-center items-center'>

                <div className="w-[98%] h-[95%]">
                    {!oStatus || !oStatus.ready ?
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
                </div>

            </div>
        </div>
    )
}

export default Controls