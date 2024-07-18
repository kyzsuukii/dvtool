package config

import "github.com/gin-contrib/multitemplate"

func ViewRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("login", "templates/layout.html", "templates/login.html")
	r.AddFromFiles("index", "templates/layout.html", "templates/index.html")
	r.AddFromFiles("output", "templates/layout.html", "templates/output.html")

	return r
}
