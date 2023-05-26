'use client';
import { Song__Output } from "@/pb/service/Song";
import React, { useState, useTransition, useRef, createElement } from "react";
// import { createNewPlaylist } from "@/actions";

type Props = {
    songs: Song__Output[] | undefined
}

const Playlist = ({ songs }: Props) => {
    const playlistRef = useRef<HTMLDivElement>(null);
    const dragItem = useRef<number | null>(null);
    const dragOverItem = useRef<number | null>(null);
    const [list, setList] = useState<Song__Output[] | undefined>(songs)
    // let [isPending, startTransition] = useTransition()

    const handleDragStart = (e: React.DragEvent<HTMLDivElement>, position: number) => {
        dragItem.current = position;
        e.dataTransfer.dropEffect = "none"
    }

    const handleDragEnter = (position: number) => {
        dragOverItem.current = position;
    }

    const handleDragEnd = () => {
        if (list && dragItem.current !== null && dragOverItem.current !== null && playlistRef.current !== null) {
            const copyListItems = [...list];
            const dragItemContent = copyListItems[dragItem.current];
            copyListItems.splice(dragItem.current, 1);
            copyListItems.splice(dragOverItem.current, 0, dragItemContent);

            setList(copyListItems);

            const dragItemDiv = playlistRef.current.children[dragOverItem.current]
            dragItemDiv.classList.remove("opacity-40")
            dragItemDiv.classList.remove("border-l-2")
            dragItemDiv.classList.remove("border-lime-500")
            dragItemDiv.classList.remove("bg-background")

            dragItem.current = null;
            dragOverItem.current = null;
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

    return (
        <div className="bg-foreground shadow-lg w-1/2 h-[95%] rounded-t-xl font-sans overflow-hidden">
            {/* <button onClick={() => startTransition(() => createNewPlaylist())}>Create new playlist</button> */}
            <div className="text-[#fff] w-[100%] my-1 text-center opacity-80 text-xl font-semibold uppercase tracking-wider">
                Playlist
            </div>

            <div className="w-[98%] h-1 mx-auto my-1 bg-background"></div>

            {list &&
                <div className="overflow-y-scroll h-full py-1 scrollbar" ref={playlistRef} onDragOver={(e) => { e.preventDefault() }}>
                    {list.map((song: Song__Output, index: number) => {
                        return <div key={song.page} className="flex gap-2 p-2 hover:bg-background hover:cursor-grab active:cursor-grabbing"
                            draggable
                            onDragStart={(e) => handleDragStart(e, index)}
                            onDragEnter={() => handleDragEnter(index)}
                            onDragOver={(e) => { handleDragOver(e) }}
                            onDragLeave={(e) => { handleDragLeave(e) }}
                            onDragEnd={() => handleDragEnd()}
                        >
                            <div className="w-12 h-12">
                                <img src={`/api/covers/${song.cover}`} />
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
        </div>
    )
}

export default Playlist