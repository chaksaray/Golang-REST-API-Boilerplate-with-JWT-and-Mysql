package controllertests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"skeleton_project/app/controllers"
	"skeleton_project/app/startup"
	"skeleton_project/config"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSignInFuncSuccess(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}

	sample := struct {
		username        string
		password     string
		errorMessage string
	}{
		username: user.Username,
		password: "123456",
		errorMessage: "",
	}

	token, _ := controllers.SignIn(server.DB, sample.username, sample.password)
	assert.NotEqual(t, token, "")
}

func TestSignInFuncFailed(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}

	sample := struct {
		username        string
		password     string
		errorMessage string
	}{
		username: user.Username,
		password: "Wrong password",
		errorMessage: "crypto/bcrypt: hashedPassword is not the hash of the given password",
	}

	_, err = controllers.SignIn(server.DB, sample.username, sample.password)
	assert.Equal(t, err, errors.New(sample.errorMessage))
}

func TestLoginSuccess(t *testing.T) {

	refreshUserTable()

	_, err := seedOneUser()
	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}
	sample := struct {
		inputJSON    string
		statusCode   int
		errorMessage string
	}{
		inputJSON:    `{"username": "chaksaray008", "password": "123456"}`,
		statusCode:   http.StatusOK,
		errorMessage: "",
	}

	req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(sample.inputJSON))
	if err != nil {
		t.Errorf("this is the error: %v", err)
	}
	rr := httptest.NewRecorder()

	server := &startup.App{}
	server.Initialize(config.GetConfig())

	handler := http.HandlerFunc(server.HandleRequest(controllers.Login))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, sample.statusCode)
	if sample.statusCode == http.StatusOK {
		assert.NotEqual(t, rr.Body.String(), "")
	}
}

func TestLoginBadRequest(t *testing.T) {

	refreshUserTable()

	_, err := seedOneUser()
	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}
	sample := struct {
		inputJSON    string
		statusCode   int
		errorMessage string
	}{
		inputJSON:    `{"username": "badrequest", "password": "123456"}`,
		statusCode:   http.StatusBadRequest,
		errorMessage: "Invalid username or password",
	}

	req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(sample.inputJSON))
	if err != nil {
		t.Errorf("this is the error: %v", err)
	}
	rr := httptest.NewRecorder()

	server := &startup.App{}
	server.Initialize(config.GetConfig())

	handler := http.HandlerFunc(server.HandleRequest(controllers.Login))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, sample.statusCode)
	if sample.statusCode == http.StatusBadRequest && sample.errorMessage != "" {
		responseMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to json: %v", err)
			}
			assert.Equal(t, responseMap["message"], sample.errorMessage)
	}
}
