package main

import (
	c "backend/config"
	d "backend/dao"
	"backend/models"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

var config = c.Config{}
var dao = d.NotesDAO{}
var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, notes)
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	note, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}
	respondWithJson(w, http.StatusOK, note)
}

func AddNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	note.ID = bson.NewObjectId()
	if err := dao.Insert(note); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, note)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var note models.Note
	note.ID = bson.ObjectIdHex(params["id"])
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(note); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	note, _ := dao.FindById(params["id"])
	if err := dao.Delete(note); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func GetQuickNote(w http.ResponseWriter, r *http.Request) {
	val, err := rdb.Get(ctx, "quickNote").Result()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, val)
}

func AddQuickNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var quickNote models.QuickNote
	if err := json.NewDecoder(r.Body).Decode(&quickNote); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := rdb.Set(ctx, "quickNote", quickNote.Content, 0).Err()
	if err != nil {
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func UpdateQuickNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var quickNote models.QuickNote
	if err := json.NewDecoder(r.Body).Decode(&quickNote); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := rdb.Set(ctx, "quickNote", quickNote.Content, 0).Err()
	if err != nil {
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteQuickNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	err := rdb.Del(ctx, "quickNote").Err()
	if err != nil {
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/notes/", GetAllNotes).Methods("GET")
	r.HandleFunc("/notes/addNote", AddNote).Methods("POST")
	r.HandleFunc("/notes/editNote/{id}", UpdateNote).Methods("PUT")
	r.HandleFunc("/notes/deleteNote/{id}", DeleteNote).Methods("DELETE")
	r.HandleFunc("/notes/{id}", GetNoteByID).Methods("GET")
	r.HandleFunc("/quickNote", GetQuickNote).Methods("GET")
	r.HandleFunc("/quickNote/addQuickNote", AddQuickNote).Methods("POST")
	r.HandleFunc("/quickNote/editQuickNote", UpdateQuickNote).Methods("PUT")
	r.HandleFunc("/quickNote/deleteQuickNote", DeleteQuickNote).Methods("DELETE")
	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
	}).Handler(r)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatal(err)
	}
}
