package redoc

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	_ "embed"
)

// ErrSpecNotFound error for when spec file not found
var ErrSpecNotFound = errors.New("spec not found")

// Redoc configuration
type Redoc struct {
	Spec        string
	DocsPath    string
	SpecPath    string
	SpecFile    string
	Title       string
	Description string
}

// HTML represents the redoc index.html page
//go:embed assets/index.html
var HTML string

// JavaScript represents the redoc standalone javascript
//go:embed assets/redoc.standalone.js
var JavaScript string

////go:embed assets/favicon.ico
//var icon string

// Body returns the final html with the js in the body
func (r Redoc) Body() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	tpl, err := template.New("redoc").Parse(string(HTML))
	if err != nil {
		return nil, err
	}

	if err = tpl.Execute(buf, map[string]string{
		"body":        string(JavaScript),
		"title":       r.Title,
		"url":         r.SpecPath,
		"description": r.Description,
		//"icon":        "favicon.ico",
	}); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Handler sets some defaults and returns a HandlerFunc
func (r Redoc) Handler() http.HandlerFunc {
	data, err := r.Body()
	if err != nil {
		panic(err)
	}
	var spec []byte
	if r.Spec != "" {
		spec = []byte(r.Spec)
	} else {
		specFile := r.SpecFile
		if specFile == "" {
			panic(ErrSpecNotFound)
		}
		spec, err = ioutil.ReadFile(specFile)
		if err != nil {
			panic(err)
		}
	}

	specPath := r.SpecPath
	if specPath == "" {
		specPath = "./openapi.json"
	}

	docsPath := r.DocsPath
	if docsPath == "" {
		docsPath = "/"
	}

	return func(w http.ResponseWriter, req *http.Request) {
		method := strings.ToLower(req.Method)

		if method != "get" && method != "head" {
			return
		}

		switch req.URL.Path {
		case docsPath:
			w.WriteHeader(200)
			w.Header().Set("content-type", "text/html")
			w.Write(data)
		case specPath:
			w.WriteHeader(200)
			w.Header().Set("content-type", "application/json")
			w.Write(spec)
		default:
		}
	}
}
