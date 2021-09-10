package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

func Home(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, map[string]string{"message": "Welcome To This Awesome API."})
}