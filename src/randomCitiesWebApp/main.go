package main

import (
	"net/http"
    "html/template"
    "RandomCities"
)

// Available templates.
var templates = template.Must(template.ParseFiles(
                "view/cities.html",
))

// Struct to save a webpage.
type Page struct{
	Title string
    RandomCities map[string]string
}

// Render templates.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Handler to show the random cities.
func viewHandler(w http.ResponseWriter, r *http.Request) {
    // File name:
    inputFileName := "src/txt/US_Cities.txt"
    // Gets randomCities from the file:
    randomCities := RandomCities.RandomCities(inputFileName)
    //Creates page:
    page := Page{Title: "RandomCities",  RandomCities: randomCities}
    // Show Page:
    renderTemplate(w, "cities", &page)
}

func main() {
    http.HandleFunc("/", viewHandler)
    http.ListenAndServe(":8080", nil)
}