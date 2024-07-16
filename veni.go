package veni

import(
	"net/http"
)

func Load(fs http.FileSystem) http.Handler{
	return http.FileServer(fs)
}
