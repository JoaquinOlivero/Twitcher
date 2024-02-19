import { checkStatus, checkTwitchCredentials, checkTwitchStreamKey, checkYoutubeParams, getBgVideos, getCurrentPlaylist, getStreamParams } from '@/actions';
import Bottom from '@/components/Bottom/Bottom';
import Top from '@/components/Top/Top';
import PCProvider from '@/context/pcContext';

export const revalidate = 0;

const Home = async () => {
    const playlist = await getCurrentPlaylist();
    const status = await checkStatus();
    const statusStreamKey = await checkTwitchStreamKey();
    const twitchCredentials = await checkTwitchCredentials();
    const bgVideos = await getBgVideos();
    const streamParams = await getStreamParams();
    const youtubeParams = await checkYoutubeParams();

    return (
        <PCProvider>
            <div className='relative w-full h-screen mx-auto flex flex-col items-center gap-2'>
                <div className='w-full h-3/5'>
                    <Top status={status} statusStreamKey={statusStreamKey} youtubeParams={youtubeParams} twitchCredentials={twitchCredentials} streamParams={streamParams} />
                </div>

                <div className='w-full h-2/5'>
                    <Bottom serverPlaylist={playlist} bgVideos={bgVideos?.videos} />
                </div>
            </div>
        </PCProvider>
    )
}

export default Home
