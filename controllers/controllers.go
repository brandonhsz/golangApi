package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int
	Name    string
	Content string
}

type allTask []task

var tasks = allTask{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some content",
	},
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid Task")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	fmt.Fprintf(w, "Tarea %v eliminada", taskID)
}

func GetAnyTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "ID INVALIDO err")
		return
	}

	for _, task := range tasks {
		if task.ID == taskId {
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}

	}

}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}
