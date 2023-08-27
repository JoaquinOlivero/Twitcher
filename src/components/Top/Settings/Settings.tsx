import { DevCredentials__Output } from '@/pb/service/DevCredentials'
import AddSongs from './AddSongs/AddSongs'
import TwitchSettings from './TwitchSettings/TwitchSettings'
import { TwitchStreamKey__Output } from '@/pb/service/TwitchStreamKey'
import StreamParams from './StreamParams/StreamParams'
import { StreamParametersResponse__Output } from '@/pb/service/StreamParametersResponse'

type Props = {
    statusStreamKey: TwitchStreamKey__Output | undefined
    twitchCredentials: DevCredentials__Output | undefined
    streamParams: StreamParametersResponse__Output | undefined
}

const Settings = ({ statusStreamKey, twitchCredentials, streamParams }: Props) => {

    return (
        <>
            <div className='w-1/4 flex flex-col items-start'>
                <div className='w-full h-full bg-foreground rounded-b-xl flex justify-center items-center'>
                    <div className="w-[98%] h-[98%] text-white">
                        <h2 className="text-center opacity-85 text-xl font-semibold uppercase tracking-wider">
                            Settings
                        </h2>

                        <div className="w-[98%] h-1 mx-auto my-1 bg-primary"></div>

                        <AddSongs />
                        <TwitchSettings statusStreamKey={statusStreamKey} twitchCredentials={twitchCredentials} />
                        <StreamParams streamParams={streamParams} />
                    </div>
                </div>
            </div>
        </>
    )
}

export default Settings