import Modal from '@/components/Modal/Modal'
import React, { useState } from 'react'
import StreamKey from './StreamKey'
import DevApp from './DevApp'
import { TwitchStreamKey__Output } from '@/pb/service/TwitchStreamKey'
import { DevCredentials__Output } from '@/pb/service/DevCredentials'

type Props = {
    statusStreamKey: TwitchStreamKey__Output | undefined
    twitchCredentials: DevCredentials__Output | undefined
}

const TwitchSettings = ({ statusStreamKey, twitchCredentials }: Props) => {
    const [isOpen, setIsOpen] = useState<boolean>(false)

    return (

        <>
            <Modal tWidth="w-1/3" tHeight="h-2/4" isOpen={isOpen} setIsOpen={setIsOpen}>
                <div className="text-center opacity-85 text-xl font-semibold uppercase tracking-wider">Twitch</div>

                <div className="flex flex-col h-full w-full gap-5 mt-5 items-center">
                    <StreamKey statusStreamKey={statusStreamKey} />
                    <DevApp twitchCredentials={twitchCredentials} />
                </div>
            </Modal>

            <div className="py-1 my-1 transition text-sm capitalize font-semibold tracking-wider opacity-80 hover:bg-background hover:opacity-100 cursor-pointer" onClick={() => setIsOpen(true)}>
                <span className="ml-2">Twitch</span>
            </div>
        </>
    )
}

export default TwitchSettings