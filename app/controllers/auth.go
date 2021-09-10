package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"skeleton_project/app/auth"
	"skeleton_project/app/models"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user.Prepare()

	err = user.Validate("login")
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	token, err := SignIn(db, user.Username, user.Password)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid username or password")
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}

func SignIn(db *gorm.DB, username string, password string) (string, error) {

	var err error

	user := models.User{}

	err = db.Model(models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}


