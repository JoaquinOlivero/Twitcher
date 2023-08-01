'use client'

import { getOverlays } from '@/actions'
import { Dispatch, MutableRefObject, ReactNode, SetStateAction, createContext, useContext, useRef, useState } from 'react'
import * as fabric from 'fabric'; // v6

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

type VideoElementSize = {
    width: number
    height: number
}

type PCContextType = {
    pc: RTCPeerConnection | null;
    newPc: () => Promise<RTCPeerConnection | null>;
    Overlays: Overlay[] | null;
    setOverlays: Dispatch<SetStateAction<Overlay[] | null>>;
    sendMsg: (msg: string) => void;
    isPreviewLoaded: boolean;
    setIsPreviewLoaded: Dispatch<SetStateAction<boolean>>;
    fabricRef: MutableRefObject<fabric.Canvas | null>;
    videoElementSize: VideoElementSize | null;
    setVideoElementSize: Dispatch<SetStateAction<VideoElementSize | null>>
}

const PCContextDefaultValue: PCContextType = {
    pc: null,
    newPc: () => Promise.resolve(null),
    Overlays: null,
    setOverlays: () => { },
    sendMsg: () => { },
    isPreviewLoaded: false,
    setIsPreviewLoaded: () => { },
    fabricRef: null!,
    videoElementSize: null,
    setVideoElementSize: () => { },
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
    const [isPreviewLoaded, setIsPreviewLoaded] = useState<boolean>(false)
    const [videoElementSize, setVideoElementSize] = useState<VideoElementSize | null>(null)
    const fabricRef = useRef<fabric.Canvas | null>(null);

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
        sendMsg,
        isPreviewLoaded,
        setIsPreviewLoaded,
        fabricRef,
        videoElementSize,
        setVideoElementSize
    };

    return <PCContext.Provider value={value}>{children}</PCContext.Provider>
}