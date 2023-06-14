'use client';

import { enablePreview, startStream } from "@/actions";

const Controls = () => {

    const handleOnSubmit = async () => {

        let pc = new RTCPeerConnection({
            iceServers: [{
                urls: 'stun:stun.l.google.com:19302'
            }]
        })

        pc.oniceconnectionstatechange = e => console.log(pc.iceConnectionState)

        pc.ontrack = function (event) {
            if (event.track.kind === "video") {
                var el: any = document.createElement(event.track.kind)
                el.srcObject = event.streams[0]
                el.autoplay = true
                el.controls = true
                el.volume = 0.1
                document.getElementById('remoteVideos')!.appendChild(el)
            }
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

    return (
        <div className='w-1/4 flex flex-col items-end'>
            <div className='w-[95%] h-full bg-foreground rounded-b-xl'>
                <button className='text-white' onClick={() => handleOnSubmit()}>Start preview</button>
                <form className="text-white" action={startStream}>
                    <button>Start stream</button>
                </form>
            </div>
        </div>
    )
}

export default Controls