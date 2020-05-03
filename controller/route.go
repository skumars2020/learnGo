package controller

import (
	"net/http"
)

func Regester() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	return mux
}
