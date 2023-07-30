import { SongPlaylist__Output } from "@/pb/service/SongPlaylist"
import Playlist from "./Playlist/Playlist"

type Props = {
    serverPlaylist: SongPlaylist__Output | undefined
}

const Bottom = ({ serverPlaylist }: Props) => {
    return (
        <div className="w-[99%] h-full flex mx-auto gap-2">
            <div className='w-1/4 h-full bg-foreground rounded-t-xl'>
                {/* Layouts */}
            </div>

            <Playlist serverPlaylist={serverPlaylist} />

            <div className='w-1/4 h-full flex flex-col items-end rounded-t-xl bg-foreground'>
            </div>
        </div>
    )
}

export default Bottom