'use client';

import { enablePreview, getCurrentPlaylist, updateSongPlaylist } from "@/actions";
import { Song__Output } from "@/pb/service/Song";
import { SongPlaylist__Output } from "@/pb/service/SongPlaylist";
import Image from "next/image";
import { useEffect, useRef, useState } from "react";

type Props = {
    serverPlaylist: SongPlaylist__Output | undefined
}

const List = ({ serverPlaylist }: Props) => {
    const playlistRef = useRef<HTMLDivElement>(null);
    const dragItem = useRef<number | null>(null);
    const dragOverItem = useRef<number | null>(null);
    const [list, setList] = useState<SongPlaylist__Output | undefined>(serverPlaylist)

    const handleDragStart = (e: React.DragEvent<HTMLDivElement>, position: number) => {
        dragItem.current = position;
        e.dataTransfer.dropEffect = "none"
    }

    const handleDragEnter = (position: number) => {
        dragOverItem.current = position;
    }

    const handleDragEnd = async () => {
        if (list && dragItem.current !== null && dragOverItem.current !== null && playlistRef.current !== null && list.songs) {
            const copyListItems = [...list.songs];
            const dragItemContent = copyListItems[dragItem.current];
            copyListItems.splice(dragItem.current, 1);
            copyListItems.splice(dragOverItem.current, 0, dragItemContent);

            setList({ songs: copyListItems });

            const dragItemDiv = playlistRef.current.children[dragOverItem.current]
            dragItemDiv.classList.remove("opacity-40")
            dragItemDiv.classList.remove("border-l-2")
            dragItemDiv.classList.remove("border-lime-500")
            dragItemDiv.classList.remove("bg-background")

            dragItem.current = null;
            dragOverItem.current = null;

            await updateSongPlaylist({ songs: copyListItems })
        }
    }

    const handleDragOver = (e: React.DragEvent<HTMLDivElement>) => {
        e.preventDefault()
        e.currentTarget.classList.add("opacity-40")
        e.currentTarget.classList.add("border-l-2")
        e.currentTarget.classList.add("border-lime-500")
        e.currentTarget.classList.add("bg-background")
    }

    const handleDragLeave = (e: React.DragEvent<HTMLDivElement>) => {
        e.preventDefault()
        e.currentTarget.classList.remove("opacity-40")
        e.currentTarget.classList.remove("border-l-2")
        e.currentTarget.classList.remove("border-lime-500")
        e.currentTarget.classList.remove("bg-background")
    }

    useEffect(() => {

        let pc = new RTCPeerConnection({
            iceServers: [
                {
                    urls: 'stun:stun.l.google.com:19302'
                }
            ]
        })

        pc.oniceconnectionstatechange = e => console.log(pc.iceConnectionState)

        let sendChannel = pc.createDataChannel('updateplaylist')

        sendChannel.onclose = () => console.log('sendChannel has closed')
        sendChannel.onopen = () => console.log('sendChannel has opened')

        sendChannel.onmessage = async e => {
            const playlist = await getCurrentPlaylist();
            setList(playlist)
        }

        pc.oniceconnectionstatechange = e => console.log(pc.iceConnectionState)

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

        pc.createOffer().then(d => pc.setLocalDescription(d)).catch(err => console.log(err))

    }, [])


    useEffect(() => {
    }, [list])


    return (
        <>
            {list && list.songs &&
                <div className="overflow-y-scroll h-full py-1 scrollbar" ref={playlistRef} onDragOver={(e) => { e.preventDefault() }}>
                    {list.songs.map((song: Song__Output, index: number) => {
                        return <div key={song.page} className="flex gap-2 p-2 hover:bg-background hover:cursor-grab active:cursor-grabbing"
                            draggable
                            onDragStart={(e) => handleDragStart(e, index)}
                            onDragEnter={() => handleDragEnter(index)}
                            onDragOver={(e) => { handleDragOver(e) }}
                            onDragLeave={(e) => { handleDragLeave(e) }}
                            onDragEnd={() => handleDragEnd()}
                        >
                            <div className="w-12 h-12">
                                <Image src={`/api/covers/${song.cover}`} width={100} height={100} alt="" />
                            </div>

                            <div className="flex flex-col">
                                <div className="text-white">
                                    {song.name}
                                </div>

                                <div className="text-white opacity-60">
                                    {song.author}
                                </div>
                            </div>

                        </div>
                    })}
                </div>
            }
        </>
    )
}

export default List