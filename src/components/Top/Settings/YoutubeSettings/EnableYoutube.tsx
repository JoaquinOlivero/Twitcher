import { manageYoutube } from "@/actions"
import { Dispatch, SetStateAction, useState } from "react"

type Props = {
    isStreamKey: boolean
    isEnabled: boolean
    setIsEnabled: Dispatch<SetStateAction<boolean>>
    streamUrl: string
}

const EnableYoutube = ({ isStreamKey, isEnabled, setIsEnabled, streamUrl }: Props) => {
    const [isWaiting, setIsWaiting] = useState<boolean>(false)


    const handleEnableYoutube = async () => {
        setIsWaiting(true)

        setIsEnabled(!isEnabled)
        await manageYoutube(!isEnabled)

        setIsWaiting(false)
    }

    return (
        <label className={`w-4/5 mb-2 inline-flex items-center cursor-pointer ${!isStreamKey || isEnabled == undefined || streamUrl.length < 1 || isWaiting ? "opacity-50 cursor-not-allowed" : "opacity-100"}`} >
            <input type="checkbox" value="" className="sr-only peer" disabled={!isStreamKey || isEnabled == undefined || streamUrl.length < 1 || isWaiting} checked={isEnabled} onClick={() => handleEnableYoutube()} />
            <div className="
                relative
                w-9 h-5
                bg-gray-200
                peer-focus:outline-none
                rounded-full peer
                dark:bg-gray-700
                peer-checked:after:translate-x-full
                rtl:peer-checked:after:-translate-x-full
                peer-checked:after:border-white after:content-['']
                after:absolute after:top-[2px]
                after:start-[2px] after:bg-white
                after:border-gray-300 after:border
                after:rounded-full
                after:h-4 after:w-4
                after:transition-all
                peer-checked:bg-purple-600
                "></div>
            <span className="ms-2 text-sm font-semibold text-gray-900 dark:text-gray-300">{isEnabled ? "Enabled" : "Enable"}</span>
        </label>
    )
}

export default EnableYoutube