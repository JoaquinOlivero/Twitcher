import Playlist from '@/components/Playlist/Playlist'
import { checkStatus, checkTwitchCredentials, checkTwitchStreamKey, getCurrentPlaylist } from '@/actions';
import Top from '@/components/Top/Top';
import PCProvider from '@/context/pcContext';

export const revalidate = 0;

const Home = async () => {
    const playlist = await getCurrentPlaylist();
    const status = await checkStatus();
    const statusStreamKey = await checkTwitchStreamKey();
    const twitchCredentials = await checkTwitchCredentials();

    return (
        <PCProvider>
            <div className='relative w-full h-screen mx-auto flex flex-col items-center'>

                <div className='w-[100%] h-3/5 flex'>
                    <Top status={status} statusStreamKey={statusStreamKey} twitchCredentials={twitchCredentials} />
                </div>

                <div className='w-[99%] h-2/5 flex items-center'>

                    <div className='w-1/4 h-[95%]'>
                        <div className='rounded-t-xl bg-foreground w-[95%] h-1/2'>
                            {/* Video Source */}
                        </div>

                        <div className='bg-foreground w-[95%] h-1/2'>
                            {/* Layouts */}
                        </div>
                    </div>

                    <Playlist serverPlaylist={playlist} />

                    <div className='w-1/4 h-[95%] flex flex-col items-end'>
                        <div className='rounded-t-xl bg-foreground w-[95%] h-full'>
                            {/* Search Song */}
                        </div>
                    </div>

                </div>

            </div>
        </PCProvider>
    )
}

export default Home
