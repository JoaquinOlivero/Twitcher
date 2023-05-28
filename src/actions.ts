"use server";

import path from 'path'
import * as grpc from "@grpc/grpc-js"
import * as protoLoader from "@grpc/proto-loader"
import { ProtoGrpcType } from "@/pb/songs"
import { Song__Output } from '@/pb/service/Song'
import { revalidatePath } from 'next/cache'
import { SongPlaylist__Output } from './pb/service/SongPlaylist';

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

export const startPreview = async () => {
    
}