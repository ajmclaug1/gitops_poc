package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type page struct {
	AppType string
	Version string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! This is the Third version of me!")

}

func indexHandler2(w http.ResponseWriter, r *http.Request) {
	p := page{AppType: "Func1", Version: "Version 1.0"}
	t := template.Must(template.ParseFiles("./html/static.html"))

	err := t.Execute(w, p)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("Basic web server is starting on port 8080...")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ver", indexHandler2)
	http.ListenAndServe(":8080", nil)
}
