package main

import (
	"log"
	"os"
	"time"

	"github.com/globalsign/mgo"
)

var db *mgo.Database
var prj ProjectDTO

const (
	PRJ_COLLECTION = "projects"
)

func init() {
	os.Setenv("MONGO_URL", "127.0.0.1:27017")
	os.Setenv("MONGO_DATABASE", "local")
}

func Connect() {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(os.Getenv("MONGO_DATABASE"))
}

func main() {
	//Insert a project
	prj := ProjectDTO{ID: 1, Code: "8001", Description: "HelloWorld project", StartDate: time.Date(2018, time.Month(12), 3, 0, 0, 0, 0, time.UTC)}
	err := db.C(PRJ_COLLECTION).Insert(&prj)

}
