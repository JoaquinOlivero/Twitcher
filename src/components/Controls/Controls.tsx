'use client';

import { enablePreview, preparePreview } from "@/actions";
import { createElement, useRef } from "react";

const Controls = () => {
    const controlsRef = useRef<HTMLDivElement>(null)

    const handleOnSubmit = async () => {

        let pc = new RTCPeerConnection({
            iceServers: [{
                urls: 'stun:stun.l.google.com:19302'
            }]
        })

        pc.oniceconnectionstatechange = e => console.log(pc.iceConnectionState)

        pc.ontrack = function (event) {
            var el: any = document.createElement(event.track.kind)
            // var el: any = createElement("video")
            el.srcObject = event.streams[0]
            el.autoplay = true
            el.controls = true

            // document.getElementById('remoteVideos')!.appendChild(el)
            const div = controlsRef.current
            div?.appendChild(el)
        }

        pc.onicecandidate = async event => {
            if (event.candidate === null) {
                const sdp = btoa(JSON.stringify(pc.localDescription))
                await preparePreview()
                const serverSdp = await enablePreview(sdp)
                // console.log(serverSdp)

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

        // preparePreview()
        // const serverSdp = await enablePreview(sdp)
        // console.log(serverSdp)
    }

    return (
        <div className='w-1/4 flex flex-col items-end' ref={controlsRef}>
            <div className='w-[95%] h-full bg-foreground rounded-b-xl'>
                <button className='text-white' onClick={() => handleOnSubmit()}>Enable preview</button>
            </div>
        </div>
    )
}

export default Controls