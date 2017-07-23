package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	flags "github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		Text            string `short:"t" long:"text" description:"The text to read" required:"true"`
		Voice           string `short:"v" long:"voice" description:"The voice ID to use to read the text" required:"true"`
		Rate            string `short:"r" long:"rate" description:"The reading speed rate" choice:"x-slow" choice:"slow" choice:"medium" choice:"fast" choice:"x-fast" default:"medium"`
		Output          string `short:"o" long:"output" description:"Path to the output file" default:"output.mp3"`
		AccessKeyID     string `long:"accessKeyID" description:"AWS Access Key ID" required:"false"`
		SecretAccessKey string `long:"SecretAccessKey" description:"AWS Secret Access Key" required:"false"`
		AWSRegion       string `long:"AWSRegion" description:"AWS Region" required:"false" default:"eu-west-1"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	config := &aws.Config{}
	config = config.WithRegion(opts.AWSRegion)
	if opts.AccessKeyID != "" && opts.SecretAccessKey != "" {
		config = config.WithCredentials(credentials.NewStaticCredentials(opts.AccessKeyID, opts.SecretAccessKey, ""))
	} else {
		config = config.WithCredentials(credentials.NewEnvCredentials())
	}

	svc := polly.New(session.Must(session.NewSession(config)))

	text := fmt.Sprintf("<speak><prosody rate='%s'>%s</prosody></speak>", strings.ToLower(opts.Rate), opts.Text)
	input := &polly.SynthesizeSpeechInput{}
	input = input.SetText(text).SetTextType("ssml")
	input = input.SetOutputFormat("mp3")
	input = input.SetVoiceId(opts.Voice)
	err = input.Validate()

	if err != nil {
		os.Exit(1)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	output, err := svc.SynthesizeSpeechWithContext(ctx, input)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	outFile, err := os.Create(opts.Output)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, output.AudioStream)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
