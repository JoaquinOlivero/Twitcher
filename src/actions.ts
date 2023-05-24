"use server";

import path from 'path'
import * as grpc from "@grpc/grpc-js"
import * as protoLoader from "@grpc/proto-loader"
import { ProtoGrpcType } from "@/pb/songs"
import { Song__Output } from '@/pb/service/Song'
import { revalidatePath } from 'next/cache'

const PROTO_FILE = "../../../../proto/songs.proto"

const packageDef = protoLoader.loadSync(path.resolve(__dirname, PROTO_FILE))
const grpcObj = (grpc.loadPackageDefinition(packageDef) as unknown) as ProtoGrpcType


export const createNewPlaylist = async () => {

    const client = new grpcObj.service.SongsManagement(
        `0.0.0.0:9000`, grpc.credentials.createInsecure()
    )

    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const playlist: Song__Output[] | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
            }

            client.CreatePlaylist({}, (err, res) => {
                if (err) {
                    console.log(err)
                    resolve(undefined)
                }
                if (res?.songs !== undefined) {
                    resolve(res.songs)
                }
            })

        })
    })

    console.log(playlist)
    revalidatePath("/")
}