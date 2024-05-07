package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/dto"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"errors"
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

	tokenUserId, error := authentication.ExtractUserId(r)

	if error !=nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	if tokenUserId != userId {
		response.Error(w, http.StatusForbidden, errors.New("you can only edit your own user"))
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

	tokenUserId, error := authentication.ExtractUserId(r)

	if error !=nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	if tokenUserId != userId {
		response.Error(w, http.StatusForbidden, errors.New("you can only delete your own user"))
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

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, error := authentication.ExtractUserId(r)

	if error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["userId"], 10, 64)

	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if (followerID == userID) {
		response.Error(w, http.StatusForbidden, errors.New("it is not possible to follow yourself"))
		return
	}

	db, error := database.Connect()

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	if error := repository.Follow(userID, followerID); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	
	response.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, error := authentication.ExtractUserId(r)

	if error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["userId"], 10, 64)

	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if (followerID == userID) {
		response.Error(w, http.StatusForbidden, errors.New("it is not possible to unfollow yourself"))
		return
	}

	db, error := database.Connect()

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	if error := repository.Unfollow(userID, followerID); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	
	response.JSON(w, http.StatusNoContent, nil)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)


	userID, error := strconv.ParseUint(params["userId"], 10, 64)

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

	followers, error := repository.GetFollowers(userID)


	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)


	userID, error := strconv.ParseUint(params["userId"], 10, 64)

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

	followers, error := repository.GetFollowing(userID)


	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	userID, error := authentication.ExtractUserId(r)

	if error != nil {
		response.Error(w, http.StatusForbidden, error)
		return
	}

	body, error := io.ReadAll(r.Body)

	if error != nil {
		response.Error(w, http.StatusForbidden, error)
		return
	}

	var updatePasswordDto dto.UpdatePassword

	if error := json.Unmarshal(body, &updatePasswordDto); error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return 
	}

	db, error := database.Connect()

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	databasePassword, error := repository.GetPassword(userID)

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}

	if error := security.VerifyPassword(databasePassword, updatePasswordDto.CurrentPassword); error != nil {
		response.Error(w, http.StatusInternalServerError, errors.New("the informed password does not match with the current one"))
		return 
	}

	hashPassword, error := security.Hash(updatePasswordDto.NewPassowrd)

	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return 
	}


	if error := repository.UpdatePassword(userID, string(hashPassword)); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}
	
	response.JSON(w, http.StatusNoContent, nil)

}