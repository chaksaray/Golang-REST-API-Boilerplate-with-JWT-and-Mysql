package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"skeleton_project/app/controllers"
	"skeleton_project/app/models"
	"skeleton_project/app/startup"
	"skeleton_project/config"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)


func TestCreateUserSuccess(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	sample := struct {
		inputJSON    string
		statusCode   int
		name     string
		username        string
		errorMessage string
	}{
		inputJSON: `{"name":"name 1", "username": "username 1", "password": "password"}`,
		statusCode: http.StatusCreated,
		name: "name 1",
		username: "username 1",
		errorMessage: "",

	}

	req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(sample.inputJSON))
	if err != nil {
		t.Errorf("this is the error: %v", err)
	}
	rr := httptest.NewRecorder()
	server := &startup.App{}
	server.Initialize(config.GetConfig())

	handler := http.HandlerFunc(server.HandleRequest(controllers.CreateUser))
	handler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr.Code, sample.statusCode)
	if sample.statusCode == http.StatusCreated {
		assert.Equal(t, responseMap["name"], sample.name)
		assert.Equal(t, responseMap["username"], sample.username)
	}
}

func TestCreateUserRequiredName(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	sample := struct {
		inputJSON    string
		statusCode   int
		name     string
		username        string
		errorMessage string
	}{
		inputJSON: `{"name":"", "username": "username 1", "password": "password"}`,
		statusCode: http.StatusBadRequest,
		name: "name 1",
		username: "username 1",
		errorMessage: "name is required",

	}

	req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(sample.inputJSON))
	if err != nil {
		t.Errorf("this is the error: %v", err)
	}
	rr := httptest.NewRecorder()
	server := &startup.App{}
	server.Initialize(config.GetConfig())

	handler := http.HandlerFunc(server.HandleRequest(controllers.CreateUser))
	handler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr.Code, sample.statusCode)
	if sample.statusCode == http.StatusBadRequest && sample.errorMessage != "" {
		assert.Equal(t, responseMap["message"], sample.errorMessage)
	}
}

func TestGetUsers(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()

	server := &startup.App{}
	server.Initialize(config.GetConfig())

	handler := http.HandlerFunc(server.HandleRequest(controllers.GetAllUsers))
	handler.ServeHTTP(rr, req)

	var users []models.User
	err = json.Unmarshal([]byte(rr.Body.String()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(users), 2)
}
