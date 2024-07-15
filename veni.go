package veni

import(
	"net/http"
)

func Load(handler http.Handler) http.Handler{
	//right now do nothing
	return handler;
}
