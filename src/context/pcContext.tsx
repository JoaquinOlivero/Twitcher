'use client'

import { getOverlays } from '@/actions'
import { Dispatch, ReactNode, SetStateAction, createContext, useContext, useState } from 'react'

type DataChannelMsg = {
    type: string
    message: Overlay | {}
}

type Overlay = {
    id: string
    width: number
    height: number
    pointX: number
    pointY: number
    show: boolean
    coverId: string
    text: string
    fontFamily: string
    fontSize: number
    lineHeight: number
    textColor: string
}

type PCContextType = {
    pc: RTCPeerConnection | null;
    newPc: () => Promise<RTCPeerConnection | null>;
    Overlays: Overlay[] | null;
    setOverlays: Dispatch<SetStateAction<Overlay[] | null>>;
    sendMsg: (msg: string) => void
}

const PCContextDefaultValue: PCContextType = {
    pc: null,
    newPc: () => Promise.resolve(null),
    Overlays: null,
    setOverlays: () => { },
    sendMsg: () => { }
}

export const PCContext = createContext<PCContextType>(PCContextDefaultValue)

type Props = {
    children: ReactNode
}

export function usePC() {
    return useContext(PCContext);
}

export default function PCProvider({ children }: Props) {
    const [pc, setPc] = useState<RTCPeerConnection | null>(null)
    const [Overlays, setOverlays] = useState<Array<Overlay> | null>(null)
    const [dataChan, setDataChan] = useState<RTCDataChannel | null>(null)

    const newPc = async () => {
        setPc(null)
        const data = await getOverlays()
        if (data && data.overlays) {
            setOverlays(data.overlays as Overlay[])
        }


        let pc = new RTCPeerConnection({
            iceServers: [{
                urls: 'stun:stun.l.google.com:19302'
            }]
        })

        createDataChannel(pc)
        setPc(pc)

        return pc
    }

    const createDataChannel = (pc: RTCPeerConnection) => {
        const dc = pc.createDataChannel('streamDataChannel')
        setDataChan(dc)

        dc.onopen = () => console.log("open data channel")
        dc.onclose = () => console.log("close data channel")

        dc.onmessage = (event) => {
            const data: DataChannelMsg = JSON.parse(event.data)
            if (data.type === "overlay") setOverlays(data.message as Overlay[])
        }
    }

    const sendMsg = async (msg: string) => {
        if (dataChan) {
            dataChan.send(msg)
        }
    }

    const value = {
        pc,
        newPc,
        Overlays,
        setOverlays,
        sendMsg
    };

    return <PCContext.Provider value={value}>{children}</PCContext.Provider>
}