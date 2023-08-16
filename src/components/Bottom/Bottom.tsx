'use client';

import { SongPlaylist__Output } from "@/pb/service/SongPlaylist"
import Playlist from "./Playlist/Playlist"
import Sources from "./Sources/Sources"
import { BackgroundVideosResponse__Output } from "@/pb/service/BackgroundVideosResponse";

type Props = {
    serverPlaylist: SongPlaylist__Output | undefined
    bgVideos: BackgroundVideosResponse__Output['videos'] | undefined
}

const Bottom = ({ serverPlaylist, bgVideos }: Props) => {
    return (
        <div className="w-[99%] h-full flex mx-auto gap-2">
            <Sources bgVideos={bgVideos} />
            <Playlist serverPlaylist={serverPlaylist} />

            <div className='w-1/4 h-full flex flex-col items-end rounded-t-xl bg-foreground'>
            </div>
        </div>
    )
}

export default Bottom