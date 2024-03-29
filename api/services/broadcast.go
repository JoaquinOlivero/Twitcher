package service

import (
	"context"
	"io"
	"log"
	"math"
	"os"
	"time"

	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/h264reader"
)

func (s *MainServer) Broadcast(sdpFromClient <-chan string, sdpForClientChannel chan<- string, exit <-chan struct{}) {
	stop := make(chan struct{})
	defer log.Println("closing broadcast func")

	// Everything below is the Pion WebRTC API, thanks for using it ❤️.
	offer := webrtc.SessionDescription{}
	Decode(<-sdpFromClient, &offer)

	peerConnectionConfig := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(peerConnectionConfig)
	if err != nil {
		panic(err)
	}
	defer func() {
		if cErr := peerConnection.Close(); cErr != nil {
			log.Printf("cannot close peerConnection: %v\n", cErr)
		}
		log.Println("peer connection closed")
	}()

	iceConnectedCtx, iceConnectedCtxCancel := context.WithCancel(context.Background())

	// Open named audio pipe
	audioPipe, err := os.OpenFile("files/stream/previewAudio", os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer audioPipe.Close()

	// Open named audio pipe
	videoPipe, err := os.OpenFile("files/stream/previewVideo", os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer videoPipe.Close()

	// Create a video track
	localVideoTrack, videoTrackErr := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeH264}, "video", "pion")
	if videoTrackErr != nil {
		panic(videoTrackErr)
	}

	rtpSender, videoTrackErr := peerConnection.AddTrack(localVideoTrack)
	if videoTrackErr != nil {
		panic(videoTrackErr)
	}

	// Read incoming RTCP packets
	// Before these packets are returned they are processed by interceptors. For things
	// like NACK this needs to be called.
	go func(stop <-chan struct{}) {
		defer log.Println("closing video rtcp")
		rtcpBuf := make([]byte, 1500)

	outer:
		for {
			select {
			case <-stop:
				break outer
			case <-exit:
				break outer
			default:
				if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
					return
				}
			}
		}
	}(stop)

	go func(exit <-chan struct{}) {
		defer log.Println("closing video")
		if err != nil {
			panic(err)
		}

		h264, h264Err := h264reader.NewReader(videoPipe)
		if h264Err != nil {
			panic(h264Err)
		}

		// Wait for connection established
		<-iceConnectedCtx.Done()

		// Send our video file frame at a time. Pace our sending so we send it at the same speed it should be played back as.
		// This isn't required since the video is timestamped, but we will such much higher loss if we send all at once.
		//
		// It is important to use a time.Ticker instead of time.Sleep because
		// * avoids accumulating skew, just calling time.Sleep didn't compensate for the time spent parsing the data
		// * works around latency issues with Sleep (see https://github.com/golang/go/issues/44343)
		h264FrameDuration := math.Round((float64(1) / float64(s.streamParams.fps)) * 1000)
		spsAndPpsCache := []byte{}
		ticker := time.NewTicker(time.Duration(h264FrameDuration * float64(time.Millisecond)))

	outer:
		for ; true; <-ticker.C {
			select {
			case <-exit:
				break outer
			default:
				nal, h264Err := h264.NextNAL()
				if h264Err == io.EOF {
					log.Printf("All video frames parsed and sent")
					break outer
				}
				if h264Err != nil {
					log.Println(h264Err)
					break outer
				}

				nal.Data = append([]byte{0x00, 0x00, 0x00, 0x01}, nal.Data...)

				if nal.UnitType == h264reader.NalUnitTypeSPS || nal.UnitType == h264reader.NalUnitTypePPS {
					spsAndPpsCache = append(spsAndPpsCache, nal.Data...)
					continue
				} else if nal.UnitType == h264reader.NalUnitTypeCodedSliceIdr {
					nal.Data = append(spsAndPpsCache, nal.Data...)
					spsAndPpsCache = []byte{}
				}

				if h264Err = localVideoTrack.WriteSample(media.Sample{Data: nal.Data, Duration: time.Second}); h264Err != nil {
					panic(h264Err)
				}
			}
		}
	}(exit)

	// Create an audio track
	localAudioTrack, audioTrackErr := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus}, "audio", "pion")
	if audioTrackErr != nil {
		panic(audioTrackErr)
	}

	rtpSenderAudio, audioTrackErr := peerConnection.AddTrack(localAudioTrack)
	if audioTrackErr != nil {
		panic(audioTrackErr)
	}

	// Read incoming RTCP packets
	// Before these packets are returned they are processed by interceptors. For things
	// like NACK this needs to be called.
	go func(stop <-chan struct{}) {
		defer log.Println("closing audio rtcp")
		rtcpBuf := make([]byte, 1500)

	outer:
		for {
			select {
			case <-stop:
				break outer
			case <-exit:
				break outer
			default:
				if _, _, rtcpErr := rtpSenderAudio.Read(rtcpBuf); rtcpErr != nil {
					return
				}
			}
		}
	}(stop)

	go func(exit <-chan struct{}) {
		defer log.Println("closing audio")
		// Open on oggfile in non-checksum mode.
		ogg, _, oggErr := NewWith(audioPipe)
		if oggErr != nil {
			panic(oggErr)
		}

		// Wait for connection established
		<-iceConnectedCtx.Done()

		// Keep track of last granule, the difference is the amount of samples in the buffer
		var lastGranule uint64

		// It is important to use a time.Ticker instead of time.Sleep because
		// * avoids accumulating skew, just calling time.Sleep didn't compensate for the time spent parsing the data
		// * works around latency issues with Sleep (see https://github.com/golang/go/issues/44343)
		oggPageDuration := time.Microsecond * 2000
		ticker := time.NewTicker(oggPageDuration)

	outer:
		for ; true; <-ticker.C {
			select {
			case <-exit:
				break outer
			default:
				pageData, pageHeader, oggErr := ogg.ParseNextPage()
				if oggErr == io.EOF {
					log.Printf("All audio pages parsed and sent")
				}

				if oggErr != nil {
					// panic(oggErr)
					break outer
				}

				// The amount of samples is the difference between the last and current timestamp
				sampleCount := float64(pageHeader.GranulePosition - lastGranule)
				lastGranule = pageHeader.GranulePosition
				sampleDuration := time.Duration((sampleCount/48000)*1000) * time.Millisecond

				if oggErr = localAudioTrack.WriteSample(media.Sample{Data: pageData, Duration: sampleDuration}); oggErr != nil {
					panic(oggErr)
				}
			}
		}
	}(exit)

	// Register data channel creation handling
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		// Register channel opening handling
		d.OnOpen(func() {
			log.Println("data channel has been opened")
		outer:
			for {
				select {
				case <-exit:
					log.Println("exit data channel for loop")
					d.Close()
					break outer
				case msg := <-s.channels.sendMsgDataChannel:
					// Send message when new song starts playing.
					err := d.SendText(msg)
					if err != nil {
						d.Close()
						break outer
					}
				case cover := <-s.channels.sendOverlayDataChannel:
					// Send message when new song starts playing.
					err := d.SendText(cover)
					if err != nil {
						d.Close()
						break outer
					}
				}
			}
		})

		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			s.manageDataChannelMessage(msg.Data)
		})

		d.OnClose(func() {
			peerConnection.Close()
		})
	})

	// Set the handler for ICE connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		log.Printf("Connection State has changed %s \n", connectionState.String())
		if connectionState == webrtc.ICEConnectionStateConnected {
			iceConnectedCtxCancel()
		}

		if connectionState.String() == "disconnected" {
			peerConnection.Close()
			return
		}
	})

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		log.Printf("Peer Connection State has changed: %s\n", s.String())

		if s == webrtc.PeerConnectionStateFailed {
			// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
			// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
			// log.Println("Peer Connection has gone to failed exiting")
			return
		}
		if s.String() == "disconnected" {
			peerConnection.Close()
			return
		}
	})

	// Set the remote SessionDescription
	err = peerConnection.SetRemoteDescription(offer)
	if err != nil {
		peerConnection.Close()
		panic(err)
	}

	// Create answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		peerConnection.Close()
		panic(err)
	}

	// Create channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// Sets the LocalDescription, and starts our UDP listeners
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		peerConnection.Close()
		panic(err)
	}

	// Block until ICE Gathering is complete, disabling trickle ICE
	// we do this because we only can exchange one signaling message
	// in a production application you should exchange ICE Candidates via OnICECandidate
	<-gatherComplete

	// Get the LocalDescription and send it to client trying to connect
	sdpForClientChannel <- Encode(*peerConnection.LocalDescription())

