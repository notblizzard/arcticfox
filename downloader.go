package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// GetImage gets image url
func GetImage(apiKey string) (imageURL string, imageUser string) {
	rand.Seed(time.Now().UnixNano())
	api := "https://api.unsplash.com/photos/random"
	fmt.Println(apiKey)
	res, err := http.Get(fmt.Sprintf("%s?client_id=%s&query=arctic+fox", api, apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	results := make(map[string]map[string]string)

	json.NewDecoder(res.Body).Decode(&results)

	return results["urls"]["raw"], results["user"]["name"]
}

func makeFolder(path string) {
	currentDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}
	fullPath := filepath.Join(currentDir, path)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		os.Mkdir(fullPath, 0700)
	}
}

// DownloadImage downloads image and returns filename
func DownloadImage(imageURL, imageUser string) string {
	img, err := http.Get(imageURL)
	if err != nil {
		log.Fatal(err)
	}
	fileName := fmt.Sprintf("%d-%s.png", time.Now().Unix(), imageUser)
	makeFolder("./Foxes/")
	file, err := os.Create(fmt.Sprintf("./Foxes/%s", fileName))
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, img.Body)
	if err != nil {
		log.Fatal(err)
	}
	return fileName
}
