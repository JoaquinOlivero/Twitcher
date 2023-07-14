import { DevCredentials__Output } from '@/pb/service/DevCredentials'
import AddSongs from './AddSongs/AddSongs'
import TwitchSettings from './TwitchSettings/TwitchSettings'
import { TwitchStreamKey__Output } from '@/pb/service/TwitchStreamKey'

type Props = {
    statusStreamKey: TwitchStreamKey__Output | undefined
    twitchCredentials: DevCredentials__Output | undefined
}

const Settings = ({ statusStreamKey, twitchCredentials }: Props) => {
    return (
        <>
            <div className='w-1/4 flex flex-col items-start'>
                <div className='w-[95%] h-full bg-foreground rounded-b-xl flex justify-center items-center'>
                    <div className="w-[98%] h-[98%] text-white">
                        <h2 className="text-center opacity-85 text-xl font-semibold uppercase tracking-wider">
                            Settings
                        </h2>

                        <div className="w-[98%] h-1 mx-auto my-1 bg-background"></div>

                        <AddSongs />
                        <TwitchSettings statusStreamKey={statusStreamKey} twitchCredentials={twitchCredentials} />
                    </div>
                </div>
            </div>
        </>
    )
}

export default Settings