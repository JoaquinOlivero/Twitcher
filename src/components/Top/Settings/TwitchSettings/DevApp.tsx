import { checkTwitchCredentials, deleteTwitchDevCredentials, saveTwitchCredentials } from "@/actions"
import ActionButton from "@/components/Utils/ActionButton"
import DeleteButton from "@/components/Utils/DeleteButton"
import { DevCredentials__Output } from "@/pb/service/DevCredentials"
import { useState } from "react"

type Props = {
    twitchCredentials: DevCredentials__Output | undefined
}


const DevApp = ({ twitchCredentials }: Props) => {
    const [clientId, setClientId] = useState<string>(twitchCredentials && twitchCredentials?.clientId ? twitchCredentials.clientId : "")
    const [secret, setSecret] = useState<string>(twitchCredentials && twitchCredentials?.active ? "000000000000000000000000000000000000000000000000000000" : "")
    const [isActive, setIsActive] = useState<boolean | undefined>(twitchCredentials?.active)
    const [isWaiting, setIsWaiting] = useState<boolean>(false)
    const [connectError, setConnectError] = useState<string | null>(null)

    const handleAuthorizeAccount = async () => {
        if (clientId.length === 0 || secret.length === 0) {
            return
        }

        setIsWaiting(true)

        const originUrl = new URL(window.location.origin)
        if (originUrl.protocol === "http:" && !originUrl.host.includes("localhost")) {
            setConnectError("http: protocol only allowed in localhost.")
            setIsWaiting(false)
            return
        }

        const success = await saveTwitchCredentials(clientId, secret)
        if (success) {
            const authTwitchUrl = `https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=${clientId}&redirect_uri=${originUrl}twitch&scope=moderator%3Aread%3Afollowers+channel%3Aread%3Asubscriptions+user%3Aread%3Aemail`
            window.open(authTwitchUrl)

            const status = await checkTwitchCredentials()
            if (status && status.active) {
                setClientId(status.clientId ? status.clientId : "")
                setSecret(status.active ? "000000000000000000000000000000000000000000000000000000" : "")
                setIsActive(status.active)
            }
        }

        setIsWaiting(false)
    }

    const handleRemoveAuth = async () => {
        await deleteTwitchDevCredentials()

        const status = await checkTwitchCredentials()
        if (status) {
            setClientId("")
            setSecret("")
            setIsActive(false)
        }
    }

    return (
        <div className='flex flex-col items-center w-full h-full'>
            <span className="text-center opacity-85 font-semibold uppercase tracking-wider">developer app</span>

            <form autoComplete="off" className='relative w-4/5 flex flex-col h-full' onSubmit={(e) => e.preventDefault()}>
                <div className={`flex flex-col gap-1 mb-3 ${isActive && 'pointer-events-none opacity-40'}`}>
                    <label htmlFor="client_id" className="font-semibold text-white opacity-85 capitalize">Client Id</label>
                    <input autoComplete="off" type="text" id="client_id" className="bg-background font-semibold tracking-wider text-white rounded w-full p-2.5 outline-none" required value={clientId} onChange={(e) => setClientId(e.target.value)} />
                </div>
                <div className={`flex flex-col gap-1 mb-3 ${isActive && 'pointer-events-none opacity-40'}`}>
                    <label htmlFor="client_secret" className="font-semibold text-white opacity-85 capitalize">Client Secret</label>
                    <input autoComplete="new-password" type="text" id="client_secret" className="bg-background tracking-wider text-white rounded w-full p-2.5 outline-none font-dots" required value={secret} onChange={(e) => setSecret(e.target.value)} />
                </div>

                <div className='flex justify-end gap-2'>
                    {isActive &&
                        <DeleteButton deleteFunc={handleRemoveAuth} />
                    }
                    <ActionButton
                        text="Authorize account"
                        width="1/3"
                        disabled={isActive}
                        isWaiting={isWaiting}
                        onClick={handleAuthorizeAccount}
                        backgroundColor="bg-purple-800"
                        backgroundColorHover="bg-purple-600"
                    />

                </div>

                {connectError && <span className='text-sm font-semibold text-red-600'>{connectError}</span>}
            </form>
        </div>
    )
}

export default DevApp