package main

import (
	"fmt"
	"github.com/KatyBlumer/Go-Eigenface-Face-Distance/faceimage"
	"image"
	"net/http"
)

type FaceData struct {
	vector faceimage.FaceVector
	img    image.Image
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/test", scratch)
	http.ListenAndServe(":8000", nil)
}

func getPath() string {
	return "/Users/kblumer/go/src/github.com/KatyBlumer/Go-Eigenface-Face-Distance/web/"
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func scratch(w http.ResponseWriter, r *http.Request) {
	f := faceimage.ToVector(getPath() + "static/img/orl_faces/1.png")
	averageFaces(5)
	saveVector(f, w, r)
	fmt.Fprint(w, "Done!")
}

func averageFaces(numFaces int) image.Image {
	filePattern := getPath() + "static/img/orl_faces/%v.png"
	filenames := make([]string, numFaces)
	for i := 0; i < numFaces; i++ {
		filenames[i] = fmt.Sprintf(filePattern, i+1)
	}
	avgFace := faceimage.AverageFaces(filenames)
	return faceimage.ToImage(avgFace)
}

func saveVector(face faceimage.FaceVector, w http.ResponseWriter, r *http.Request) {
}

// var rootTemplate = template.Must(template.New("root").Parse(rootTemplateHTML))

const rootTemplateHTML = `
<html><body>
<form action="{{.}}" method="POST" enctype="multipart/form-data">
Upload File: <input type="file" name="file"><br>
<input type="submit" name="submit" value="Submit">
</form>
</body></html>
`
