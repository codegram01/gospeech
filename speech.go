package main

import (
	"context"
	"os"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

func GenSpeech(textInp string, outPath string) error {
	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
			// Set the text input to be synthesized.
			Input: &texttospeechpb.SynthesisInput{
					InputSource: &texttospeechpb.SynthesisInput_Text{Text: textInp},
			},
			// Build the voice request, select the language code ("en-US") and the SSML
			// voice gender ("neutral").
			Voice: &texttospeechpb.VoiceSelectionParams{
					LanguageCode: "en-US",
					SsmlGender:   texttospeechpb.SsmlVoiceGender_MALE,
					Name: "en-US-Wavenet-I",
			},
			// Select the type of audio file you want returned.
			AudioConfig: &texttospeechpb.AudioConfig{
					AudioEncoding: texttospeechpb.AudioEncoding_MP3,
			},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		return err
	}

	err = os.WriteFile(outPath, resp.AudioContent, 0644)
	if err != nil {
		return err
	}
	
	return nil
}