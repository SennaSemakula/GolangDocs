package main

import (
	"html/template"
	"github.com/SennaSemakula/server/wiki"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := wiki.LoadPage(title)

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	p, err := wiki.LoadPage(title)

	// Create an empty page with a title if file does not exist
	if err != nil {
		p = &wiki.Page{Title: title}
	}

	fmt.Fprintf(w, "<h1> Editing %s </h1>"+
					"<form action=\"save/ %s\" method=\"POST\">"+
					"<textarea name=\"body\">%s</textarea><br>"+
					"<input type=\"submit\" value=\"Save\">"+
					"</form>", p.Title, p.Title, p.Body)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]

	fmt.Println(r.Body)

	// wiki.SaveF
	fmt.Fprintf(w, "<h1>Saving file %s</h1>", title)
}

func main() {
	// http.HandleFunc("/", handle)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	// Always returns an error. We wrap it in a log fatal so that we can view that error
	log.Fatal(http.ListenAndServe(":8080", nil))
}
