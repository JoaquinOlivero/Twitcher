import { SongPlaylist__Output } from "@/pb/service/SongPlaylist";
import List from "./components/List";
import { createNewPlaylist } from "@/actions";

type Props = {
    serverPlaylist: SongPlaylist__Output | undefined
}

export const revalidate = 0;

const Playlist = ({ serverPlaylist }: Props) => {

    return (
        <div className="bg-foreground shadow-lg w-1/2 h-[95%] rounded-t-xl font-sans overflow-hidden">
            <div className="text-[#fff] w-full my-1 relative">
                <div className="text-center opacity-80 text-xl font-semibold uppercase tracking-wider">
                    Playlist
                </div>
            </div>


            <div className="w-[98%] h-1 mx-auto my-1 bg-background"></div>


            {serverPlaylist && serverPlaylist.songs &&
                <List serverPlaylist={serverPlaylist} />
            }
        </div>
    )
}

export default Playlist