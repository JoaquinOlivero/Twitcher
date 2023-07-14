import { findNewSongsNCS, statusNCS } from "@/actions"
import { useEffect, useState } from "react"

const NCS = () => {
    const [isFinding, setIsFinding] = useState<boolean>(false)

    const findNewSongs = async (find: boolean, checkStatus: boolean) => {
        if (find && !checkStatus) {
            setIsFinding(true)
            findNewSongsNCS()
        }

        const status = await statusNCS()

        if (status && status.active) {
            setIsFinding(status.active)
            await sleep(150)
            findNewSongs(false, true)
            return
        }

        setIsFinding(false)
    }

    useEffect(() => {
        findNewSongs(false, true)
    }, [])


    return (
        <>
            <div className="text-center opacity-85 text-xl font-semibold uppercase tracking-wider">ncs</div>
            <div className="relative flex h-full w-full gap-5 items-center justify-center">
                <button onClick={() => findNewSongs(true, false)} className={`text-center bg-blue-600/30 px-2 py-1.5 rounded-lg font-semibold tracking-wider capitalize cursor-pointer transition hover:bg-blue-600/60 flex items-center gap-1 ${isFinding && "pointer-events-none opacity-40"}`} disabled={isFinding}>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor" className="w-5 h-5">
                        <path strokeLinecap="round" strokeLinejoin="round" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
                    </svg>
                    {!isFinding ?
                        <span>find new songs</span>
                        :
                        <span>finding new songs. Please wait...</span>
                    }
                </button>
                {isFinding &&
                    <span className="absolute bottom-1 text-sm text-center opacity-70 font-semibold">This may take a while. You can close this tab.</span>
                }
            </div>
        </>
    )
}

export default NCS

const sleep = (ms: number) => new Promise(res => setTimeout(res, ms));