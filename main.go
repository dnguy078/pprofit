package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^(.+)@golang.org$")
	path := r.URL.Path[1:]
	match := re.FindAllStringSubmatch(path, -1)

	if match != nil {
		fmt.Fprintln(w, "$$$")
		return
	}

	fmt.Fprintln(w, "Hello, World!")
}
