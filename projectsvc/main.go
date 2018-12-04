package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Project structure
// contains information on a project item
type Project struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description,omitempty"`
	StartDate   time.Time `json:"from,omitempty"`
	EndDate     time.Time `json:"to,omitempty"`
}

// Global variable containing projects
var projects []Project

// Retreive the project list
func listProjectHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(projects)
}

// get details on a specific project
func getProjectHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	pid, err := strconv.Atoi(params["pid"])
	if err != nil {
		return
	}

	for _, item := range projects {
		if item.ID == pid {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Project{})
}

// add a new project
func addProjectHandler(w http.ResponseWriter, req *http.Request) {
	var project Project
	_ = json.NewDecoder(req.Body).Decode(&project)
	project.ID = len(projects)
	projects = append(projects, project)
	json.NewEncoder(w).Encode(project)
}

//Entry point
func main() {

	projects = append(projects, Project{ID: 1, Code: "8001", Description: "HelloWorld project", StartDate: time.Date(2018, time.Month(12), 3, 0, 0, 0, 0, time.UTC)})
	projects = append(projects, Project{ID: 2, Code: "8002", Description: "SOCGEN", StartDate: time.Date(2017, time.Month(12), 3, 0, 0, 0, 0, time.UTC)})

	router := mux.NewRouter()

	router.HandleFunc("/project", listProjectHandler).Methods("GET")
	router.HandleFunc("/project/{pid}", getProjectHandler).Methods("GET")
	router.HandleFunc("/project", addProjectHandler).Methods("POST")

	log.Println("Starting on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
