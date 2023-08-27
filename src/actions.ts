"use server";

import path from 'path'
import * as grpc from "@grpc/grpc-js"
import * as protoLoader from "@grpc/proto-loader"
import { ProtoGrpcType } from "@/pb/main"
import { SongPlaylist__Output } from './pb/service/SongPlaylist';
import { revalidatePath } from 'next/cache';
import { StatusResponse__Output } from './pb/service/StatusResponse';
import { StatusNCSResponse__Output } from './pb/service/StatusNCSResponse';
import { TwitchStreamKey__Output } from './pb/service/TwitchStreamKey';
import { DevCredentials__Output } from './pb/service/DevCredentials';
import { Overlays__Output } from './pb/service/Overlays';
import { BackgroundVideosResponse__Output } from './pb/service/BackgroundVideosResponse';
import { BackgroundVideo } from './pb/service/BackgroundVideo';
import { StreamResponse__Output } from './pb/service/StreamResponse';
import { StreamParametersResponse__Output } from './pb/service/StreamParametersResponse';
import { SaveStreamParametersRequest } from './pb/service/SaveStreamParametersRequest';

const PROTO_FILE = "../proto/main.proto"

const packageDef = protoLoader.loadSync(path.resolve(process.cwd(), PROTO_FILE))
const grpcObj = (grpc.loadPackageDefinition(packageDef) as unknown) as ProtoGrpcType

const client = new grpcObj.service.Main(
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
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const serverSdp: string = await new Promise<string>(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(err.message)
                return err.message
            }

            client.Preview({ sdp: clientSdp }, (err, res) => {
                if (err) {
                    console.log(err)
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
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const res: StreamResponse__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
                return
            }

            client.StartStream({}, (err, res) => {
                if (err) {
                    console.log(err)
                    resolve(undefined)
                    return
                }
                resolve(res)
            })
        })
    })

    return res
}

export const stopStream = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const res: StatusResponse__Output | undefined = await new Promise(resolve => {

        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
                return
            }

            client.StopStream({}, (err, res) => {
                if (err) {
                    console.log(err)
                    resolve(undefined)
                    return
                }

                if (res) {
                    resolve(res)
                    return
                }
            })
        })
    })

    return res
}


export const startPreview = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const res: StatusResponse__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
                return
            }

            client.StartPreview({}, (err, res) => {
                if (err) {
                    console.log(err)
                    resolve(undefined)
                    return
                }

                if (res) {
                    resolve(res)
                    return
                }
            })
        })

    })

    return res
}

export const stopPreview = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const res: StatusResponse__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
                return
            }

            client.StopPreview({}, (err, res) => {
                if (err) {
                    console.log(err)
                    resolve(undefined)
                    return
                }

                if (res) {
                    resolve(res)
                    return
                }
            })
        })
    })

    return res
}


export const checkStatus = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const status: StatusResponse__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
            }

            client.Status({}, (err, res) => {
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

    return status
}

export const findNewSongsNCS = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    client.waitForReady(deadline, (err) => {
        if (err) {
            return
        }

        client.FindNewSongsNCS({}, (err, res) => {
            if (err) {
                return
            }
        })
    })
}

export const statusNCS = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const status: StatusNCSResponse__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
            }

            client.StatusNCS({}, (err, res) => {
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

    return status
}

export const saveTwitchStreamKey = async (streamKey: string) => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const success: boolean = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(false)
            }

            client.TwitchSaveStreamKey({ key: streamKey }, (err, res) => {
                if (err) {
                    console.log(err)
                    resolve(false)
                }
                if (res !== undefined) {
                    resolve(true)
                } else {
                    resolve(false)
                }
            })

        })
    })

    return success
}

export const checkTwitchStreamKey = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const status: TwitchStreamKey__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                resolve(undefined)
            }

            client.CheckTwitchStreamKey({}, (err, res) => {
                if (err) {
                    const status = { active: false }
                    resolve(status)
                }
                if (res !== undefined) {
                    resolve(res)
                } else {
                    const status = { active: false }
                    resolve(status)
                }
            })

        })
    })

    return status
}

export const deleteTwitchStreamKey = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const success: boolean = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(false)
            }

            client.DeleteTwitchStreamKey({}, (err, res) => {
                if (err) {
                    resolve(false)
                    console.log(err)
                }
                if (res !== undefined) {
                    resolve(true)
                } else {
                    resolve(false)
                }
            })

        })
    })

    return success
}

