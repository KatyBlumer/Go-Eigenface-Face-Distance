package faceimage

import (
	"image"
	"image/color"
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
	im := readImage(filename)

	face := faceVector{
		width:  im.Bounds().Max.X - im.Bounds().Min.X,
		height: im.Bounds().Max.Y - im.Bounds().Min.Y,
	}
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
	im := image.NewGray16(bounds)
	for y := 0; y < face.height; y++ {
		for x := 0; x < face.width; x++ {
			// ORL database images are 16-bit grayscale
			value := uint16(face.pixels[(y*face.width)+x])
			im.SetGray16(x, y, color.Gray16{value})
		}
	}
	return im
}

func AverageFaces(filenames []string) faceVector {
	faces := make([]faceVector, len(filenames))
	for i := 0; i < len(filenames); i++ {
		faces[i] = ToVector(filenames[i])
	}
	return average(faces)
}

func average(faces []faceVector) faceVector {
	width := faces[0].width
	height := faces[0].height
	avg := make([]float64, width*height)

	for i := 0; i < len(faces); i++ {
		face := faces[i]
		if face.width != width || face.height != height {
			return faceVector{}
		}
		for j := 0; j < width*height; j++ {
			// TODO check what this does to precision
			avg[j] += face.pixels[j]
		}
	}

	for j := 0; j < width*height; j++ {
		avg[j] = avg[j] / float64(len(faces))
	}
	return faceVector{
		width:  width,
		height: height,
		pixels: avg,
	}
}
