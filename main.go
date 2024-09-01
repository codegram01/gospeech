package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	message = flag.String("m", "", "Message Speech")
	filePath = flag.String("f", "", "Input File")
	outPath = flag.String("o", "", "Output Path")
)

func main() {
	flag.Parse()

	if *message == "" && *filePath == "" {
		log.Fatal("Message or FileInp require")
	}
	if *message != "" && *filePath != "" {
		log.Fatal("Not Both message and file input")
	}


	if *message != "" {
		if *outPath == "" {
			*outPath = "out.mp3"
		}
		err := GenSpeech(*message, *outPath)

		if err != nil {
			log.Fatal(err)
		}
		return
	}

    // we do read file here
	if *outPath == "" {
		*outPath = "."
	}


	lines, err := ReadFileByLine(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range lines {
		outSpeechPath := fmt.Sprintf("%s/%d-%s.mp3", *outPath, i, line.Name)

		fmt.Println(outSpeechPath)

		err := GenSpeech(line.Content, outSpeechPath)

		if err != nil {
			log.Fatal(err)
		}
	}
}