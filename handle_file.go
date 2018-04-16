package gofroala

import (
	"net/http"
)

func uploadFileValidation(w http.ResponseWriter, r *http.Request) {
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	upload(w, r, FileUploadOptions)
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
}
