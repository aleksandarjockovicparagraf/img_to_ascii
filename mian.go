package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	_ = os.Args
	_ = os.Args[1:]
	arg := os.Args[1]

	image, err := get_image_from_file_path(arg)

	if err != nil {
		log.Fatal(err)
	}

	grayscale := get_grayscale(image)
	map_grayscale_to_char(grayscale)
}

func get_image_from_file_path(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	image, _, err := image.Decode(f)

	width := image.Bounds().Dx()
	height := image.Bounds().Dy()
	aspectRatio := float64(width) / float64(height)

	newHeight := uint(float64(80) / aspectRatio)

	resizedImage := resize.Resize(uint(80), newHeight, image, resize.Lanczos3)

	return resizedImage, err

}

func get_grayscale(image image.Image) [][]float64 {

	bounds := image.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	grayscale := make([][]float64, height)
	for y := range grayscale {
		grayscale[y] = make([]float64, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := image.At(x, y)
			r, g, b, _ := pixel.RGBA()
			gray := (0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)) / 65535.0
			gray = math.Round(gray*10) / 10.0
			grayscale[y][x] = gray
		}
	}

	return grayscale
}

func map_grayscale_to_char(grayscale [][]float64) [][]string {

	char_map := make([][]string, len(grayscale))
	for y := range grayscale {
		char_map[y] = make([]string, len(grayscale[y]))
	}

	for y := range grayscale {
		for x := range grayscale[y] {

			gray := grayscale[y][x]

			switch gray {
			case 0.0:
				char_map[y][x] = " "
			case 0.1:
				char_map[y][x] = " "
			case 0.2:
				char_map[y][x] = "."
			case 0.3:
				char_map[y][x] = ":"
			case 0.4:
				char_map[y][x] = "-"
			case 0.5:
				char_map[y][x] = "="
			case 0.6:
				char_map[y][x] = "+"
			case 0.7:
				char_map[y][x] = "*"
			case 0.8:
				char_map[y][x] = "#"
			case 0.9:
				char_map[y][x] = "%"
			default:
				char_map[y][x] = "@"
			}

			fmt.Print(char_map[y][x])
		}

		fmt.Println()
	}

	return char_map
}
