package web

import (
	"fmt"
	"github.com/KatyBlumer/Go-Eigenface-Face-Distance/faceimage"
	"image"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/test", scratch)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, wrld!")
}

func scratch(w http.ResponseWriter, r *http.Request) {
	faceimage.ToVector("static/img/orl_faces/1.png")
	averageFaces(5)
	fmt.Fprint(w, "Done!")
}

func averageFaces(numFaces int) image.Image {
	filePattern := "static/img/orl_faces/%v.png"
	filenames := make([]string, numFaces)
	for i := 0; i < numFaces; i++ {
		filenames[i] = fmt.Sprintf(filePattern, i+1)
	}
	avgFace := faceimage.AverageFaces(filenames)
	return faceimage.ToImage(avgFace)
}
