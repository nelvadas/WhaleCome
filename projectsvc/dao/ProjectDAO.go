package dao

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
	. "github.com/nelvadas/WhaleCome/projectsvc/models"
	"gopkg.in/mgo.v2/bson"
)

//ProjectDAO interface
type ProjectDAO struct {
}

var db *mgo.Database
var prj ProjectDTO

//PRJCOLLECTION Project collection name
const PRJCOLLECTION = "projects"

func init() {
	os.Setenv("MONGO_URL", "127.0.0.1:27017")
	os.Setenv("MONGO_DATABASE", "local")
	log.Printf("MONGO_URL=%v", os.Getenv("MONGO_URL"))
	log.Printf("MONGO_DATABASE=%v", os.Getenv("MONGO_DATABASE"))
	Connect()
}

//Connect to Mongo DB
func Connect() {
	log.Printf("Connecting...")
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to DB %v", os.Getenv("MONGO_DATABASE"))
	db = session.DB(os.Getenv("MONGO_DATABASE"))
}

//InsertProject creates  new project in the database
func (dao *ProjectDAO) InsertProject(prjdto ProjectDTO) error {
	err := db.C(PRJCOLLECTION).Insert(&prjdto)
	return err

}

//FindProjectByID find project by id
func (dao *ProjectDAO) FindProjectByID(id string) (ProjectDTO, error) {
	var project ProjectDTO
	err := db.C(PRJCOLLECTION).FindId(bson.ObjectIdHex(id)).One(&project)
	return project, err

}

//FindAll find project by id
func (dao *ProjectDAO) FindAll() ([]ProjectDTO, error) {
	var projects []ProjectDTO
	err := db.C(PRJCOLLECTION).Find(bson.M{}).All(&projects)
	return projects, err

}
