package faceimage

import (
	"github.com/KatyBlumer/Go-Eigenface-Face-Distance/eigenface"
	"image"
	"image/color"
	_ "image/png"
	"os"
)

func ToVector(filename string) eigenface.FaceVector {
	im := readImage(filename)

	face := eigenface.FaceVector{
		Width:  im.Bounds().Max.X - im.Bounds().Min.X,
		Height: im.Bounds().Max.Y - im.Bounds().Min.Y,
	}
	face.Pixels = make([]float64, face.Width*face.Height)
	minX := im.Bounds().Min.X
	minY := im.Bounds().Min.Y

	// iterate through image row by row
	for y := 0; y < face.Height; y++ {
		for x := 0; x < face.Width; x++ {
			color := im.At(x-minX, y-minY)
			// ORL database images are 16-bit grayscale, so can use any of RGB values
			value, _, _, _ := color.RGBA()
			face.Pixels[(y*face.Width)+x] = float64(value)
		}
	}
	return face
}

func readImage(filename string) image.Image {
	file, _ := os.Open(filename)
	im, _, _ := image.Decode(file)
	return im
}

func ToImage(face eigenface.FaceVector) image.Image {
	bounds := image.Rect(0, 0, face.Width, face.Height)
	im := image.NewGray16(bounds)
	for y := 0; y < face.Height; y++ {
		for x := 0; x < face.Width; x++ {
			// ORL database images are 16-bit grayscale
			value := uint16(face.Pixels[(y*face.Width)+x])
			im.SetGray16(x, y, color.Gray16{value})
		}
	}
	return im
}

func AverageFaces(filenames []string) eigenface.FaceVector {
	faces := make([]eigenface.FaceVector, len(filenames))
	for i := 0; i < len(filenames); i++ {
		faces[i] = ToVector(filenames[i])
	}
	return eigenface.Average(faces)
}
