package controllers

import (
	"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateHandphone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var handphone entities.Handphone
	json.NewDecoder(r.Body).Decode(&handphone)
	database.Instance.Create(&handphone)
	json.NewEncoder(w).Encode(handphone)
}

func GetHandphoneById(w http.ResponseWriter, r *http.Request) {
	handphoneId := mux.Vars(r)["id"]
	if checkIfHandphoneExists(handphoneId) == false {
		json.NewEncoder(w).Encode("Handphone Not Found!")
		return
	}
	var handphone entities.Handphone
	database.Instance.First(&handphone, handphoneId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(handphone)
}

func GetHandphone(w http.ResponseWriter, r *http.Request) {
	var handphones []entities.Handphone
	database.Instance.Find(&handphones)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handphones)
}

func UpdateHandphone(w http.ResponseWriter, r *http.Request) {
	handphoneId := mux.Vars(r)["id"]
	if checkIfHandphoneExists(handphoneId) == false {
		json.NewEncoder(w).Encode("Handphone Not Found!")
		return
	}
	var handphone entities.Handphone
	database.Instance.First(&handphone, handphoneId)
	json.NewDecoder(r.Body).Decode(&handphone)
	database.Instance.Save(&handphone)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(handphone)
}

func DeleteHandphone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	handphoneId := mux.Vars(r)["id"]
	if checkIfHandphoneExists(handphoneId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Handphone Not Found!")
		return
	}
	var handphone entities.Handphone
	database.Instance.Delete(&handphone, handphoneId)
	json.NewEncoder(w).Encode("Handphone Deleted Successfully!")
}

func checkIfHandphoneExists(handphoneId string) bool {
	var handphone entities.Handphone
	database.Instance.First(&handphone, handphoneId)
	if handphone.ID == 0 {
		return false
	}
	return true
}
