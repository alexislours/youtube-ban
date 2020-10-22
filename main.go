package main

import (
	"bytes"
	"fmt"
	"github.com/icza/mjpeg"
	"image"
	"image/color"
	"image/jpeg"
	"time"
)

func main() {
	fmt.Printf("Generating video...\n")
	start := time.Now()

	width, height := 1920, 1080

	img := createImage(width, height)
	aw, _ := mjpeg.New("out.mp4", 1920, 1080, 1)

	buf := &bytes.Buffer{}
	jpeg.Encode(buf, img, nil)
	for i := 0; i < 10; i++ {
		aw.AddFrame(buf.Bytes())
	}

	aw.Close()
	end := time.Since(start)
	fmt.Printf("\nCreated video as \"out.mp4\". Took %.2fs\nUpload to YouTube for a quick account termination :)\n", end.Seconds())
}

func createImage(width int, height int) *image.RGBA {
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	yellow := color.RGBA{255, 255, 0, 255}
	green := color.RGBA{0, 0, 255, 255}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if x%2 == 0 {
				img.Set(x, y, green)
				y++
				img.Set(x, y, yellow)
			} else {
				img.Set(x, y, yellow)
				y++
				img.Set(x, y, green)
			}
		}
	}
	return img
}
