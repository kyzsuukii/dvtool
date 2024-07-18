package config

import "github.com/gin-contrib/multitemplate"

func ViewRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("login", "templates/layout.html", "templates/login.html")
	r.AddFromFiles("index", "templates/layout.html", "templates/index.html")

	return r
}
