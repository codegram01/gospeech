# GoSpeech - Golang Speech

GoSpeech can generate audio for you from CLI param or from file 

[Video when i create GoSpeech](https://www.youtube.com/live/UpTZ_D2PUKU?si=XzEP6NNoX6MR0247)

## Usage
```sh
gospeech -m 'Hello World' -o hello.mp3
gospeech -f content.txt -o content
```

## Install
Prerequisites: [Install Golang](https://go.dev/doc/install) 

### Option 1: Install Binary
```sh
go install github.com/codegram01/gospeech@latest
```

### Option 2: Install and build from source
```sh
git clone https://github.com/codegram01/gospeech.git

cd gospeech

# Run code: 
go run . 

# Build 
go build .

# Install command 
go install 
```

