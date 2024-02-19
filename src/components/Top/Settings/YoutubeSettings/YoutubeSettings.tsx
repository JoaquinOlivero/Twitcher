import Modal from "@/components/Modal/Modal"
import { useState } from "react"
import StreamKey from "./StreamKey"
import StreamUrl from "./StreamUrl"
import { YoutubeParams__Output } from "@/pb/service/YoutubeParams"

type Props = {
    params: YoutubeParams__Output | undefined
}

function YoutubeSettings({ params }: Props) {
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [isStreamKey, setIsStreamKey] = useState<boolean>(params && params.isKeyActive ? params.isKeyActive : false)
    const [streamUrl, setStreamUrl] = useState<string>(params && params.url && params.url.length > 0 ? params.url : "")
    const [isEnabled, setIsEnabled] = useState<boolean>(params && params.enabled ? params.enabled : false)

    const handleEnableYoutube = () => {
    }

    return (
        <>
            <Modal tWidth="w-1/3" tHeight="h-2/5" isOpen={isOpen} setIsOpen={setIsOpen}>
                <div className="text-center opacity-85 text-xl font-semibold uppercase tracking-wider">Youtube</div>

                <div className="flex flex-col h-full w-full gap-5 mt-5 items-center">
                    <StreamKey isStreamKey={isStreamKey} setIsStreamKey={setIsStreamKey} />
                    <StreamUrl streamUrl={streamUrl} setStreamUrl={setStreamUrl} />
                    <label className={`inline-flex items-center cursor-pointer ${!isStreamKey || params == undefined || streamUrl.length < 1 ? "opacity-50 cursor-not-allowed" : "opacity-100"}`} onClick={handleEnableYoutube}>
                        <input type="checkbox" value="" className="sr-only peer" disabled={!isStreamKey || params == undefined || streamUrl.length < 1} defaultChecked={isEnabled} />
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
                        <span className="ms-3 text-sm font-medium text-gray-900 dark:text-gray-300">Enable</span>
                    </label>
                </div>
            </Modal>

            <div className="py-1 my-1 transition text-sm capitalize font-semibold tracking-wider opacity-80 hover:bg-background hover:opacity-100 cursor-pointer" onClick={() => setIsOpen(true)}>
                <span className="ml-2">Youtube</span>
            </div>
        </>
    )
}

export default YoutubeSettings