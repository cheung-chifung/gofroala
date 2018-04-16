package gofroala

import "net/http"

func uploadImage(w http.ResponseWriter, r *http.Request) {
	upload(w, r, ImageUploadOptions)
}

func uploadImageValidation(w http.ResponseWriter, r *http.Request) {
}

func uploadImageResize(w http.ResponseWriter, r *http.Request) {
}

func deleteImage(w http.ResponseWriter, r *http.Request) {
}

func loadImages(w http.ResponseWriter, r *http.Request) {
	list(w, r, ImageUploadOptions)
}
