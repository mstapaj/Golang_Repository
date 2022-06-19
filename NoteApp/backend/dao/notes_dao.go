package dao

import (
	n "backend/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type NotesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "notes"
)

func (m *NotesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *NotesDAO) FindAll() ([]n.Note, error) {
	var notes []n.Note
	err := db.C(COLLECTION).Find(bson.M{}).All(&notes)
	return notes, err
}

func (m *NotesDAO) FindById(id string) (n.Note, error) {
	var note n.Note
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&note)
	return note, err
}

func (m *NotesDAO) Insert(note n.Note) error {
	err := db.C(COLLECTION).Insert(&note)
	return err
}

func (m *NotesDAO) Delete(note n.Note) error {
	err := db.C(COLLECTION).Remove(&note)
	return err
}

func (m *NotesDAO) Update(note n.Note) error {
	err := db.C(COLLECTION).UpdateId(note.ID, &note)
	return err
}
