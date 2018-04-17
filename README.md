# gofroala (WIP)

Gofroala is a Go library for implementing the [Froala WYSIWYG HTML Editor](https://github.com/froala/wysiwyg-editor) server side with Amazon S3, Google Cloud Storage etc.

## Installation

### Requirements
* Go 1.9+
* [Stow](https://github.com/graymeta/stow)

```go
go get github.com/graymeta/stow
go get github.com/keekun/gofroala
```

## Usage

Please read the code [example](example/main.go).


## Development

```
go run example/main.go
```

and open

```
http://localhost:8080/
```

with [Refresh](https://github.com/markbates/refresh), example code would be rebuid and re-run automatically.

```
go get github.com/markbates/refresh
refresh init
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)

## Project status

Work in process
