package main

import (
	"github.com/hookrock/gooo"
)

func main() {
	r := gooo.Default()
	r.GET("/", func(c *gooo.Context) {
		c.View("index.tmpl", nil)
	})
	r.Run()
}