export const saveTwitchCredentials = async (client_id: string, secret: string) => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const success: boolean = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(false)
            }

            client.SaveTwitchDevCredentials({ clientId: client_id, secret: secret }, (err, res) => {
                if (err) {
                    resolve(false)
                    console.log(err)
                }
                if (res !== undefined) {
                    resolve(true)
                } else {
                    resolve(false)
                }
            })

        })
    })

    return success
}

export const checkTwitchCredentials = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const credentials: DevCredentials__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                resolve(undefined)
            }

            client.CheckTwitchDevCredentials({}, (err, res) => {
                if (err) {
                    const status = { active: false }
                    resolve(status)
                }
                if (res !== undefined) {
                    resolve(res)
                } else {
                    const status = { active: false }
                    resolve(status)
                }
            })

        })
    })

    return credentials
}

export const deleteTwitchDevCredentials = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const success: boolean = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(false)
            }

            client.DeleteTwitchDevCredentials({}, (err, res) => {
                if (err) {
                    resolve(false)
                    console.log(err)
                }
                if (res !== undefined) {
                    resolve(true)
                } else {
                    resolve(false)
                }
            })

        })
    })

    return success
}

export const twitchAccessToken = async (code: string) => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const success: boolean = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(false)
            }

            client.TwitchAccessToken({ code: code }, (err, res) => {
                if (err) {
                    resolve(false)
                    console.log(err)
                }
                if (res !== undefined) {
                    resolve(true)
                } else {
                    resolve(false)
                }
            })

        })
    })

    return success
}

export const getOverlays = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const overlays: Overlays__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
            }

            client.GetOverlays({}, (err, res) => {
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

    return overlays
}

export const getBgVideos = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const videos: BackgroundVideosResponse__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
            }

            client.BackgroundVideos({}, (err, res) => {
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

    return videos
}

export const swapBgVideo = async (video: BackgroundVideo) => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
            }

            client.SwapBackgroundVideo(video, (err, res) => {
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

    return
}

export const uploadVideoFile = async (formData: FormData) => {
    const file = formData.get("file") as File

    type UploadResponse = {
        ok: boolean,
        msg: string
    }

    const res: UploadResponse = await new Promise(async (resolve) => {
        const call = client.UploadVideo((error, res) => {
            if (error) {
                if (error.code === 6) {
                    resolve({ ok: false, msg: "file already exists" })
                    call.end()
                    return
                }

                resolve({ ok: false, msg: error.message })
                call.end()
                return
            } else {
                resolve({ ok: true, msg: "" })
            }

        })

        call.write({ info: { fileName: file.name, size: file.size.toString(), type: file.type } })

        var chunkCounter = 0;

        //break into 2 MB chunks fat minimum
        const chunkSize = 2097152;

        var start = 0;
        var chunkEnd = start + chunkSize;

        do {
            // Create chunk
            chunkCounter++;
            chunkEnd = Math.min(start + chunkSize, file.size);
            const chunk = await file.slice(start, chunkEnd).arrayBuffer();

            const buffer = Buffer.from(chunk);

            // upload chunk
            call.write({ chunk: buffer })

            start += chunkSize;

        } while (start < file.size);

        call.end()

        return
    })

    return res
}

export const deleteBgVideo = async (video: BackgroundVideo) => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    type DeleteResponse = {
        ok: boolean,
        msg: string
    }

    const res: DeleteResponse = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve({ ok: false, msg: err.message })
            }

            client.DeleteBackgroundVideo(video, (err, res) => {
                if (err) {
                    console.log(err)
                    resolve({ ok: false, msg: err.message })
                }

                if (res !== undefined) {
                    resolve({ ok: true, msg: "" })
                } else {
                    resolve({ ok: false, msg: "" })
                }
            })
        })
    })

    return res
}

export const getStreamParams = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const res: StreamParametersResponse__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(undefined)
                return
            }

            client.StreamParameters({}, (err, res) => {
                if (err) {
                    console.log(err)
                    resolve(undefined)
                    return
                }

                if (res) {
                    resolve(res)
                    return
                }
            })
        })
    })

    if (res) {
        if (res.volume === 101) {
            res.volume = 0
        }

        res.volume = Math.round(res.volume! * 100)
    }

    return res
}

export const saveStreamParams = async (params: SaveStreamParametersRequest) => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const res: boolean = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                console.log(err)
                resolve(false)
                return
            }

            client.SaveStreamParameters(params, (err, res) => {
                if (err) {
                    console.log(err)
                    resolve(false)
                    return
                }

                if (res) {
                    resolve(true)
                    return
                }
            })
        })
    })

    return res
}