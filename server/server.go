package main

import (
	"html/template"
	"github.com/SennaSemakula/server/wiki"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := wiki.LoadPage(title)

	// Redirect user if page does not exist
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

func renderTemplate(w http.ResponseWriter, file string, p *wiki.Page) {
	tmpl, _ := template.ParseFiles(file + ".html")

	tmpl.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	p, err := wiki.LoadPage(title)

	// Create an empty page with a title if file does not exist
	if err != nil {
		p = &wiki.Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]

	body := r.FormValue("body")

	p := &wiki.Page{Title: title, Body: []byte(body)}

	p.Save()

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	// http.HandleFunc("/", handle)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	// Always returns an error. We wrap it in a log fatal so that we can view that error
	log.Fatal(http.ListenAndServe(":8080", nil))
}
