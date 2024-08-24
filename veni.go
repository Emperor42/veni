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
	if r.URL.Path != "/call" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Call Complete!")
}
