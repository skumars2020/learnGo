package main

import (
	"net/http"

	"github.com/skumars2020/learnGo/controller"
)

func main() {
	mux := controller.Regester()
	http.ListenAndServe(":3000", mux)
}
