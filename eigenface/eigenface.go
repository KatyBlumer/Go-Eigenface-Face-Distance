package eigenface

import "github.com/KatyBlumer/gonum/matrix/mat64"

type FaceVector struct {
	Pixels        []float64
	Width, Height int
}

func Average(faces []FaceVector) FaceVector {
	width := faces[0].Width
	height := faces[0].Height
	avg := make([]float64, width*height)

	for i := 0; i < len(faces); i++ {
		face := faces[i]
		if face.Width != width || face.Height != height {
			return FaceVector{}
		}
		for j := 0; j < width*height; j++ {
			// TODO check what this does to precision
			avg[j] += face.Pixels[j]
		}
	}

	for j := 0; j < width*height; j++ {
		avg[j] = avg[j] / float64(len(faces))
	}
	return FaceVector{
		Width:  width,
		Height: height,
		Pixels: avg,
	}
}

func difference(face1, face2 FaceVector) FaceVector {
	width := face1.Width
	height := face1.Height
	diff := make([]float64, width*height)

	for i := 0; i < width*height; i++ {
		diff[i] += face1.Pixels[i] - face2.Pixels[i]
	}
	return FaceVector{
		Width:  width,
		Height: height,
		Pixels: diff,
	}
}

func Normalize(faces []FaceVector) []FaceVector {
	faceDiffs := make([]FaceVector, len(faces))

	avg := Average(faces)
	for i := 0; i < len(faces); i++ {
		faceDiffs[i] = difference(faces[i], avg)
	}

	return faceDiffs
}

func Eigenfaces(faces []FaceVector) [][]float64 {
	mat := matrix(faces)
	epsilon := 0.01
	small := 0.01
	eigenvalues := mat64.SVD(mat, epsilon, small, true /*wantu*/, false /*wantv*/).U
}

func matrix(faces []FaceVector) [][]float64 {
	mat := make([]float64, len(faces))
	for i := 0; i < height; i++ {
		mat[i] = faces[i].Pixels
	}
	return transpose(mat)
}

func transpose(mat [][]float64) (t [][]float64) {
	height := len(mat)
	width := len(mat[0])

	t = makeMat(height, width)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			t[i][j] = mat[j][i]
		}
	}
	return
}

func makeMat(width, height int) (mat [][]float64) {
	mat = make([][]float64, height)
	for i := 0; i < height; i++ {
		mat[i] = make([]float64, width)
	}
	return
}
