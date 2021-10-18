package main

import (
	"github.com/afocus/captcha"
)

func main() {
	cap := captcha.New()
	cap.SetFont("/home/jonathan/go/pkg/mod/github.com/afocus/captcha@v0.0.0-20191010092841-4bd1f21c8868/examples/comic.ttf")

	cap.Create(5, captcha.ALL)

}
