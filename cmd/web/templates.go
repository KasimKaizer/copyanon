package main

import (
	"path/filepath"
	"text/template"

	"github.com/KasimKaizer/copyanon/internal/models"
)

type TemplateData struct {
	Gist  models.Gist
	Gists []models.Gist
}

func NewTemplateCache() (map[string]*template.Template, error) {

	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob("ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.ParseFiles("ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
