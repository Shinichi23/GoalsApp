package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Goal struct {
	Title string
	// Number int
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	tmpl := template.Must(template.ParseFiles("index.html"))

	goals := map[string][]Goal{
		"Goals": {
			{"Hello World"},
			{"Hello Mars"},
			{"Hello Moon"},
		},
	}

	tmpl.Execute(w, goals)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/add-goal/", AddGoal)

	fmt.Println("App running ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func AddGoal(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("addgoal")
	fmt.Println(r.Method)
	fmt.Println(title)
	test := fmt.Sprintf("<li>%s</li>", title)
	tmpl, _ := template.New("t").Parse(test)
	tmpl.Execute(w, nil)
}
