package veni

import (
	"fmt"
	"net/http"
)

type VeniContext struct {
	Name string
}

func (v *VeniContext) ProcessHeader() {
	fmt.Println("temp")
}

func (v *VeniContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() != "/call" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, v.Name)
	}

	fmt.Fprintf(w, "Call Complete!")
}
