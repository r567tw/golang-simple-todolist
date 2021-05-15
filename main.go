package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Task      string
	Completed bool
}

type templateData struct {
	Name string
}

var todos []Todo

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")
	data := templateData{Name: "Jimmy"}
	if err != nil {
		log.Fatal(err)
	}
	err = html.Execute(writer, data)

	if err != nil {
		log.Fatal(err)
	}
}

func todoCreateHandler(writer http.ResponseWriter, request *http.Request) {
	task := request.FormValue("task")
	println(task)
	todos[0] = Todo{Task: task, Completed: false}
	println(todos[0].Task)
}

func main() {
	todos = make([]Todo, 5)

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/todo/create", todoCreateHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
