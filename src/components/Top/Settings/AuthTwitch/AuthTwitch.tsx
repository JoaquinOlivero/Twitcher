import Modal from '@/components/Modal/Modal'
import React, { useState } from 'react'

const AuthTwitch = () => {
    const [isOpen, setIsOpen] = useState<boolean>(false)

    return (
        <>
            {/* <Modal tWidth="w-1/3" tHeight="h-2/3" isOpen={isOpen} setIsOpen={setIsOpen}>
                <div>Auth twitch</div>
            </Modal> */}
            <div className="py-1 my-1 transition text-sm capitalize font-semibold tracking-wider opacity-80 hover:bg-background hover:opacity-100 cursor-pointer" onClick={() => setIsOpen(true)}>
                <span className="ml-2">Authorize Twitch</span>
            </div>
        </>
    )
}

export default AuthTwitch