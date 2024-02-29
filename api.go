package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks = getAllTasksRepository()
	json.NewEncoder(w).Encode(tasks)
}

func getTaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID := 0
	fmt.Sscanf(params["id"], "%d", &taskID)
	var task = getTaskByIdRepository(taskID)
	if task.ID > 0 {
		json.NewEncoder(w).Encode(task)
		return
	}
	json.NewEncoder(w).Encode(&Task{})
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	var task = createTaskRepository(newTask)
	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID := 0
	fmt.Sscanf(params["id"], "%d", &taskID)
	var updatedTask Task
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)
	var task = updateTaskRepository(taskID, updatedTask)
	if task.ID > 0 {
		json.NewEncoder(w).Encode(task)
		return
	}
	json.NewEncoder(w).Encode(&Task{})
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID := 0
	fmt.Sscanf(params["id"], "%d", &taskID)
	var isDeleted = deleteTaskRepository(taskID)
	if isDeleted {
		json.NewEncoder(w).Encode("{\"message\": \"Task deleted\"}")
		return
	} else {
		json.NewEncoder(w).Encode("{\"message\": \"Task deleted\"}")
		return
	}
}
