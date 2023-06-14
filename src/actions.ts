"use server";

import path, { resolve } from 'path'
import * as grpc from "@grpc/grpc-js"
import * as protoLoader from "@grpc/proto-loader"
import { ProtoGrpcType } from "@/pb/songs"
import { SongPlaylist__Output } from './pb/service/SongPlaylist';
import { AudioStream__Output } from './pb/service/AudioStream';
import { OutputResponse__Output } from './pb/service/OutputResponse';
import { revalidatePath } from 'next/cache';
import { SDP__Output } from './pb/service/SDP';

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
                console.log(err)
                resolve(undefined)
            }

            client.CurrentSongPlaylist({}, (err, res) => {
                if (err) {
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

    return playlist
}

export const createNewPlaylist = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const playlist: SongPlaylist__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
            }

            client.CreateSongPlaylist({}, (err, res) => {
                if (err) {
                    console.log(err)
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
            console.log(err)
            return
        }

        client.UpdateSongPlaylist(songs, (err, res) => {
            if (err) {
                console.log(err)
                return
            }
        })
    })

}

export const enablePreview = async (clientSdp: string) => {

    await startAudio()

    await startOutput()

    const serverSdp: string = await new Promise<string>(resolve => {
        var call = client.Preview({sdp: clientSdp});
        call.on("data", async (res: SDP__Output) => {
            if (res.sdp) {
                resolve(res.sdp)
            }
        })
    
        call.on("end", () => {
            // The server has finished sending data.
        })
    
        call.on("error", (err) => {
            // An error has occurred and the stream is closed.
            console.log(err)
            return
        })
    
        call.on("status", (status) => {
            // process status
            // console.log(status)
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
            console.log(err)
            return
        }

        client.StartTwitch({}, (err, res) => {
            if (err) {
                console.log(err)
                return
            }
            console.log(res)
        })
    })

}

const startAudio = async () => {
    // Listen to audio grpc stream.
    const audio: boolean = await new Promise(resolve => {
        var call = client.Audio({});
        call.on("data", async (res: AudioStream__Output) => {
            if (res.playlist) {
                // console.log(res)
            }

            if (res.ready) {
                resolve(res.ready)
            }
        })
    
        call.on("end", () => {
            // The server has finished sending data.
        })
    
        call.on("error", (err) => {
            // An error has occurred and the stream is closed.
            console.log(err)
            resolve(false)
        })
    
        call.on("status", (status) => {
            // process status
            // console.log(status)
        })
    
    })

    return audio
}

const startOutput = async () => {
    const output: boolean = await new Promise(resolve => {
        var call = client.Output({});
        call.on("data", async (res: OutputResponse__Output) => {
            if (res.time) {
                // console.log("time: ", res.time)
            }

            if (res.bitrate) {
                // console.log("bitrate: ", res.bitrate)
            }

            if (res.ready) {
                resolve(res.ready)
            }

        })
    
        call.on("end", () => {
            // The server has finished sending data.
        })
    
        call.on("error", (err) => {
            // An error has occurred and the stream is closed.
            console.log(err)
            resolve(false)

        })
    
        call.on("status", (status) => {
            // process status
            // console.log(status)
        })
    })

    return output

}