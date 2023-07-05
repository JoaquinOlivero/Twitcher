import React from 'react'
import AddSongs from './AddSongs/AddSongs'
import AuthTwitch from './AuthTwitch/AuthTwitch'

const Settings = () => {
    return (
        <>
            <div className='w-1/4 flex flex-col items-start'>
                <div className='w-[95%] h-full bg-foreground rounded-b-xl flex justify-center items-center'>
                    <div className="w-[98%] h-[98%] text-white">
                        <h2 className="text-center opacity-85 text-xl font-semibold uppercase tracking-wider">
                            Settings
                        </h2>

                        <div className="w-[98%] h-1 mx-auto my-1 bg-background"></div>

                        <AddSongs />
                        <AuthTwitch />
                    </div>
                </div>
            </div>
        </>
    )
}

export default Settings