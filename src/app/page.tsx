import Preview from '@/components/Preview/Preview'
import Playlist from '@/components/Playlist/Playlist'

const Home = async () => {

    return (
        <div className='w-[99%] h-screen mx-auto flex flex-col'>

            <div className='w-full h-3/5'>
                <Preview />
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

                <Playlist />

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
