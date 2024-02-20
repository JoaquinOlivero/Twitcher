import Modal from "@/components/Modal/Modal"
import { useState } from "react"
import StreamKey from "./StreamKey"
import StreamUrl from "./StreamUrl"
import { YoutubeParams__Output } from "@/pb/service/YoutubeParams"
import EnableYoutube from "./EnableYoutube"

type Props = {
    params: YoutubeParams__Output | undefined
}

function YoutubeSettings({ params }: Props) {
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [isStreamKey, setIsStreamKey] = useState<boolean>(params && params.isKeyActive ? params.isKeyActive : false)
    const [streamUrl, setStreamUrl] = useState<string>(params && params.url && params.url.length > 0 ? params.url : "")
    const [isEnabled, setIsEnabled] = useState<boolean>(params && params.enabled ? params.enabled : false)

    return (
        <>
            <Modal tWidth="w-1/3" tHeight="h-auto" isOpen={isOpen} setIsOpen={setIsOpen}>
                <div className="text-center opacity-85 text-xl font-semibold uppercase tracking-wider">Youtube</div>

                <div className="flex flex-col h-full w-full gap-5 mt-5 items-center">
                    <StreamKey isStreamKey={isStreamKey} setIsStreamKey={setIsStreamKey} setIsEnabled={setIsEnabled} />
                    <StreamUrl streamUrl={streamUrl} setStreamUrl={setStreamUrl} setIsEnabled={setIsEnabled} />
                    <EnableYoutube isStreamKey={isStreamKey} streamUrl={streamUrl} isEnabled={isEnabled} setIsEnabled={setIsEnabled} />
                </div>
            </Modal>

            <div className="py-1 my-1 transition text-sm capitalize font-semibold tracking-wider opacity-80 hover:bg-background hover:opacity-100 cursor-pointer" onClick={() => setIsOpen(true)}>
                <span className="ml-2">Youtube</span>
            </div>
        </>
    )
}

export default YoutubeSettings