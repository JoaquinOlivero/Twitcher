import Preview from '@/components/Preview/Preview'
import Playlist from '@/components/Playlist/Playlist'
import Controls from '@/components/Controls/Controls'
import { checkOutputStatus, getCurrentPlaylist } from '@/actions';
import Top from '@/components/Top/Top';

export const revalidate = 0;

const Home = async () => {
    const playlist = await getCurrentPlaylist();
    const outputStatus = await checkOutputStatus();

    return (
        <div className='w-[99%] h-screen mx-auto flex flex-col'>

            <div className='relative w-full h-3/5 flex'>
                <div className='w-1/4'>

                </div>
                {/* <Preview outputStatus={outputStatus} />
                <Controls outputStatus={outputStatus} /> */}
                <Top outputStatus={outputStatus} />
            </div>

            <div className='w-full h-2/5 flex items-center'>

                <div className='w-1/4 h-[95%]'>
                    <div className='rounded-t-xl bg-foreground w-[95%] h-1/2'>
                        Video Source
                    </div>

                    <div className='bg-foreground w-[95%] h-1/2'>
                        Layouts
                    </div>
                </div>

                <Playlist serverPlaylist={playlist} />

                <div className='w-1/4 h-[95%] flex flex-col items-end'>
                    <div className='rounded-t-xl bg-foreground w-[95%] h-full'>
                        Search Song
                    </div>
                </div>

            </div>

        </div>
    )
}

export default Home
