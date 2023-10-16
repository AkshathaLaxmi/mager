package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
)

func resizer() {

	var input string

	fmt.Println("Please enter the image file name: ")
	fmt.Scanln(&input)
	filename := strings.TrimSpace(input)

	fmt.Println("Please enter the new width: ")
	fmt.Scanln(&input)

	newWidth, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	splitString := strings.Split(filename, ".")

	format := strings.ToLower(splitString[len(splitString)-1])

	reader, err := os.Open(fmt.Sprintf("images/%s", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	var resizedImage *image.RGBA = resizeImage(reader, filename, newWidth)

	output, err := os.Create(fmt.Sprintf("outputImg/%s", filename))

	if err != nil {
		log.Fatal(err)
	}

	defer output.Close()

	switch format {
	case "png":
		err = png.Encode(output, resizedImage)
	case "jpg", "jpeg":
		err = jpeg.Encode(output, resizedImage, nil)
	case "gif":
		err = gif.Encode(output, resizedImage, nil)
	default:
		log.Fatal("Format not supported")
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Image resized and saved as outputImg/%s\n", filename)
}

func resizeImage(reader *os.File, filename string, newWidth int) *image.RGBA {
	img, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()

	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()

	newHeight := (newWidth * originalHeight) / originalWidth

	resizedmage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			xOriginal := x * originalWidth / newWidth
			yOriginal := y * originalHeight / newHeight

			color := img.At(xOriginal, yOriginal)

			resizedmage.Set(x, y, color)
		}
	}

	return resizedmage
}

func main() {
	resizer()
}
