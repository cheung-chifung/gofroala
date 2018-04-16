package main

import (
	"net/http"
	"os"

	"log"

	"github.com/graymeta/stow"
	"github.com/graymeta/stow/local"
	"github.com/keekun/gofroala"
)

func main() {

	// Setup stow: https://github.com/graymeta/stow.git
	localConfig := stow.ConfigMap{"path": "example/files"}
	l, err := stow.Dial(local.Kind, localConfig)
	if err != nil {
		log.Fatal(err)
	}
	// rebuild container every time for test
	_ = os.RemoveAll("example/files/test")
	container, err := l.CreateContainer("test")
	if err != nil {
		log.Fatal(err)
	}

	gofroala.Container = container
	gofroala.FileUploadOptions.ItemURLFunc = itemURLFunc
	gofroala.ImageUploadOptions.ItemURLFunc = itemURLFunc

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("example/files"))))
	http.Handle("/froala/", gofroala.NewServeMux("/froala"))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func itemURLFunc(item stow.Item) (string, error) {
	return "/files/test/" + item.Name(), nil
}

func index(w http.ResponseWriter, r *http.Request) {
	body := []byte(`<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <!-- Include external CSS. -->
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.4.0/css/font-awesome.min.css" rel="stylesheet" type="text/css" />
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.25.0/codemirror.min.css">
    <!-- Include Editor style. -->
    <link href="https://cdnjs.cloudflare.com/ajax/libs/froala-editor/2.8.0/css/froala_editor.pkgd.min.css" rel="stylesheet" type="text/css" />
    <link href="https://cdnjs.cloudflare.com/ajax/libs/froala-editor/2.8.0/css/froala_style.min.css" rel="stylesheet" type="text/css" />
  </head>
  <body style="padding: 50px 0">
    <!-- Create a tag that we will use as the editable area. -->
    <!-- You can use a div tag as well. -->
    <textarea></textarea>
    <!-- Include external JS libs. -->
    <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
    <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.25.0/codemirror.min.js"></script>
    <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.25.0/mode/xml/xml.min.js"></script>
    <!-- Include Editor JS files. -->
    <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/froala-editor/2.8.0/js/froala_editor.pkgd.min.js"></script>
    <!-- Initialize the editor. -->
    <script>
      $(function() {
        $('textarea').froalaEditor({
          "toolbarButtons": ["insertImage", "insertLink", "insertTable"],
					imageUploadURL: "/froala/upload_image",
					imageManagerLoadURL: "/froala/load_images",
					imageManagerDeleteURL: "/froala/delete_image",
        });
      });
    </script>
  </body>
</html>`)

	w.Header().Set("Content-Type", "text/html")
	w.Write(body)
}
