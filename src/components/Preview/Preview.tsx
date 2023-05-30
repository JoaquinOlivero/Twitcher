'use client';
import { useEffect, useState } from "react";
import ReactPlayer from "react-player"

const Preview = () => {
    const [isLoaded, setIsLoaded] = useState<boolean>(false)
    const [muted, setMuted] = useState<boolean>(true)
    const [volume, setVolume] = useState<number>(0.1)

    // useEffect(() => {
    //     if (isLoaded === false) {
    //         setIsLoaded(true)
    //     }
    // }, [isLoaded])

    // const handleOnError = (error: any) => {
    //     console.log(error)
    //     setIsLoaded(false)
    // }

    return (
        <div className="w-1/2 h-full mx-auto relative">
            <div className="bg-foreground w-full h-full z-0 rounded-b-xl"></div>

            {/* Controls */}
            < div className="absolute z-10 w-[98%] h-[98%] top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 text-white flex items-end">
                {isLoaded &&
                    <button onClick={() => setMuted(!muted)} className="z-3">
                        {muted ?
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-5 h-5">
                                <path strokeLinecap="round" strokeLinejoin="round" d="M17.25 9.75L19.5 12m0 0l2.25 2.25M19.5 12l2.25-2.25M19.5 12l-2.25 2.25m-10.5-6l4.72-4.72a.75.75 0 011.28.531V19.94a.75.75 0 01-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.506-1.938-1.354A9.01 9.01 0 012.25 12c0-.83.112-1.633.322-2.395C2.806 8.757 3.63 8.25 4.51 8.25H6.75z" />
                            </svg>
                            :
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-5 h-5">
                                <path strokeLinecap="round" strokeLinejoin="round" d="M19.114 5.636a9 9 0 010 12.728M16.463 8.288a5.25 5.25 0 010 7.424M6.75 8.25l4.72-4.72a.75.75 0 011.28.53v15.88a.75.75 0 01-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.507-1.938-1.354A9.01 9.01 0 012.25 12c0-.83.112-1.633.322-2.396C2.806 8.756 3.63 8.25 4.51 8.25H6.75z" />
                            </svg>
                        }
                    </button>
                }
            </div>

            {/* {
                isLoaded ?
                    <ReactPlayer
                        // url="/api/preview/master.m3u8"
                        url="/preview"
                        playing={true}
                        controls={false}
                        volume={volume}
                        muted={muted}
                        width="100%"
                        height="100%"
                        // onError={(error) => handleOnError(error)}
                        // config={{
                        //     file: {
                        //         forceSafariHLS: true,
                        //         hlsOptions: {
                        //             "lowLatencyMode": true
                        //         }
                        //     }
                        // }}

                        style={{ position: "absolute", top: "0" }}
                    />
                    :
                    null
            } */}

            <div id="remoteVideos" className="absolute top-0 w-full h-full">

            </div>

        </div >
    )
}

export default Preview