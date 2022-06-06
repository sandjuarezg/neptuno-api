package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Session struct {
	Type         string `json:"type"`
	Token        string `json:"token"`
	RefreshToken bool   `json:"refreshToken"`
}

func Login(r *http.Request) (userSession Session, errMess string, err error) {
	if err = r.ParseForm(); err != nil {
		return
	}

	userJSON, err := json.Marshal(
		User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		})
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		"https://quiet-fortress-17520.herokuapp.com/api/v1/usuarios/login",
		bytes.NewBuffer(userJSON),
	)
	if err != nil {
		return
	}

	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var errMessage []ErrorMessage

	if response.StatusCode != http.StatusOK {
		err = json.Unmarshal(content, &errMessage)
		if err != nil {
			return
		}

		errMess = errMessage[0].Message

		return
	}

	err = json.Unmarshal(content, &userSession)
	if err != nil {
		return
	}

	return
}

func Signin(r *http.Request) (err error) {
	userJSON, err := json.Marshal(
		User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		})
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		"https://quiet-fortress-17520.herokuapp.com/api/v1/usuarios/registro",
		bytes.NewBuffer(userJSON),
	)
	if err != nil {
		return
	}

	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		return
	}

	return
}
