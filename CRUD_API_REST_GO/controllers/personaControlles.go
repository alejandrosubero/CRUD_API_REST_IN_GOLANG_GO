package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/alejandrosubero/crud-api/commons"
	"github.com/alejandrosubero/crud-api/models"
	"github.com/gorilla/mux"
)

func GetALL(writer http.ResponseWriter, request *http.Request) {

	personas := []models.Persona{}
	db := commons.GetConnection()
	defer db.Close()
	db.LogMode(true)
	db.Preload("Direcciones").Find(&personas)
	json, _ := json.Marshal(personas)
	commons.SendResponse(writer, http.StatusOK, json)
	db.LogMode(false)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	id := mux.Vars(request)["id"]
	db := commons.GetConnection()
	db.LogMode(true)
	defer db.Close()
	db.Find(&persona, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)
		commons.SendResponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
	db.LogMode(false)
}

func Save(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	db := commons.GetConnection()
	db.LogMode(true)
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&persona)
	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	// error = db.Create(&persona).Error
	error = db.Save(&persona).Error
	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)
	commons.SendResponse(writer, http.StatusCreated, json)
	db.LogMode(false)
}

func Delete(writer http.ResponseWriter, request *http.Request) {

	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()
	id := mux.Vars(request)["id"]

	db.Find(&persona, id)

	if persona.ID > 0 {
		db.Delete(persona)
		commons.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

func Base(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Golang!")
}

func FileDowload(writer http.ResponseWriter, request *http.Request) {
	nombre := mux.Vars(request)["nombre"]
	writer.Header().Set("Content-Disposition", "attachment; filename="+nombre+".txt")
	writer.Header().Set("Content-Type", request.Header.Get("Content-Type"))
	writer.Header().Set("Content-Length", request.Header.Get("Content-Length"))
	http.ServeFile(writer, request, filepath.Join("file", nombre+".txt"))
}
