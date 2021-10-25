package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Page Wiki contents
type Page struct {
	Title string
	Body  []byte // means "a byte slice.", body element is a []byte rather than string because that is the type expected by io libraries will use.
}

// save method will save the Page's Body to a next file.
func (p *Page) save() error {
	filename := fmt.Sprintf("%s.txt", p.Title)

	return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage constructs the file name from the title parameter,reads the file's contents into a new variable body.
func loadPage(title string) (*Page, error) {
	filename := fmt.Sprintf("%s.txt", title)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

/*
func main() {
	p1 := &Page{Title: "LastPage", Body: []byte("This is a sample Page")}
	_ = p1.save()
	p2, _ := loadPage("LastPage")
	fmt.Println(string(p2.Body))
}
*/

// Using net/http to serve wiki page

// templates Template caching
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// validPath validate the title with a regular expression
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type WikiHandler func(w http.ResponseWriter, r *http.Request, title string)

func makeHandler(f WikiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		f(w, r, m[2])
	}
}

/*
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}

	return m[2], nil // The Title is the second subexpression
}
*/

// viewHandler allow users to view a wiki page, handle URLs prefixed with "/views/"
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		// handling non-existent pages
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// editHandler loads the edit page.
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// saveHandler handle the submission of forms located on the edit pages.
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// renderTemplate use the template, instead of the hard-coded HTML
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	/*
		t, err := template.ParseFiles(tmpl + ".html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/
	// Using template caching
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
