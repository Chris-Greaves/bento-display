package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"math"
	"os"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

// maxSize: max size in bytes. Does not use compression if maxSize <= 0.
func imageToJpeg(in image.Image, out *os.File, maxSize int64, minQuality int) error {
	quality := 100

	if err := jpeg.Encode(out, in, &jpeg.Options{Quality: quality}); err != nil {
		return err
	}
	if maxSize <= 0 {
		return nil
	}

	stat, err := out.Stat()
	if err != nil {
		return err
	}

	size := stat.Size()

	for size > maxSize && quality > 0 {
		if err = out.Truncate(0); err != nil {
			return err
		}
		if _, err = out.Seek(0, io.SeekStart); err != nil {
			return err
		}
		quality -= 5
		if err = jpeg.Encode(out, in, &jpeg.Options{Quality: quality}); err != nil {
			return err
		}
		stat, err = out.Stat()
		if err != nil {
			return err
		}
		size = stat.Size()
	}

	if quality <= minQuality {
		return errors.New("can't reach desired filesize (quality below minimum set)")
	}

	return err
}

// createOptimisedFile resizes and compresses an image file to fit within desired constraints.
// outPath must be a jpeg file. as quality will be adjusted to meet max size.
func createOptimisedFile(imagePath string, outPath string, maxSize int64, minQuality int, maxDimensions int) (bool, error) {
	inputImage, err := os.Stat(imagePath)
	if err != nil {
		return false, err
	}

	if inputImage.Size() < maxSize {
		return false, nil
	}

	img, err := imgio.Open(imagePath)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	isLandscape := img.Bounds().Dx() > img.Bounds().Dy()

	var scaleFactor float64 = 1.0
	var newWidth int = img.Bounds().Dx()
	var newHeight int = img.Bounds().Dy()

	if isLandscape && img.Bounds().Dx() > maxDimensions {
		tempWidth := float64(newWidth)
		for i := 0; tempWidth > float64(maxDimensions); i++ {
			scaleFactor = (100.0 - float64(i)*5.0) / 100.0
			tempWidth = float64(img.Bounds().Dx()) * scaleFactor
		}
		newWidth = int(tempWidth)
		newHeight = int(math.Floor(float64(img.Bounds().Dy()) * scaleFactor))
	} else if img.Bounds().Dy() > maxDimensions { // Portrait
		tempHeight := float64(newHeight)
		for i := 0; tempHeight > float64(maxDimensions); i++ {
			scaleFactor = (100.0 - float64(i)*5.0) / 100.0
			tempHeight = float64(img.Bounds().Dy()) * scaleFactor
		}
		newWidth = int(math.Floor(float64(img.Bounds().Dx()) * scaleFactor))
		newHeight = int(tempHeight)
	}

	resizedImg := transform.Resize(img, newWidth, newHeight, transform.Linear)

	if err := imgio.Save(outPath, resizedImg, imgio.JPEGEncoder(100)); err != nil {
		fmt.Println(err)
		return false, err
	}

	quality := 100
	for quality >= minQuality {
		optimisedFile, err := os.Stat(outPath)
		if err != nil {
			return false, err
		}
		if optimisedFile.Size() < maxSize {
			log.Printf("Reduced image size by %.2f%% to %dx%d, the compressed using quality %d\n", (1.0-scaleFactor)*100.0, newWidth, newHeight, quality)
			return true, nil
		}
		quality -= 5
		if err := imgio.Save(outPath, resizedImg, imgio.JPEGEncoder(quality)); err != nil {
			fmt.Println(err)
			return false, err
		}
	}

	return true, nil
}
