package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var automaton *Automaton

func main() {
	automaton = NewAutomaton()

	automaton.BuildFromWords([]string{"testing", "nesting"})

	http.HandleFunc("/", index)

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, automaton)
}
