package web

import (
	"fmt"
	"image"
	_ "image/png"
)

func readImage(filename string) image.Image {
	return image.Decode(r)
}
