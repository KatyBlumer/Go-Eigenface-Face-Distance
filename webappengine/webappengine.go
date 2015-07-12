package webappengine

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"github.com/KatyBlumer/Go-Eigenface-Face-Distance/faceimage"
	"google.golang.org/cloud"
	"image"
	"net/http"
)

type FaceData struct {
	vector faceimage.FaceVector
	img    image.Image
}

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/test", scratch)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, wrld!")
}

func scratch(w http.ResponseWriter, r *http.Request) {
	f := faceimage.ToVector("static/img/orl_faces/1.png")
	averageFaces(5)
	saveVector(f, w, r)
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

func saveVector(face faceimage.FaceVector, w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "faceimage", nil), &face)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Stored a face vector: %v\n", face)

	var face2 faceimage.FaceVector
	if err = datastore.Get(c, key, &face2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Retrieved a face vector: %v\n", face2)

	if face.Width == face2.Width { // can't compare struct with []float64
		fmt.Fprint(w, "Successfully retrieved same vector!\n")
	} else {
		fmt.Fprint(w, "Failed to retrieve same vector!\n")
	}
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

// func saveImage(im image.Image) {
// 	c := appengine.NewContext(r)
// 	blobs, _, err := blobstore.ParseUpload(r)
// 	if err != nil {
// 		serveError(c, w, err)
// 		return
// 	}
// 	file := blobs["file"]
// 	if len(file) == 0 {
// 		c.Errorf("no file uploaded")
// 		http.Redirect(w, r, "/", http.StatusFound)
// 		return
// 	}
// 	http.Redirect(w, r, "/serve/?blobKey="+string(file[0].BlobKey), http.StatusFound)
// }
