package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/dto"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"fmt"
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

	if post.ID == 0 {
		response.Error(w, http.StatusNotFound, errors.New(fmt.Sprintf("the post with id %d does not exist", postId)))
		return 
	}

	response.JSON(w, http.StatusOK, post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {

	userID, error := authentication.ExtractUserId(r)

	if error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	db, error := database.Connect()

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}

	defer db.Close()

	repository := repositories.NewPostsRepository(db)

	posts, error := repository.GetPosts(userID)

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return  
	}


	response.JSON(w, http.StatusOK, posts)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, error := strconv.ParseUint(params["postId"], 10, 64)

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

	databasePost, error := repository.GetPostByID(postID)


	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}

	userID, error := authentication.ExtractUserId(r)

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}
	
	if databasePost.AuthorID != userID {
		response.Error(w, http.StatusForbidden, errors.New("cannot update post that you are not the author"))
		return 
	}

	body, error := io.ReadAll(r.Body)

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}

	var post models.Post

	if error := json.Unmarshal(body, &post); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return  
	}

	if error := post.Prepare(); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return  
	}

	if error := repository.UpdatePost(postID, post); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return  
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, error := strconv.ParseUint(params["postId"], 10, 64)

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

	databasePost, error := repository.GetPostByID(postID)

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}

	if databasePost.ID == 0 {
		response.Error(w, http.StatusNotFound, errors.New(fmt.Sprintf("the post with id %d does not exist", postID)))
		return 
	}

	userID, error := authentication.ExtractUserId(r)

	if databasePost.AuthorID != userID {
		response.Error(w, http.StatusForbidden, errors.New("cannot delete post that you are not the author"))
		return 
	}

	if error := repository.DeletePost(postID); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return 
	}

	response.JSON(w, http.StatusNoContent, nil)
}