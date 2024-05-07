package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	error = user.Prepare("register")
	
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
	nameOrUsername := strings.ToLower(r.URL.Query().Get("user"))

	db, error := database.Connect()

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, error := repository.GetUsers(nameOrUsername)

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusOK, users)
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)

	userId, error := strconv.ParseUint(params["userId"], 10, 64)

	if error !=nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()

	if error !=nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	user, error := repository.GetUserById(userId)

	if error !=nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusOK, user)

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, error := strconv.ParseUint(params["userId"], 10, 64)

	if error !=nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	body, error := io.ReadAll(r.Body)

	if error !=nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}


	var user models.User

	if error := json.Unmarshal(body, &user); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	
	if error := user.Prepare("update"); error != nil {
		response.Error(w, http.StatusBadRequest, error)
	return
	}

	db, error := database.Connect()

	if error !=nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	repository.UpdateUser(userId, user)

	response.JSON(w, http.StatusNoContent, nil)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, error := strconv.ParseUint(params["userId"], 10, 64)

	if error !=nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()

	if error !=nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	repository.DeleteUser(userId)

	response.JSON(w, http.StatusNoContent, nil)
}