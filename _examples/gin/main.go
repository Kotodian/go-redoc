package main

import (
	_ "embed"
	"github.com/Kotodian/go-redoc"
	ginredoc "github.com/Kotodian/go-redoc/gin"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

//go:embed openapi.json
var spec string

func main() {
	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		//SpecFile:    "./openapi.json",
		SpecPath: "/openapi.json",
		DocsPath: "/docs",
		Spec:     spec,
	}

	r := gin.New()
	r.Use(favicon.New("./_examples/gin/favicon.ico"))
	r.Use(ginredoc.New(doc))
	r.StaticFile("/favicon", "./favicon.ico")
	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Run(":8000"))
}
