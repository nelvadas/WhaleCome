package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	. "github.com/nelvadas/WhaleCome/projectsvc/dao"
	. "github.com/nelvadas/WhaleCome/projectsvc/models"
	"gopkg.in/mgo.v2/bson"
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

var dao = ProjectDAO{}

// Retreive the project list
func listProjectHandler(w http.ResponseWriter, req *http.Request) {
	projects, err := dao.FindAll()
	if err == nil {
		json.NewEncoder(w).Encode(projects)
	} else {
		json.NewEncoder(w).Encode([]ProjectDTO{})
	}

}

// get details on a specific project
func getProjectHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["pid"]
	log.Printf("Looking for id with id=%v", id)
	projectItem, err := dao.FindProjectByID(id)
	if err == nil {
		json.NewEncoder(w).Encode(projectItem)
	} else {
		json.NewEncoder(w).Encode(ProjectDTO{})
	}

}

// add a new project
func addProjectHandler(w http.ResponseWriter, req *http.Request) {
	var project ProjectDTO
	_ = json.NewDecoder(req.Body).Decode(&project)
	project.ID = bson.NewObjectId()

	log.Printf("%v", project)
	err := dao.InsertProject(project)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("%v", project2)
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
