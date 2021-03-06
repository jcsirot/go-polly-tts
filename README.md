# go-polly-tts
A very simple TTS application using [AWS Polly](https://aws.amazon.com/polly/) service

## Install

    go get -u github.com/jcsirot/go-polly-tts

## Build

    go build -o go-polly-tts

## Usage

``` plain
Usage:
  go-polly-tts [OPTIONS] text

Application Options:
  -v, --voice=                                The voice ID to use to read the text
  -r, --rate=[x-slow|slow|medium|fast|x-fast] The reading speed rate (default: medium)
  -o, --output=                               Path to the output file (default: output.mp3)
      --accessKeyID=                          AWS Access Key ID
      --SecretAccessKey=                      AWS Secret Access Key
      --AWSRegion=                            AWS Region (default: eu-west-1)

Help Options:
  -h, --help                                  Show this help message

Arguments:
  text:                                       The text to read
```

### Authentication and credentials

Using Amazon Polly requires authentication (for billing). You can use `--accessKeyID` and `--SecretAccessKey` options to pass amazon credentials on the command line. If one or both arguments are not present the credentials are read from environment variables `AWS_ACCESS_KEY` and `AWS_SECRET_ACCESS_KEY`.

## Usage

    go-polly-tts -v Mizuki -r slow -o message.mp3 -- お客様は神様です