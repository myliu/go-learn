package main

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/russross/blackfriday"
)

func main() {
	m := martini.Classic()

	m.Post("/generate", func(r *http.Request) []byte {
		body := r.FormValue("body")
		return blackfriday.MarkdownBasic([]byte(body))
	})

	m.Run()
}
