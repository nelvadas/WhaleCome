package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//ProjectDTO structure
// contains information on a project item
type ProjectDTO struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Code        string        `bson:"code" json:"code"`
	Description string        `bson:"description" json:"description,omitempty"`
	StartDate   time.Time     `bson:"dateFrom" json:"from,omitempty"`
	EndDate     time.Time     `bson:"dateTo" json:"to,omitempty"`
}
