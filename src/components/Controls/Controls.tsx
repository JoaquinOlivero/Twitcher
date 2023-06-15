'use client';

import { createNewPlaylist, enablePreview, startStream, stopOutput } from "@/actions";
import { OutputResponse__Output } from "@/pb/service/OutputResponse";
import { forwardRef, useState } from "react";

type Props = {
    outputStatus: OutputResponse__Output | undefined
    addVideoElement: Function
    removeVideoElement: Function
}

const Controls = ({ outputStatus, addVideoElement, removeVideoElement }: Props) => {
    const [oStatus, setOStatus] = useState<OutputResponse__Output | undefined>(outputStatus)

    const handleStartPreview = async () => {

        await createNewPlaylist()

        let pc = new RTCPeerConnection({
            iceServers: [{
                urls: 'stun:stun.l.google.com:19302'
            }]
        })

        pc.oniceconnectionstatechange = e => console.log(pc.iceConnectionState)

        pc.ontrack = function (event) {
            addVideoElement(event)
            setOStatus({ ready: true })
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

    const handleStopPreview = async () => {
        await stopOutput()

        removeVideoElement()
        setOStatus({ ready: false })
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