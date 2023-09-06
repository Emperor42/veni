// Copyright 2010 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style

// license that can be found in the LICENSE file.


//go:build ignore


package main


import (

	"html/template"

	"log"

	"net/http"

	"os"

	"regexp"

)


type Page struct {

	Title string

	Body  []byte

}


func (p *Page) save() error {

	filename := "../"+p.Title + ".txt"

	return os.WriteFile(filename, p.Body, 0600)

}


func loadPage(title string) (*Page, error) {

	filename := "../"+title + ".txt"

	body, err := os.ReadFile(filename)

	if err != nil {

		return nil, err

	}

	return &Page{Title: title, Body: body}, nil

}


func viewHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)

	if err != nil {

		http.Redirect(w, r, "/edit/"+title, http.StatusFound)

		return

	}

	renderTemplate(w, "view", p)

}


func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)

	if err != nil {

		p = &Page{Title: title}

	}

	renderTemplate(w, "edit", p)

}


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

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)

	if err != nil {

		http.Redirect(w, r, "/edit/"+title, http.StatusFound)

		return

	}

	renderTemplate(w, "view", p)

}

func indexHandler(w http.ResponseWriter, r *http.Request, title string) {

	return http.HandleFunc("/VENI/view/index", makeHandler(viewHandler))

}

var templates = template.Must(template.ParseFiles("edit.html", "view.html", "index.html"))


func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {

	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

//handlers for HTTP functions from API, these are specified in modules outside of the main server

func connectHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return

	}

	//COMPLETE STUB CONNECT

}

func deleteHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	//COMPLETE STUB delete

}

func getHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	//COMPLETE STUB get

}

func headHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	//COMPLETE STUB head

}

func optionsHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	//COMPLETE STUB options

}

func patchHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	//COMPLETE STUB patch

}

func postHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	//COMPLETE STUB post

}

func putHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	//COMPLETE STUB put

}

func traceHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	//COMPLETE STUB trace

}

//The path for the VENI system directly, hardnend compared to API path
var validVENIPath = regexp.MustCompile("^/VENI/(edit|save|view)/([a-zA-Z0-9]+)$")

var validAPIPath = regexp.MustCompile("^/API/(get|post|head|put|delete|connect|options|trace|patch)/([a-zA-Z0-9]+)$")


func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		m := validVENIPath.FindStringSubmatch(r.URL.Path)
		aux := validAPIPath.FindStringSubmatch(r.URL.Path)

		if m == nil && aux == nil{

			http.NotFound(w, r)

			return

		}

		fn(w, r, m[2])

	}

}


func main() {
	//index

	http.HandleFunc("/", indexHandler);

	//core VENI template view, edit and save (based on https://go.dev/doc/articles/wiki/)

	http.HandleFunc("/VENI/view/", makeHandler(viewHandler))

	http.HandleFunc("/VENI/edit/", makeHandler(editHandler))

	http.HandleFunc("/VENI/save/", makeHandler(saveHandler))

	//API functionality

	http.HandleFunc("/API/connect/", makeHandler(connectHandler))

	http.HandleFunc("/API/delete/", makeHandler(deleteHandler))

	http.HandleFunc("/API/get/", makeHandler(getHandler))

	http.HandleFunc("/API/head/", makeHandler(headHandler))

	http.HandleFunc("/API/options/", makeHandler(optionsHandler))

	http.HandleFunc("/API/patch/", makeHandler(patchHandler))

	http.HandleFunc("/API/post/", makeHandler(postHandler))

	http.HandleFunc("/API/put/", makeHandler(putHandler))

	http.HandleFunc("/API/trace/", makeHandler(traceHandler))




	log.Fatal(http.ListenAndServe(":8080", nil))

}