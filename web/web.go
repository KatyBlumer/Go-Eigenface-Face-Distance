package main

import (
	"fmt"
	"github.com/KatyBlumer/Go-Eigenface-Face-Distance/faceimage"
	"image"
	"image/png"
	"net/http"
	"os"
)

type FaceData struct {
	vector faceimage.FaceVector
	img    image.Image
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/test", scratch)
	http.Handle("/static/", http.FileServer(http.Dir(getPath())))
	http.Handle("/tmp/", http.FileServer(http.Dir(getPath())))
	http.ListenAndServe(":8000", nil)
}

func getPath() string {
	return "/Users/kblumer/go/src/github.com/KatyBlumer/Go-Eigenface-Face-Distance/web/"
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func scratch(w http.ResponseWriter, r *http.Request) {
	avg := averageFaces(5)
	tempPath := "tmp/avg.png"
	saveImage(avg, getPath()+tempPath)
	fmt.Fprint(w, fmt.Sprintf(imageTemplateHTML, tempPath))
}

func averageFaces(numFaces int) faceimage.FaceVector {
	filePattern := getPath() + "static/img/orl_faces/%v.png"
	filenames := make([]string, numFaces)
	for i := 0; i < numFaces; i++ {
		filenames[i] = fmt.Sprintf(filePattern, i+1)
	}
	avgFace := faceimage.AverageFaces(filenames)
	return avgFace
}

func saveImage(face faceimage.FaceVector, path string) {
	img := faceimage.ToImage(face)
	out, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = png.Encode(out, img)
}

// var rootTemplate = template.Must(template.New("root").Parse(rootTemplateHTML))

const formTemplateHTML = `
<html><body>
<form action="{{.}}" method="POST" enctype="multipart/form-data">
Upload File: <input type="file" name="file"><br>
<input type="submit" name="submit" value="Submit">
</form>
</body></html>
`

const imageTemplateHTML = `
<html><body>
<img src="%s">
</body></html>
`
