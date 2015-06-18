package web

import (
	"fmt"
	"github.com/KatyBlumer/Go-Eigenface-Face-Distance/faceimage"
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
	fmt.Fprint(w, "Done!")
}
