import path from 'path'
import * as grpc from "@grpc/grpc-js"
import * as protoLoader from "@grpc/proto-loader"
import { ProtoGrpcType } from "../pb/songs"
import { Song__Output } from '@/pb/service/Song'
import Preview from '@/components/Preview/Preview'
import Playlist from '@/components/Playlist/Playlist'

const PROTO_FILE = "../../../../proto/songs.proto"

const packageDef = protoLoader.loadSync(path.resolve(__dirname, PROTO_FILE))
const grpcObj = (grpc.loadPackageDefinition(packageDef) as unknown) as ProtoGrpcType


const getData = async () => {

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
                    // console.log(err)
                    resolve(undefined)
                }
                if (res?.songs !== undefined) {
                    resolve(res.songs)
                }
            })

        })
    })

    return playlist
}

const Home = async () => {
    const data: Song__Output[] | undefined = await getData()

    return (
        <div className='w-[99%] h-screen mx-auto flex flex-col'>

            <div className='w-full h-3/5'>
                <Preview />
            </div>

            <div className='w-full h-2/5 flex items-center'>

                <div className='w-1/4 h-[95%]'>
                    <div className='rounded-t-xl bg-foreground w-[95%] h-1/2'>
                        Video Source
                    </div>

                    <div className='bg-foreground w-[95%] h-1/2'>
                        Layouts
                    </div>
                </div>

                <Playlist songs={data} />

                <div className='w-1/4 h-[95%] flex flex-col items-end'>
                    <div className='rounded-t-xl bg-foreground w-[95%] h-full'>
                        Search Song
                    </div>
                </div>

            </div>

        </div>
    )
}

export default Home
