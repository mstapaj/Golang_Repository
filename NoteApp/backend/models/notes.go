package models

import "gopkg.in/mgo.v2/bson"

type Note struct {
	ID      bson.ObjectId `bson:"_id" json:"_id"`
	Title   string        `bson:"title" json:"title"`
	Content string        `bson:"content" json:"content"`
}
