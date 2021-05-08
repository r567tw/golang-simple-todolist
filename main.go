package main

import (
	"html/template"
	"log"
	"net/http"
)

type TemplateData struct {
	Name string
}

type Todo struct {
	Task      string
	Completed bool
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")

	if err != nil {
		log.Fatal(err)
	}
	data := TemplateData{Name: "Jimmy"}
	err = html.Execute(writer, data)

	if err != nil {
		log.Fatal(err)
	}
}

func todoCreateHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")

	if err != nil {
		log.Fatal(err)
	}
	data := TemplateData{Name: "Jimmy"}
	err = html.Execute(writer, data)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/todo/create", todoCreateHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
