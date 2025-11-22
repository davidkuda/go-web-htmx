package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type templateData struct {
	Title    string
	Path     string
	RootPath string
	User     User
	Error    Error
}

type User struct {
	Authenticated bool
	ID            string
	Email         string
}

type Error struct {
	HTTPStatusCode int
	HTTPStatusText string
	Method         string
	Path           string
}

func newError(r *http.Request, errorCode int) Error {
	return Error{
		HTTPStatusCode: errorCode,
		HTTPStatusText: http.StatusText(errorCode),
		Method:         r.Method,
		Path:           r.URL.Path,
	}
}

func (app *application) newTemplateData(r *http.Request) templateData {

	var rootPath, title string
	i := 1
	for i < len(r.URL.Path) && r.URL.Path[i] != '/' {
		i++
	}
	rootPath = r.URL.Path[0:i]
	title = strings.ToTitle(r.URL.Path[1:i])

	return templateData{
		Title:    title,
		RootPath: rootPath,
		Path:     r.URL.Path,
		User:     User{},
	}
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data *templateData) {
	var err error

	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("couldn't find template \"%s\" in app.templateCache", page)
		app.serverError(w, r, err)
		return
	}

	buf := bytes.Buffer{}

	isHTMX := r.Header.Get("HX-Request") == "true"
	if isHTMX {
		err = ts.ExecuteTemplate(&buf, "main", data)
	} else {
		err = ts.ExecuteTemplate(&buf, "base", data)
	}
	if err != nil {
		errMsg := fmt.Errorf("error executing template %s: %s", page, err.Error())
		app.serverError(w, r, errMsg)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func newTemplateCache(includeBase bool) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	funcs := template.FuncMap{
		"formatDate": formatDate,
	}

	pages, err := filepath.Glob("./ui/html/*.tmpl.html")
	if err != nil {
		return nil, fmt.Errorf("failed filepath.Glob for pages: %v", err)
	}

	partials, err := filepath.Glob("./ui/html/partials/*.tmpl.html")
	if err != nil {
		return nil, fmt.Errorf("failed filepath.Glob for partials: %v", err)
	}

	for _, page := range pages {
		name := filepath.Base(page)

		var files []string

		N := 1 + len(partials) + 1
		files = make([]string, N)
		files[0] = "./ui/html/base.tmpl.html"
		for i, partial := range partials {
			files[i+1] = partial
		}
		files[N-1] = page

		tmpl := template.New("base").Funcs(funcs)
		t, err := tmpl.ParseFiles(files...)
		if err != nil {
			return nil, fmt.Errorf("Error parsing template files: %s", err.Error())
		}

		cache[name] = t
	}

	return cache, nil
}

func formatDate(t time.Time) string {
	return t.Format("January 2, 2006")
}