outer:
	for {
		select {
		case <-exit:
			break outer
		case sdp := <-sdpFromClient:
			stop := make(chan struct{})

			recvOnlyOffer := webrtc.SessionDescription{}
			Decode(sdp, &recvOnlyOffer)

			// Create a new PeerConnection
			peerConnection, err := webrtc.NewPeerConnection(peerConnectionConfig)
			if err != nil {
				peerConnection.Close()
				panic(err)
			}

			rtpSender, err = peerConnection.AddTrack(localVideoTrack)
			if err != nil {
				peerConnection.Close()
				panic(err)
			}

			// Read incoming RTCP packets
			// Before these packets are returned they are processed by interceptors. For things
			// like NACK this needs to be called.
			go func(stop <-chan struct{}) {
				rtcpBuf := make([]byte, 3000)

			outer:
				for {
					select {
					case <-stop:
						break outer
					case <-exit:
						break outer
					default:
						if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
							peerConnection.Close()
							return
						}
					}
				}
			}(stop)

			rtpAudioSender, err := peerConnection.AddTrack(localAudioTrack)
			if err != nil {
				panic(err)
			}

			// Read incoming RTCP packets
			// Before these packets are returned they are processed by interceptors. For things
			// like NACK this needs to be called.
			go func(stop <-chan struct{}) {
				rtcpBuf := make([]byte, 3000)

			outer:
				for {
					select {
					case <-exit:
						break outer
					case <-stop:
						break outer
					default:
						if _, _, rtcpErr := rtpAudioSender.Read(rtcpBuf); rtcpErr != nil {
							peerConnection.Close()
							return
						}
					}
				}
			}(stop)

			// Register data channel creation handling
			peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
				// Register channel opening handling
				d.OnOpen(func() {
					log.Println("data channel has been opened")
				outer:
					for {
						select {
						case <-exit:
							log.Println("exit data channel for loop")
							d.Close()
							break outer
						case msg := <-s.channels.sendMsgDataChannel:
							// Send message when new song starts playing.
							err := d.SendText(msg)
							if err != nil {
								d.Close()
								break outer
							}
						case cover := <-s.channels.sendOverlayDataChannel:
							// Send message when new song starts playing.
							err := d.SendText(cover)
							if err != nil {
								d.Close()
								break outer
							}
						}
					}
				})

				d.OnMessage(func(msg webrtc.DataChannelMessage) {
					s.manageDataChannelMessage(msg.Data)
				})

				d.OnClose(func() {
					peerConnection.Close()
				})
			})

			// Set the handler for ICE connection state
			// This will notify you when the peer has connected/disconnected
			peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
				log.Printf("Connection State has changed %s \n", connectionState.String())
				if connectionState == webrtc.ICEConnectionStateConnected {
					iceConnectedCtxCancel()
				}

				if connectionState.String() == "disconnected" {
					peerConnection.Close()
					close(stop)
					return
				}
			})

			// Set the handler for Peer connection state
			// This will notify you when the peer has connected/disconnected
			peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
				log.Printf("Peer Connection State has changed: %s\n", s.String())

				if s == webrtc.PeerConnectionStateFailed {
					// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
					// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
					// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
					// log.Println("Peer Connection has gone to failed exiting")
					peerConnection.Close()
					return
				}
				if s.String() == "disconnected" {
					peerConnection.Close()
					return
				}
			})

			// Set the remote SessionDescription
			err = peerConnection.SetRemoteDescription(recvOnlyOffer)
			if err != nil {
				peerConnection.Close()
				panic(err)
			}

			// Create answer
			answer, err := peerConnection.CreateAnswer(nil)
			if err != nil {
				peerConnection.Close()
				panic(err)
			}

			// Create channel that is blocked until ICE Gathering is complete
			gatherComplete = webrtc.GatheringCompletePromise(peerConnection)

			// Sets the LocalDescription, and starts our UDP listeners
			err = peerConnection.SetLocalDescription(answer)
			if err != nil {
				peerConnection.Close()
				panic(err)
			}

			// Block until ICE Gathering is complete, disabling trickle ICE
			// we do this because we only can exchange one signaling message
			// in a production application you should exchange ICE Candidates via OnICECandidate
			<-gatherComplete

			// Get the LocalDescription and send it to client trying to connect
			sdpForClientChannel <- Encode(*peerConnection.LocalDescription())
		}
	}

}
