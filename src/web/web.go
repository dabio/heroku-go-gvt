package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello, World!\n\nThis is a demo on how to use gvt on Heroku.\ngithub.com/dabio/heroku-go-gvt")
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Hello, "+p.ByName("name"))
}

func main() {
	port := flag.String("port", "8080", "listening port")
	flag.Parse()

	r := httprouter.New()
	r.GET("/", index)
	r.GET("/hello/:name", hello)

	log.Fatal(http.ListenAndServe(":"+*port, r))
}
