package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Goal struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var goals []Goal
var goalsMux sync.Mutex

func main() {
	http.HandleFunc("/goals", allGoals)
	http.HandleFunc("/add", addGoal)
	http.HandleFunc("/update", upGoal)
	http.HandleFunc("/delete", delGoal)

	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("Running App ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func allGoals(w http.ResponseWriter, r *http.Request) {

	goalsMux.Lock()
	defer goalsMux.Unlock()

	goalsJson, err := json.Marshal(goals)
	if err != nil {
		http.Error(w, "Error encoding goals", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(goalsJson)
}

func addGoal(w http.ResponseWriter, r *http.Request) {

	goalsMux.Lock()
	defer goalsMux.Unlock()

	var newGoal Goal
	err := json.NewDecoder(r.Body).Decode(&newGoal)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	newGoal.ID = len(goals) + 1
	goals = append(goals, newGoal)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newGoal)
}

func upGoal(w http.ResponseWriter, r *http.Request) {

	goalsMux.Lock()
	defer goalsMux.Unlock()

	var updatedGoal Goal
	err := json.NewDecoder(r.Body).Decode(&updatedGoal)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	for i, goal := range goals {
		if goal.ID == updatedGoal.ID {
			goals[i] = updatedGoal
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedGoal)
			return
		}
	}

	http.Error(w, "Goal not found", http.StatusNotFound)
}

func delGoal(w http.ResponseWriter, r *http.Request) {

	goalsMux.Lock()
	defer goalsMux.Unlock()

	var targetID int
	err := json.NewDecoder(r.Body).Decode(&targetID)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	for i, goal := range goals {
		if goal.ID == targetID {
			goals = append(goals[:i], goals[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}

	}

	http.Error(w, "Goal not found", http.StatusNotFound)
}
