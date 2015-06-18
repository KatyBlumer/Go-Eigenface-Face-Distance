package faceimage

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

type faceVector struct {
	pixels        []float64
	width, height int
}

func readImage(filename string) image.Image {
	file, _ := os.Open(filename)
	im, _, _ := image.Decode(file)
	return im
}

func ToVector(filename string) faceVector {
	fmt.Fprintln(os.Stderr, "HIIIIIIIIIIIIIIIII___________________________________________________________________")
	im := readImage(filename)
	fmt.Fprintf(os.Stderr, "%v, %v, %v, %v \n", im.Bounds().Min.X, im.Bounds().Min.Y, im.Bounds().Max.X, im.Bounds().Max.Y)

	face := faceVector{
		width:  im.Bounds().Max.X - im.Bounds().Min.X,
		height: im.Bounds().Max.Y - im.Bounds().Min.Y,
	}
	fmt.Fprintf(os.Stderr, "%v, %v \n", face.width, face.height)
	face.pixels = make([]float64, face.width*face.height)
	minX := im.Bounds().Min.X
	minY := im.Bounds().Min.Y
	// iterate through image row by row
	for y := 0; y < face.height; y++ {
		for x := 0; x < face.width; x++ {
			color := im.At(x-minX, y-minY)
			// ORL database images are 16-bit grayscale, so can use any of RGB values
			value, _, _, _ := color.RGBA()
			face.pixels[(y*face.width)+x] = float64(value)
		}
	}
	return face
}

func ToImage(face faceVector) image.Image {
	bounds := image.Rect(0, 0, face.width, face.height)
	return image.NewRGBA(bounds)
}
