package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/dto"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {

	userID, error := authentication.ExtractUserId(r)

	if error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	body, error := io.ReadAll(r.Body)

	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	var createPostDto dto.CreatePost

	if error := json.Unmarshal(body, &createPostDto); error !=nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()

	if error !=nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewPostsRepository(db)


	var post = models.Post{
		Title: createPostDto.Title,
		Content: createPostDto.Content,
		AuthorID: userID,
		CreatedAt: time.Now(),
	}

	if error := post.Prepare(); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}

	postID, error := repository.CreatePost(userID, post.Title, post.Content)

	if error !=nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	post.ID = postID

	response.JSON(w, http.StatusCreated, post)

}

func GetPostById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	postId, error := strconv.ParseUint(params["postId"], 10, 64)

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

	repository := repositories.NewPostsRepository(db)


	post, error := repository.GetPostByID(postId)

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	} 

	response.JSON(w, http.StatusOK, post)
}