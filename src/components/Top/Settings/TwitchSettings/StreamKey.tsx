import { checkTwitchStreamKey, deleteTwitchStreamKey, saveTwitchStreamKey } from "@/actions"
import ActionButton from "@/components/Utils/ActionButton"
import DeleteButton from "@/components/Utils/DeleteButton"
import { TwitchStreamKey__Output } from "@/pb/service/TwitchStreamKey"
import { useState } from "react"

type Props = {
    statusStreamKey: TwitchStreamKey__Output | undefined
}

const StreamKey = ({ statusStreamKey }: Props) => {
    const [streamKey, setStreamKey] = useState<string>(statusStreamKey && statusStreamKey?.active ? "000000000000000000000000000000000000000000000000000000" : "")
    const [isStreamKey, setIsStreamKey] = useState<boolean | undefined>(statusStreamKey?.active)

    const handleSaveStreamKey = async () => {
        if (streamKey.length === 0) {
            return
        }

        const success = await saveTwitchStreamKey(streamKey)
        if (success) {
            const status = await checkTwitchStreamKey()
            if (status && status.active) {
                setIsStreamKey(status.active)
            }
        }
    }

    const handleDeleteStreamKey = async () => {
        await deleteTwitchStreamKey()

        const status = await checkTwitchStreamKey()
        if (status) {
            setStreamKey("")
            setIsStreamKey(status.active)
        }
    }

    return (
        <form autoComplete="off" className={`w-4/5 flex flex-col`} onSubmit={(e) => e.preventDefault()}>
            <div className={`flex flex-col gap-1 mb-3 ${isStreamKey && 'pointer-events-none opacity-40'}`}>
                <label htmlFor="key" className="font-semibold text-white opacity-85 capitalize">Stream key</label>
                <input autoComplete="new-password" type="text" id="key" className="bg-background tracking-wider text-white rounded w-full p-2.5 outline-none font-dots text-xs" required value={streamKey} onChange={(e) => setStreamKey(e.currentTarget.value)} />
            </div>
            <div className='flex justify-end gap-2'>
                {isStreamKey &&
                    <DeleteButton
                        deleteFunc={handleDeleteStreamKey}
                    />
                }
                <ActionButton
                    text="save"
                    width="1/6"
                    toggle={isStreamKey}
                    onClick={handleSaveStreamKey}
                    backgroundColor="bg-purple-800"
                    backgroundColorHover="bg-purple-600"
                />
            </div>
        </form>

    )
}

export default StreamKey