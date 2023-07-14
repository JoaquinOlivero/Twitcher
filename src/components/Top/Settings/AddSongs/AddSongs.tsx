import Modal from "@/components/Modal/Modal"
import { useState } from "react"
import NCS from "./NCS"
import SoundCloud from "./Soundcloud"
import UploadFile from "./UploadFile"

const AddSongs = () => {
    const [tWidth, setTWidth] = useState<string>("w-1/3")
    const [tHeight, setTHeight] = useState<string>("h-1/4")
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [ncs, setNcs] = useState<boolean>(false)
    const [soundCloud, setSoundCloud] = useState<boolean>(false)
    const [uploadFile, setUploadFile] = useState<boolean>(false)
    const [isSubPage, setIsSubPage] = useState<boolean>(false)

    const showNcs = () => {
        setNcs(true)
        setIsSubPage(true)
        setTHeight("h-1/3")
    }

    const showSoundCloud = () => {
        setSoundCloud(true)
        setIsSubPage(true)
        setTHeight("h-1/3")
    }

    const showUploadFile = () => {
        setUploadFile(true)
        setIsSubPage(true)
        setTHeight("h-2/3")
    }


    const goBack = () => {
        setNcs(false)
        setSoundCloud(false)
        setUploadFile(false)
        setIsSubPage(false)
        setTWidth("w-1/3")
        setTHeight("h-1/4")
    }

    return (
        <>
            <Modal tWidth={tWidth} tHeight={tHeight} isOpen={isOpen} setIsOpen={setIsOpen} goBack={goBack} isSubPage={isSubPage}>
                <div className="flex flex-col h-full">
                    {!ncs && !soundCloud && !uploadFile &&
                        <>
                            <div className="text-center opacity-85 text-xl font-semibold uppercase tracking-wider">add songs</div>
                            <div className="flex h-full w-full gap-5 items-center justify-center">
                                <div onClick={() => showNcs()} className="text-center bg-blue-600/30 w-1/5 py-1.5 rounded-lg font-semibold tracking-wider capitalize cursor-pointer transition hover:bg-blue-600/60">
                                    <span>NCS</span>
                                </div>
                                <div onClick={() => showSoundCloud()} className="text-center bg-blue-600/30 w-1/5 py-1.5 rounded-lg font-semibold tracking-wider capitalize cursor-pointer transition hover:bg-blue-600/60">
                                    <span>SoundCloud</span>
                                </div>
                                <div onClick={() => showUploadFile()} className="text-center bg-blue-600/30 w-1/5 py-1.5 rounded-lg font-semibold tracking-wider capitalize cursor-pointer transition hover:bg-blue-600/60">
                                    <span>Upload file</span>
                                </div>
                            </div>
                        </>
                    }

                    {ncs && !soundCloud && !uploadFile &&
                        <NCS />
                    }

                    {soundCloud && !ncs && !uploadFile &&
                        <SoundCloud />
                    }

                    {uploadFile && !ncs && !soundCloud &&
                        <UploadFile />
                    }


                </div>
            </Modal>

            <div className="py-1 my-1 transition text-sm capitalize font-semibold tracking-wider opacity-80 hover:bg-background hover:opacity-100 cursor-pointer" onClick={() => setIsOpen(true)}>
                <span className="ml-2">Add songs</span>
            </div>
        </>
    )
}

export default AddSongs
