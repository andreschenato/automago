package main

import "fmt"

type Automaton struct {
	TransitionTable map[int]map[rune]int
	CurrentState    int
	FinalStates     map[int]bool
	ErrorState      int
	TotalStates     int
}

func NewAutomaton() *Automaton {
	return &Automaton{
		TransitionTable: make(map[int]map[rune]int),
		CurrentState:    0,
		FinalStates:     make(map[int]bool),
		ErrorState:      -1,
		TotalStates:     1,
	}
}

func (a *Automaton) BuildFromWords(words []string) {
	a.TransitionTable = make(map[int]map[rune]int)
	a.FinalStates = make(map[int]bool)
	a.TotalStates = 1
	nextFreeState := 1

	for _, word := range words {
		currentState := 0
		for _, char := range word {
			if a.TransitionTable[currentState] == nil {
				a.TransitionTable[currentState] = make(map[rune]int)
			}

			if next, exists := a.TransitionTable[currentState][char]; exists {
				currentState = next
			} else {
				newState := nextFreeState
				nextFreeState++
				a.TransitionTable[currentState][char] = newState
				currentState = newState
				a.TotalStates = nextFreeState
			}
		}
		a.FinalStates[currentState] = true
	}
}

func (a *Automaton) IsFinal(state int) bool {
	return a.FinalStates[state]
}

func (a *Automaton) GetAlphabet() []rune {
	return []rune("abcdefghijklmnopqrstuvwxyz")
}

func (a *Automaton) ListStates() []int {
	list := make([]int, a.TotalStates)
	for i := 0; i < a.TotalStates; i++ {
		list[i] = i
	}
	return list
}

func (a *Automaton) GetTransitionDisplay(state int, char rune) string {
	if transitions, ok := a.TransitionTable[state]; ok {
		if dest, ok := transitions[char]; ok {
			return fmt.Sprintf("q%d", dest)
		}
	}
	return "-"
}

func (a *Automaton) Accept(word string) bool {
	currentState := 0

	for _, char := range word {
		if transitions, ok := a.TransitionTable[currentState]; ok {
			if nextState, exists := transitions[char]; exists {
				currentState = nextState
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return a.IsFinal(currentState)
}
