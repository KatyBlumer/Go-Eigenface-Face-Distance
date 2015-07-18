package main

import (
	"fmt"
	"github.com/KatyBlumer/Go-Eigenface-Face-Distance/eigenface"
	"github.com/KatyBlumer/Go-Eigenface-Face-Distance/faceimage"
	"image"
	"image/png"
	"net/http"
	"os"
)

type FaceData struct {
	vector eigenface.FaceVector
	img    image.Image
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/test", scratch)
	http.HandleFunc("/testimages", testImages)
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

func averageFaces(numFaces int) eigenface.FaceVector {
	filePattern := getPath() + "static/img/orl_faces/%v.png"
	filenames := make([]string, numFaces)
	for i := 0; i < numFaces; i++ {
		filenames[i] = fmt.Sprintf(filePattern, i+1)
	}
	avgFace := faceimage.AverageFaces(filenames)
	return avgFace
}

func saveImage(face eigenface.FaceVector, path string) {
	img := faceimage.ToImage(face)
	out, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = png.Encode(out, img)
}

func scratch(w http.ResponseWriter, r *http.Request) {
	avg := averageFaces(40)
	tempPath := "tmp/avg.png"
	saveImage(avg, getPath()+tempPath)
	fmt.Fprint(w, fmt.Sprintf(imageTemplateHTML, tempPath))
}

func testImages(w http.ResponseWriter, r *http.Request) {
	faceNum := 4
	faces := normalizeFaces(faceNum)
	face := faces[faceNum-1]
	tempPath := "tmp/normalized.png"
	saveImage(face, getPath()+tempPath)
	fmt.Fprint(w, fmt.Sprintf(imageTemplateHTML, tempPath))
}

func normalizeFaces(numFaces int) []eigenface.FaceVector {
	filePattern := getPath() + "static/img/orl_faces/%v.png"
	faces := make([]eigenface.FaceVector, numFaces)
	for i := 0; i < numFaces; i++ {
		filename := fmt.Sprintf(filePattern, i+1)
		faces[i] = faceimage.ToVector(filename)
	}
	return eigenface.Normalize(faces)
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
