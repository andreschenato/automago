package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var automaton *Automaton

type ProcessResponse struct {
	State   int    `json:"state"`
	Message string `json:"message"`
}

func main() {
	automaton = NewAutomaton()

	automaton.BuildFromWords([]string{"testing", "nesting"})

	http.HandleFunc("/", index)
	http.HandleFunc("/process", process)
	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		automaton.Reset()
	})

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

func process(w http.ResponseWriter, r *http.Request) {
	charStr := r.FormValue("char")
	resp := ProcessResponse{}

	if len(charStr) > 0 {
		r := rune(charStr[0])
		automaton.Step(r)
	}

	resp.State = automaton.CurrentState
	resp.Message = fmt.Sprintf("Current State: q%d", resp.State)
	if resp.State == -1 {
		resp.Message = "Error State"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
