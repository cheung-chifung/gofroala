package gofroala

import "net/http"

func NewServeMux(prefix string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(prefix+"/upload_file", UploadFileHandler)
	mux.Handle(prefix+"/upload_file_validation", UploadFileValidationHandler)
	mux.Handle(prefix+"/upload_image", UploadImageHandler)
	mux.Handle(prefix+"/upload_image_validation", UploadImageValidationHandler)
	mux.Handle(prefix+"/upload_image_resize", UploadImageResizeHandler)
	mux.Handle(prefix+"/delete_file", DeleteFileHandler)
	mux.Handle(prefix+"/delete_image", DeleteImageHandler)
	mux.Handle(prefix+"/load_images", LoadImagesHandler)

	return mux
}

var UploadFileHandler = http.HandlerFunc(uploadFile)
var UploadFileValidationHandler = http.HandlerFunc(uploadFileValidation)
var UploadImageHandler = http.HandlerFunc(uploadImage)
var UploadImageValidationHandler = http.HandlerFunc(uploadImageValidation)
var UploadImageResizeHandler = http.HandlerFunc(uploadImageResize)
var DeleteFileHandler = http.HandlerFunc(deleteFile)
var DeleteImageHandler = http.HandlerFunc(deleteImage)
var LoadImagesHandler = http.HandlerFunc(loadImages)
