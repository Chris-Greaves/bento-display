package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	maxImageSize  int64  = 500 * 1024 // 500 KB
	maxDimensions int    = 1500       // Max width or height in pixels
	imagesDir     string = "images"
	optimisedDir  string = "optimized_images"
)

type Photo struct {
	FullPath          string `json:"fullpath"`
	WebPath           string `json:"web_path"`
	Filename          string `json:"filename"`
	Filesize          int64  `json:"filesize"`
	OptimisedFilePath string `json:"optimised_file_path"`
	WebOptimisedPath  string `json:"web_optimised_path"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mediaDir, exists := os.LookupEnv("MEDIA_DIR")
	if !exists {
		log.Fatal("MEDIA_DIR environment variable not set")
	}

	staticDir, exists := os.LookupEnv("STATIC_DIR")
	if !exists {
		log.Fatal("STATIC_DIR environment variable not set")
	}

	maxImageSizeStr, exists := os.LookupEnv("MAX_IMAGE_SIZE")
	if exists {
		maxImageSize, err = strconv.ParseInt(maxImageSizeStr, 10, 64)
		if err != nil {
			log.Fatal("Error parsing MAX_IMAGE_SIZE:", err)
		}
	}

	maxDimensionsStr, exists := os.LookupEnv("MAX_DIMENSIONS")
	if exists {
		maxDimensions, err = strconv.Atoi(maxDimensionsStr)
		if err != nil {
			log.Fatal("Error parsing MAX_DIMENSIONS:", err)
		}
	}

	var photos []Photo

	_ = filepath.Walk(mediaDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only deal with files we can handle
		ext := filepath.Ext(info.Name())
		if ext != ".jpg" && ext != ".jpeg" {
			return nil
		}

		outputPath := filepath.Join(staticDir, imagesDir, info.Name())

		if !info.IsDir() {
			photo := Photo{
				FullPath: outputPath,
				WebPath:  filepath.Join(imagesDir, info.Name()),
				Filename: info.Name(),
				Filesize: info.Size(),

				// Use original path, will be updated if compressed
				OptimisedFilePath: outputPath,
				WebOptimisedPath:  filepath.Join(imagesDir, info.Name()),
			}

			if info.Size() > maxImageSize {
				compressedPath := filepath.Join(staticDir, imagesDir, optimisedDir, info.Name())
				err = os.MkdirAll(filepath.Dir(compressedPath), os.ModePerm)
				if err != nil {
					log.Fatal(err)
				}

				//qualityOnlyCompression(path, compressedPath, &photo)
				resizeFirstCompression(path, compressedPath, &photo)
			}

			os.Symlink(path, outputPath)

			photos = append(photos, photo)
		}

		return nil
	})

	//fmt.Println(photos) // prints the list of photos in JSON format

	for _, photo := range photos {
		fmt.Printf("Path: %s, Optimised Path: %s, Filename: %s, Size: %d bytes\n", photo.FullPath, photo.OptimisedFilePath, photo.Filename, photo.Filesize)
	}

	imageData, err := os.Create(filepath.Join(staticDir, imagesDir, "imageData.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer imageData.Close()

	// Write JSON data to the file
	encoder := json.NewEncoder(imageData)
	err = encoder.Encode(photos)
	if err != nil {
		log.Fatal(err)
	}
}

func qualityOnlyCompression(path string, compressedPath string, photo *Photo) {
	original, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()

	originalImage, _, err := image.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(compressedPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = imageToJpeg(originalImage, file, maxImageSize, 50)
	if err != nil {
		log.Fatal(err)
	}

	photo.OptimisedFilePath = filepath.Join(compressedPath)
}

func resizeFirstCompression(path, compressedPath string, photo *Photo) {
	created, err := createOptimisedFile(path, compressedPath, maxImageSize, 50, maxDimensions)
	if err != nil {
		log.Fatal(err)
	}
	if created {
		photo.OptimisedFilePath = compressedPath
		photo.WebOptimisedPath = filepath.Join(imagesDir, optimisedDir, photo.Filename)
	}
}
