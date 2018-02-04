package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type RepositoryInfo struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	Name          string
	Owner         string
	FullName      string
	Description   string
	Stars         int
	Forks         int
	UpdatedAt     time.Time
	LastUpdatedBy string
	ReadMe        string
	Tags          []string
	Categories    []string
	User          User
	Processed     bool
	ProcessedAt   time.Time
}
