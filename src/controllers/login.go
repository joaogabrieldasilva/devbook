package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)


func Login(w http.ResponseWriter, r *http.Request) {

	body, error := io.ReadAll(r.Body)

	if error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User

	if error := json.Unmarshal(body, &user); error != nil {
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

	databaseUser, error := repository.GetUserByEmail(user.Email)

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}


	if error := security.VerifyPassword(databaseUser.Password, user.Password);  error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return 
	}

	token, error := authentication.CreateToken(databaseUser.ID)

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}

	w.Write([]byte(token))

}
