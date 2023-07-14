'use client';

import { twitchAccessToken } from '@/actions';
import Link from 'next/link';
import { useSearchParams } from 'next/navigation'
import { useEffect, useState } from 'react';

const Twitch = () => {
    const searchParams = useSearchParams()
    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [code, setCode] = useState<boolean>(false)

    const handleAccessToken = async (code: string) => {
        const success = await twitchAccessToken(code)
        console.log(success)
        window.close()
    }

    useEffect(() => {
        const code = searchParams.get("code")
        if (!code) {
            setCode(false)
            setIsLoading(false)
        } else {
            // send code to grpc endpoint and wait for confirmation.
            setCode(true)
            handleAccessToken(code)
        }

    }, [])


    return (
        <>
            {!code && !isLoading &&
                <div className='flex h-full justify-center items-center text-white font-semibold text-lg'><Link href="/">Return to home</Link></div>
            }

            {code && isLoading &&
                <div className='flex h-full justify-center items-center text-white font-semibold text-lg'>
                    Getting access token. Please wait...
                </div>
            }
        </>
    )

}

export default Twitch