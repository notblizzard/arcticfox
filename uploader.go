package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
)

// UploadToDropbox uploads image to dropbox
func UploadToDropbox(fileName string, apiKey string) {
	config := dropbox.Config{
		Token:    apiKey,
		LogLevel: dropbox.LogInfo,
	}
	file := files.New(config)
	img, err := os.Open(fmt.Sprintf("./Foxes/%s", fileName))
	if err != nil {
		log.Fatal(err)
	}
	file.Upload(&files.CommitInfo{Path: fmt.Sprintf("/%s", fileName), Autorename: true, Mode: &files.WriteMode{Update: "aaaaaaaaa", Tagged: dropbox.Tagged{Tag: "update"}}}, bufio.NewReader(img))
}
