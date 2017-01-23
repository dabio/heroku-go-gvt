package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello, World!\n\nThis is a demo on how to use gvt on Heroku.\ngithub.com/dabio/heroku-go-gvt")
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Hello, "+p.ByName("name"))
}

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/hello/:name", hello)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
