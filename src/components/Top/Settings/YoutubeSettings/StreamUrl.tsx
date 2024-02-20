import { checkYoutubeParams, deleteYoutubeStreamUrl, saveYoutubeStreamUrl } from "@/actions"
import ActionButton from "@/components/Utils/ActionButton"
import DeleteButton from "@/components/Utils/DeleteButton"
import { Dispatch, SetStateAction, useState } from "react"

type Props = {
    streamUrl: string
    setStreamUrl: Dispatch<SetStateAction<string>>
    setIsEnabled: Dispatch<SetStateAction<boolean>>
}

const StreamUrl = ({ streamUrl, setStreamUrl, setIsEnabled }: Props) => {
    const [isStreamUrl, setIsStreamUrl] = useState<boolean | undefined>(streamUrl && streamUrl?.length > 0 ? true : false)
    const [isWaiting, setIsWaiting] = useState<boolean>(false)

    const handleSaveStreamUrl = async () => {
        if (streamUrl.length === 0) {
            return
        }

        setIsWaiting(true)

        const success = await saveYoutubeStreamUrl(streamUrl)
        if (success) {
            const status = await checkYoutubeParams()
            if (status && status.url !== undefined) {
                setIsStreamUrl(true)
            }
        }

        setIsWaiting(false)
    }

    const handleDeleteStreamUrl = async () => {
        await deleteYoutubeStreamUrl()

        const status = await checkYoutubeParams()
        if (status && status.url == undefined) {
            setStreamUrl("")
            setIsStreamUrl(false)
            setIsEnabled(false)
        } else if (status && status.url) {
            setStreamUrl(status.url)
            setIsStreamUrl(true)
        }
    }

    return (
        <form autoComplete="off" className={`w-4/5 flex flex-col`} onSubmit={(e) => e.preventDefault()}>
            <div className={`flex flex-col gap-1 mb-3 ${isStreamUrl && 'pointer-events-none opacity-40'}`}>
                <label htmlFor="key" className="font-semibold text-white opacity-85 capitalize">Stream URL</label>
                <input autoComplete="new-password" type="text" id="key" className="bg-background tracking-wider text-white rounded w-full p-2.5 outline-none font-semibold" required value={streamUrl} onChange={(e) => setStreamUrl(e.currentTarget.value)} />
            </div>
            <div className='flex justify-end gap-2'>
                {isStreamUrl &&
                    <DeleteButton
                        deleteFunc={handleDeleteStreamUrl}
                    />
                }
                <ActionButton
                    text="save"
                    width="1/5"
                    disabled={isStreamUrl}
                    isWaiting={isWaiting}
                    onClick={handleSaveStreamUrl}
                    backgroundColor="bg-purple-800"
                    backgroundColorHover="bg-purple-600"
                />
            </div>
        </form>

    )
}

export default StreamUrl