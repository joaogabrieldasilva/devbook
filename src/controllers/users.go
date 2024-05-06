package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {
	
	body, error := io.ReadAll(r.Body)

	if error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User 

	if error = json.Unmarshal(body, &user); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	error = user.Prepare()
	
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	user.ID, error = repository.Create(user)
	
	if error !=nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusCreated, user)

}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}