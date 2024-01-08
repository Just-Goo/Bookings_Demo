package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Just-Goo/Bookings_Demo/pkg/config"
	"github.com/Just-Goo/Bookings_Demo/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData  {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		// get the template cache from the 'AppConfig'
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// Retrieve the requested template from template cache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData)

	_ = template.Execute(buf, templateData)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

//==== A complex template cache ====//

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{} // create a map for caching the templates

	// get all of the files named *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html") // Get all the files ending with '.page.html' from the 'templates' folder
	if err != nil {
		return myCache, err
	}

	// range through all files ending with '*.page.html'
	for _, page := range pages {
		// Each page contains the full file path, but we only need the name of the file. So we use filepath.Base() to get just the file name minus the full path to it
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page) //
		if err != nil {
			return myCache, nil
		}

		matches, err := filepath.Glob("./templates/*.layout.html") // Get all the files ending with '.layout.html' from the 'templates' folder
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 { // If there are any available 'layouts' in templates folder
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateSet
	}
	return myCache, nil

}

//==== A simple template cache ====//

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have the template in our map
// 	_, inMap := tc[t]
// 	if !inMap {
// 		log.Println("creating template and adding to cache")
// 		// create the template
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	} else {
// 		// template already in cache

// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}

// 	// Parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	// Add template to cache [map]
// 	tc[t] = tmpl

// 	return nil
// }
