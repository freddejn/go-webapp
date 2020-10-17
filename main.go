package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	templatesPath = "templates/"
)

func main() {
	port := os.Getenv("PORT")
	print("running")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello</h1>"))
}
