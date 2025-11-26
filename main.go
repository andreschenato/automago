package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"unicode"
)

var automaton *Automaton

type ProcessResponse struct {
	State   int    `json:"state"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	automaton = NewAutomaton()

	automaton.BuildFromWords([]string{"testing", "nesting"})

	http.HandleFunc("/", index)
	http.HandleFunc("/api/configure", configure)
	http.HandleFunc("/api/process", process)

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	automaton.Reset()

	tmpl.Execute(w, automaton)
}

func configure(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		input := r.FormValue("words")

		splitter := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		fields := strings.FieldsFunc(input, splitter)

		var cleanWords []string
		for _, p := range fields {
			cleanWords = append(cleanWords, strings.ToLower(p))
		}

		automaton.BuildFromWords(cleanWords)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func process(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	resp := ProcessResponse{}

	if action == "validate" {
		fullToken := r.FormValue("token")
		accepted := automaton.SimulateValidation(fullToken)

		if accepted {
			resp.Status = "accepted"
			resp.Message = "ACCEPTED"
		} else {
			resp.Status = "rejected"
			resp.Message = "REJECTED"
		}

		automaton.Reset()
		resp.State = 0

	} else {
		currentToken := r.FormValue("token")

		if currentToken == "" {
			resp.State = 0
			resp.Message = "Aguardando..."
		} else {
			finalState := automaton.GetStateForString(currentToken)
			resp.State = finalState

			if resp.State == -1 {
				resp.Message = "ERRO"
			} else {
				resp.Message = fmt.Sprintf("State q%d", resp.State)
			}
		}

		resp.Status = "processing"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
