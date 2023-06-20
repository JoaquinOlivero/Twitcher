"use server";

import path from 'path'
import * as grpc from "@grpc/grpc-js"
import * as protoLoader from "@grpc/proto-loader"
import { ProtoGrpcType } from "@/pb/songs"
import { SongPlaylist__Output } from './pb/service/SongPlaylist';
import { OutputResponse__Output } from './pb/service/OutputResponse';
import { revalidatePath } from 'next/cache';

const PROTO_FILE = "../../../../proto/songs.proto"

const packageDef = protoLoader.loadSync(path.resolve(__dirname, PROTO_FILE))
const grpcObj = (grpc.loadPackageDefinition(packageDef) as unknown) as ProtoGrpcType

const client = new grpcObj.service.StreamManagement(
    `0.0.0.0:9000`, grpc.credentials.createInsecure()
)

export const getCurrentPlaylist = async () => {

    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const playlist: SongPlaylist__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                // console.log(err)
                resolve(undefined)
            }

            client.CurrentSongPlaylist({}, (err, res) => {
                if (err) {
                    resolve(undefined)
                }

                if (res !== undefined) {
                    resolve(res)
                    revalidatePath("/")
                } else {
                    resolve(undefined)
                }

            })

        })
    })

    return playlist
}

export const createNewPlaylist = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const playlist: SongPlaylist__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                // console.log(err)
                resolve(undefined)
            }

            client.CreateSongPlaylist({}, (err, res) => {
                if (err) {
                    // console.log(err)
                    resolve(undefined)
                }
                
                if (res !== undefined) {
                    resolve(res)
                    revalidatePath("/")
                } else {
                    resolve(undefined)
                }
            })

        })
    })

    return playlist
}

export const updateSongPlaylist = async (songs: SongPlaylist__Output) => {
    
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    client.waitForReady(deadline, (err) => {
        if (err) {
            // console.log(err)
            return
        }

        client.UpdateSongPlaylist(songs, (err, res) => {
            if (err) {
                // console.log(err)
                return
            }
        })
    })

}

export const enablePreview = async (clientSdp: string) => {
    await startAudio()

    await startOutput()

    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)
    
    const serverSdp: string = await new Promise<string>(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                // console.log(err)
                resolve(err.message)
                return err.message
            }
            
            client.Preview({sdp: clientSdp}, (err, res) => {
                if (err) {
                    // console.log(err)
                    resolve(err.message)
                    return err.message
                }
                
                if (res && res.sdp) {
                    resolve(res.sdp)
                    return res.sdp
                }
            })
        })
        
    })


    return serverSdp
}


export const startStream = async () => {

    await startAudio()

    await startOutput()

    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    client.waitForReady(deadline, (err) => {
        if (err) {
            // console.log(err)
            return
        }

        client.StartTwitch({}, (err, res) => {
            if (err) {
                // console.log(err)
                return
            }
        })
    })

}

const startAudio = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const ready: boolean = await new Promise<boolean>(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                // console.log(err)
                resolve(false)
                return false
            }
            
            client.StartAudio({}, (err, res) => {
                if (err) {
                    // console.log(err)
                    resolve(false)
                    return false
                }
                
                if (res && res.ready) {
                    resolve(res.ready)
                    return res.ready
                }
            })
        })
        
    })

    return ready
}
    
const startOutput = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)
    
    const ready: boolean = await new Promise<boolean>(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                // console.log(err)
                resolve(false)
                return false
            }
            
            client.StartOutput({}, (err, res) => {
                if (err) {
                    // console.log(err)
                    resolve(false)
                    return
                }
                
                if (res && res.ready) {
                    resolve(res.ready)
                    return res.ready
                }
            })
        })
        
    })

    return ready
}

export const checkOutputStatus = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const status: OutputResponse__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                // console.log(err)
                resolve(undefined)
            }

            client.outputStatus({}, (err, res) => {
                if (err) {
                    // console.log(err)
                    resolve(undefined)
                }
                
                if (res !== undefined) {
                    resolve(res)
                } else {
                    resolve(undefined)
                }
            })

        })
    })

    return status
}

export const stopOutput = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    client.waitForReady(deadline, (err) => {
        if (err) {
            // console.log(err)
            return
        }

        client.StopOutput({}, (err, res) => {
            if (err) {
                // console.log(err)
                return
            }
        })
    })
}