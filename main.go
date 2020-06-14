package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func download(unsplashKey, dropboxKey string) {
	imgURL, imageUser := GetImage(unsplashKey)
	fileName := DownloadImage(imgURL, imageUser)
	UploadToDropbox(fileName, dropboxKey)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	unsplashKey := os.Getenv("UNSPLASH_KEY")
	dropboxKey := os.Getenv("DROPBOX_KEY")
	hourTicker := time.NewTicker(60 * time.Minute)

	// run once, then start the loop
	download(unsplashKey, dropboxKey)
	for {
		select {
		case <-hourTicker.C:
			download(unsplashKey, dropboxKey)
		}
	}
}
