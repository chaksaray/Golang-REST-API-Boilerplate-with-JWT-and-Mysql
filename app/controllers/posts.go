package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"skeleton_project/app/auth"
	"skeleton_project/app/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func CreatePost(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	post := models.Post{}
	json.Unmarshal(body, &post)
	post.Prepare()

	err := post.Validate()
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		RespondError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	if userId != post.AuthorID {
		RespondError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	createdPost, err := post.CreatePost(db)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, createdPost)
}

func GetAllPosts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	_, err := auth.ExtractTokenID(r)
	if err != nil {
		RespondError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	post := models.Post{}
	posts, err := post.FindAllPosts(db)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, posts)
}

func GetPostById(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	_, err := auth.ExtractTokenID(r)
	if err != nil {
		RespondError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	// string to int
	pId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	post := models.Post{}

	existedPost, err := post.FindPostByID(db, uint32(pId))
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, existedPost)
}

func UpdatePost(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// string to int
	pId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// check if the post exist
	post := models.Post{}
	err = db.Model(models.Post{}).Where("id = ?", pId).Take(&post).Error
	if err != nil {
		RespondError(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	// check if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		RespondError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	// if a user attempt to update a post not belonging to him
	if uid != post.AuthorID {
		RespondError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	// read the data posted
	body, _ := ioutil.ReadAll(r.Body)

	// processing the request data
	updatedPost := models.Post{}
	json.Unmarshal(body, &updatedPost)

	// check if the request user id is equal to the one gotten from token
	if uid != updatedPost.AuthorID {
		RespondError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	updatedPost.Prepare()
	err = updatedPost.Validate()
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedPost.ID = post.ID

	newPost, err := updatedPost.UpdateAPost(db)

	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, newPost)
}

func DeletePost(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// string to int
	pId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// check if the post exist
	post := models.Post{}
	err = db.Model(models.Post{}).Where("id = ?", pId).Take(&post).Error
	if err != nil {
		RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		RespondError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	// Is the authenticated user, the owner of this post?
	if uid != post.AuthorID {
		RespondError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	_, err = post.DeleteAPost(db, uint32(pId), uid)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"message": "Delete."})
}
