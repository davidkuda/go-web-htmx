# Go


### An idea to instantiate HTMX templates: a set of templates without base

```go

type application struct {
	templateCache     map[string]*template.Template
	templateCacheHTMX map[string]*template.Template
}


// for HTMX, you will not need the base.tmpl.html template.
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

		if includeBase {
			// with base.tmpl.html
			N := 1 + len(partials) + 1
			files = make([]string, N)
			files[0] = "./ui/html/base.tmpl.html"
			for i, partial := range partials {
				files[i+1] = partial
			}
			files[N-1] = page
		} else {
			// without base.tmpl.html
			N := 1 + len(partials)
			files = make([]string, N)
			files[0] = page
			for i, partial := range partials {
				files[i+1] = partial
			}
		}

		tmpl := template.New("base").Funcs(funcs)
		t, err := tmpl.ParseFiles(files...)
		if err != nil {
			return nil, fmt.Errorf("Error parsing template files: %s", err.Error())
		}

		cache[name] = t
	}

	return cache, nil
}

func main() {
	addr := flag.String("addr", ":8875", "HTTP network address")

	app := &application{}

	app.templateCache, err = newTemplateCache(true)
	if err != nil {
		log.Fatalf("could not initialise templateCache: %v\n", err)
	}

	app.templateCacheHTMX, err = newTemplateCache(false)
	if err != nil {
		log.Fatalf("could not initialise templateCache: %v\n", err)
	}

	log.Printf("Starting web server, listening on %s", *addr)
	err = http.ListenAndServe(*addr, app.routes())
	log.Fatal(err)
}
```
