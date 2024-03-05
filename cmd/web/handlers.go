package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KasimKaizer/copyanon/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	gists, err := app.gists.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	data := TemplateData{
		Gists: gists,
	}
	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) gistCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
	}

	// Temp data
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.gists.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/gist/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) gistView(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || num < 1 {
		app.notFound(w)
		return
	}

	gist, err := app.gists.Get(num)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		}
		app.serverError(w, r, err)
		return
	}
	data := TemplateData{
		Gist: gist,
	}
	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}
