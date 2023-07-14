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
    await startAudio()

    await startOutput("preview")

    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const serverSdp: string = await new Promise<string>(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                // console.log(err)
                resolve(err.message)
                return err.message
            }

            client.Preview({ sdp: clientSdp }, (err, res) => {
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

    await startOutput("stream")

    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    client.waitForReady(deadline, (err) => {
        if (err) {
            // console.log(err)
            return
        }

        client.StartStream({}, (err, res) => {
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

const startOutput = async (mode: string) => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const ready: boolean = await new Promise<boolean>(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                // console.log(err)
                resolve(false)
                return false
            }

            client.StartOutput({ mode: mode }, (err, res) => {
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


export const stopStream = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    client.waitForReady(deadline, (err) => {
        if (err) {
            // console.log(err)
            return
        }

        client.StopStream({}, (err, res) => {
            if (err) {
                // console.log(err)
                return
            }
        })
    })
}

export const checkStatus = async () => {
    const deadline = new Date()
    deadline.setSeconds(deadline.getSeconds() + 5)

    const status: StatusResponse__Output | undefined = await new Promise(resolve => {
        client.waitForReady(deadline, (err) => {
            if (err) {
                // console.log(err)
                resolve(undefined)
            }

            client.Status({}, (err, res) => {
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
                // console.log(err)
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
                    revalidatePath("/")
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
                    revalidatePath("/")
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
                    revalidatePath("/")
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
