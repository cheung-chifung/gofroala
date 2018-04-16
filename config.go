package gofroala

import "github.com/graymeta/stow"

var (
	FileUploadOptions  *UploadOptions
	ImageUploadOptions *UploadOptions
	Container          stow.Container
)

func init() {
	FileUploadOptions = DefaultFileUploadOptions
	ImageUploadOptions = DefaultImageUploadOptions
}

type UploadOptions struct {
	FieldName   string
	Validation  *UploadValidationOptions
	Resize      string
	ItemURLFunc func(item stow.Item) (string, error)
}

type UploadValidationOptions struct {
	AllowedExts      []string
	AllowedMimeTypes []string
}

var DefaultFileUploadOptions = &UploadOptions{
	FieldName: "file",
	Validation: &UploadValidationOptions{
		AllowedExts: []string{
			"txt",
			"pdf",
			"doc",
		},
		AllowedMimeTypes: []string{
			"text/plain",
			"application/msword",
			"application/x-pdf",
			"application/pdf",
		},
	},
	ItemURLFunc: func(item stow.Item) (string, error) {
		return item.URL().String(), nil
	},
}

var DefaultImageUploadOptions = &UploadOptions{
	FieldName: "file",
	Validation: &UploadValidationOptions{
		AllowedExts: []string{
			"gif",
			"jpeg",
			"jpg",
			"png",
			"svg",
			"blob",
		},
		AllowedMimeTypes: []string{
			"image/gif",
			"image/jpeg",
			"image/pjpeg",
			"image/x-png",
			"image/png",
			"image/svg+xml",
		},
	},
	ItemURLFunc: func(item stow.Item) (string, error) {
		return item.URL().String(), nil
	},
	// string resize param from http://docs.wand-py.org/en/0.4.3/guide/resizecrop.html#transform-images
	// Examples: '100x100', '100x100!'. Find more on http://www.imagemagick.org/script/command-line-processing.php#geometry
	Resize: "",
}
