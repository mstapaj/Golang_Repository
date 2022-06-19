package models

type QuickNote struct {
	Content string `bson:"content" json:"content"`
}
