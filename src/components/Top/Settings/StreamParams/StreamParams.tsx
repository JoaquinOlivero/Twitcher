import { saveStreamParams } from "@/actions"
import Modal from "@/components/Modal/Modal"
import ActionButton from "@/components/Utils/ActionButton"
import { SaveStreamParametersRequest } from "@/pb/service/SaveStreamParametersRequest"
import { StreamParametersResponse__Output } from "@/pb/service/StreamParametersResponse"
import { useState } from "react"

type Props = {
    streamParams: StreamParametersResponse__Output | undefined
}

const StreamParams = ({ streamParams }: Props) => {
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [params, setParams] = useState<StreamParametersResponse__Output | undefined>(streamParams)
    const [isSaving, setIsSaving] = useState<boolean>(false)

    const handleSaveParameters = async () => {
        setIsSaving(true)
        const saveParams: SaveStreamParametersRequest = {
            width: params?.width,
            height: params?.height,
            fps: params?.fps,
            preset: params?.preset
        }
        await saveStreamParams(saveParams)
        setIsSaving(false)
    }

    const handleChangeResolution = async (height: number) => {
        switch (height) {
            case 1080:
                setParams({ height: height, width: 1920, fps: params?.fps, preset: params?.preset })
                break;
            case 720:
                setParams({ height: height, width: 1280, fps: params?.fps, preset: params?.preset })
                break;
            case 480:
                setParams({ height: height, width: 640, fps: params?.fps, preset: params?.preset })
                break;
        }
    }

    return (
        <>
            <Modal tWidth="w-1/3" tHeight="h-2/4" isOpen={isOpen} setIsOpen={setIsOpen}>
                <div className="text-center opacity-85 text-xl font-semibold uppercase tracking-wider">Stream parameters</div>

                <div className="flex flex-col h-full w-full gap-5 mt-5 items-center">
                    <form autoComplete="off" className={`w-4/5 h-full flex flex-col gap-5`} onSubmit={(e) => e.preventDefault()}>
                        <div className="flex flex-col gap-1 items-center">
                            <label className="font-semibold text-white opacity-85 uppercase">Resolution</label>
                            <select value={params?.height} className="bg-background tracking-wider text-white rounded p-2.5 outline-none font-semibold"
                                onChange={(e) => handleChangeResolution(parseInt(e.currentTarget.value))}
                            >
                                <option value={1080} className="font-semibold">1920x1080</option>
                                <option value={720} className="font-semibold">1280x720</option>
                                <option value={480} className="font-semibold">640x480</option>
                            </select>
                        </div>

                        <div className="flex flex-col gap-1 items-center">
                            <label className="font-semibold text-white opacity-85 uppercase">fps</label>
                            <select value={params?.fps} className="bg-background tracking-wider text-white rounded p-2.5 outline-none font-semibold"
                                onChange={(e) => setParams({ height: params?.height, width: params?.width, fps: parseInt(e.currentTarget.value), preset: params?.preset })}
                            >
                                <option value={60} className="font-semibold">60</option>
                                <option value={59} className="font-semibold">59</option>
                                <option value={30} className="font-semibold">30</option>
                                <option value={25} className="font-semibold">25</option>
                                <option value={24} className="font-semibold">24</option>
                            </select>
                        </div>

                        <div className="flex flex-col gap-1 items-center">
                            <label className="font-semibold text-white opacity-85 uppercase">preset</label>
                            <select value={params?.preset} className="bg-background tracking-wider text-white rounded p-2.5 outline-none font-semibold"
                                onChange={(e) => setParams({ height: params?.height, width: params?.width, fps: params?.fps, preset: e.currentTarget.value })}
                            >
                                <option value="ultrafast" className="font-semibold">ultrafast</option>
                                <option value="superfast" className="font-semibold">superfast</option>
                                <option value="veryfast" className="font-semibold">veryfast</option>
                                <option value="faster" className="font-semibold">faster</option>
                                <option value="fast" className="font-semibold">fast</option>
                                <option value="medium" className="font-semibold">medium</option>
                                <option value="slow" className="font-semibold">slow</option>
                                <option value="slower" className="font-semibold">slower</option>
                                <option value="veryslow" className="font-semibold">veryslow</option>
                            </select>
                        </div>

                        <div className="w-full flex justify-center mt-10">
                            <ActionButton
                                text="save"
                                width="1/3"
                                disabled={isSaving}
                                isWaiting={isSaving}
                                onClick={handleSaveParameters}
                                backgroundColor="bg-primary"
                                backgroundColorHover="bg-purple-600"
                            />
                        </div>
                    </form>
                </div>
            </Modal>

            <div className="py-1 my-1 transition text-sm capitalize font-semibold tracking-wider opacity-80 hover:bg-background hover:opacity-100 cursor-pointer" onClick={() => setIsOpen(true)}>
                <span className="ml-2">Stream parameters</span>
            </div>
        </>
    )
}

export default StreamParams