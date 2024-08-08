package main

import (
	speech "cloud.google.com/go/speech/apiv1"
	"context"
	"fmt"
	"github.com/gordonklaus/portaudio"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	"io"
	"log"
)

const (
	sampleRate   = 16000
	channelCount = 1
	bufferSize   = 4096
)

func captureAudio(audioChannel chan []int16) {
	portaudio.Initialize()
	defer portaudio.Terminate()

	buffer := make([]int16, 4096)
	stream, err := portaudio.OpenDefaultStream(1, 0, 16000, len(buffer), buffer)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}
	defer stream.Stop()

	for {
		if err := stream.Read(); err != nil {
			log.Fatal(err)
		}
		audioChannel <- buffer
	}
}

func transcribeAudio(ctx context.Context, client *speech.Client, audioChannel chan []int16) {
	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		log.Fatalf("Could not create stream: %v", err)
	}
	defer stream.CloseSend()

	go func() {
		for audio := range audioChannel {
			req := &speechpb.StreamingRecognizeRequest{
				StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
					AudioContent: int16SliceToByteSlice(audio),
				},
			}
			if err := stream.Send(req); err != nil {
				log.Fatalf("Could not send audio: %v", err)
			}
		}
	}()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive response: %v", err)
		}

		for _, result := range resp.Results {
			fmt.Printf("Transcript: %s\n", result.Alternatives[0].Transcript)
		}
	}
}

func int16SliceToByteSlice(data []int16) []byte {
	buf := make([]byte, len(data)*2)
	for i, v := range data {
		buf[i*2] = byte(v)
		buf[i*2+1] = byte(v >> 8)
	}
	return buf
}

func main() {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	audioChannel := make(chan []int16)
	go captureAudio(audioChannel)
	transcribeAudio(ctx, client, audioChannel)
}
