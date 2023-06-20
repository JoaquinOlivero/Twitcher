'use client'

import { ReactNode, createContext, useContext, useState } from 'react'

type PCContextType = {
    pc: RTCPeerConnection | null;
    newPc: () => Promise<RTCPeerConnection | null>;
    updatePlaylistDataChan: RTCDataChannel | null;
}

const PCContextDefaultValue: PCContextType = {
    pc: null,
    newPc: () => Promise.resolve(null),
    updatePlaylistDataChan: null
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
    const [updatePlaylistDataChan, setUpdatePlaylistDataChan] = useState<RTCDataChannel | null>(null)

    const newPc = async () => {
        let pc = new RTCPeerConnection({
            iceServers: [{
                urls: 'stun:stun.l.google.com:19302'
            }]
        })

        let dataChan = pc.createDataChannel('updateplaylist')

        setUpdatePlaylistDataChan(dataChan)

        setPc(pc)

        return pc
    }

    const value = {
        pc,
        newPc,
        updatePlaylistDataChan
    };

    return <PCContext.Provider value={value}>{children}</PCContext.Provider>
}