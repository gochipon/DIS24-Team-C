package main

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	"log"
	"os"
	"os/signal"
)

const sampleRate = 44100
const framesPerBuffer = 64

func main() {
	err := portaudio.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer portaudio.Terminate()

	in := make([]int16, framesPerBuffer)
	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(in), in)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	err = stream.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	fmt.Println("Recording. Press Ctrl+C to stop.")
	for {
		select {
		case <-sig:
			fmt.Println("\nStopping.")
			return
		default:
			err := stream.Read()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Audio data:", in)
		}
	}
}
