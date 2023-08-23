import { findNewSongsNCS, statusNCS } from "@/actions"
import ActionButton from "@/components/Utils/ActionButton"
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
            await findNewSongs(false, true)
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
                <ActionButton
                    text="find new songs"
                    waitingText="please wait"
                    width="1/3"
                    disabled={isFinding}
                    isWaiting={isFinding}
                    backgroundColor="bg-primary"
                    onClick={() => findNewSongs(true, false)}
                />
                {isFinding &&
                    <span className="absolute bottom-1 text-sm text-center opacity-70 font-semibold">This may take a while. You can close this tab.</span>
                }
            </div>
        </>
    )
}

export default NCS

const sleep = (ms: number) => new Promise(res => setTimeout(res, ms));