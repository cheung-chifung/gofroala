package gofroala

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"path"

	"github.com/graymeta/stow"
)

func upload(w http.ResponseWriter, r *http.Request, opts *UploadOptions) {
	file, header, err := r.FormFile(opts.FieldName)
	if err != nil {
		return
	}
	defer file.Close()

	ext := path.Ext(header.Filename)
	filename := generateFilename(ext)

	item, err := Container.Put(filename, file, header.Size, nil)
	if err != nil {
		return
	}
	url, err := opts.ItemURLFunc(item)
	if err != nil {
		return
	}
	jsonRespond(w, map[string]interface{}{
		"link": url,
	})
}

type listItem struct {
	URL   string `json:"url"`
	Thumb string `json:"thumb"`
	Name  string `json:"name"`
}

func list(w http.ResponseWriter, r *http.Request, opts *UploadOptions) {
	stepSize := 500
	res := make([]*listItem, 0, stepSize)
	err := stow.Walk(Container, stow.NoPrefix, stepSize, func(item stow.Item, err error) error {
		if err != nil {
			return err
		}
		url, err := opts.ItemURLFunc(item)
		if err != nil {
			return err
		}
		res = append(res, &listItem{
			URL:   url,
			Thumb: item.URL().String(),
			Name:  item.Name(),
		})
		return nil
	})
	if err != nil {
		return
	}
	jsonRespond(w, res)
}

func generateFilename(ext string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes) + ext
}

func jsonRespond(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
