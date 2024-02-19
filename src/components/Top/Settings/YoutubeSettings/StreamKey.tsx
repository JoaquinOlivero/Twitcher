import { checkYoutubeParams, deleteYoutubeStreamKey, saveYoutubeStreamKey } from "@/actions"
import ActionButton from "@/components/Utils/ActionButton"
import DeleteButton from "@/components/Utils/DeleteButton"
import { Dispatch, SetStateAction, useState } from "react"

type Props = {
    isStreamKey: boolean
    setIsStreamKey: Dispatch<SetStateAction<boolean>>
}

const StreamKey = ({ isStreamKey, setIsStreamKey }: Props) => {
    const [streamKey, setStreamKey] = useState<string>(isStreamKey ? "000000000000000000000000000000000000000000000000000000" : "")
    const [isWaiting, setIsWaiting] = useState<boolean>(false)

    const handleSaveStreamKey = async () => {
        if (streamKey.length === 0) {
            return
        }

        setIsWaiting(true)

        const success = await saveYoutubeStreamKey(streamKey)
        if (success) {
            const status = await checkYoutubeParams()
            if (status && status.isKeyActive) {
                setIsStreamKey(true)
            }
        }

        setIsWaiting(false)
    }

    const handleDeleteStreamKey = async () => {
        await deleteYoutubeStreamKey()

        const status = await checkYoutubeParams()
        if (status && status.isKeyActive) {
            setStreamKey("000000000000000000000000000000000000000000000000000000")
            setIsStreamKey(true)
        } else {
            setStreamKey("")
            setIsStreamKey(false)
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
                    width="1/5"
                    disabled={isStreamKey}
                    isWaiting={isWaiting}
                    onClick={handleSaveStreamKey}
                    backgroundColor="bg-purple-800"
                    backgroundColorHover="bg-purple-600"
                />
            </div>
        </form>

    )
}

export default StreamKey