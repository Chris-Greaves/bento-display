package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	// Get number from command line argument
	var num int
	flag.IntVar(&num, "n", 5, "Number of images to pull (max 2000)")
	flag.Parse()

	var pics []struct {
		Download_URL string `json:"download_url"`
	}

	for i := 1; i <= num; i++ {
		page := (i - 1) * 10
		resp, err := http.Get("https://picsum.photos/v2/list?" + "page=" + strconv.Itoa(page) + "&limit=10")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var newPics []struct {
			Download_URL string `json:"download_url"`
		}
		json.Unmarshal(body, &newPics)

		pics = append(pics, newPics...)
	}

	err := os.MkdirAll("../../test/media", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < num; i++ {
		resp, err := http.Get(pics[i].Download_URL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		filename := "image_" + strconv.Itoa(i) + ".jpg"
		file, err := os.Create(filepath.Join("../../test/media", filename))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.Write(body)
	}
}
