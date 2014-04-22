package main

import (
	"io/ioutil"
	"net/http"
    "html/template"
    "RandomCities"
)

var templates = template.Must(template.ParseFiles(
                "view/cities.html",
))

type Page struct{
	Title string
	Body []byte
    RandomCities map[string]string
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    inputFileName := "input/US_Cities"
    randomCities := RandomCities.RandomCities(inputFileName)

    page := Page{Title: "RandomCities", Body: []byte("null"), RandomCities: randomCities}
    
    renderTemplate(w, "cities", &page)
}

func main() {
    http.HandleFunc("/", viewHandler)
    http.ListenAndServe(":8080", nil)
}