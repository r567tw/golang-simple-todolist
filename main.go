package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"fmt"
)

type TodoList struct {
	Todos []string
}

// type Todo struct {
// 	Task string
// 	Completed bool
// }

func getTodos(fileName string) []string {
	var todos []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		todos = append(todos, scanner.Text())
	}
	return todos
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")
	todos := getTodos("todolist.txt")
	if err != nil {
		log.Fatal(err)
	}

	todolist := TodoList{
		Todos:  todos,
	}

	err = html.Execute(writer, todolist)

	if err != nil {
		log.Fatal(err)
	}
}

func todoCreateHandler(writer http.ResponseWriter, request *http.Request) {
	task := request.FormValue("task")

	file, _ := os.OpenFile("todolist.txt", os.O_WRONLY | os.O_APPEND | os.O_CREATE , os.FileMode(0600))
	fmt.Fprintln(file, task)
	defer file.Close()
	http.Redirect(writer, request, "/", http.StatusFound)
}

func main() {

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/todo/create", todoCreateHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
