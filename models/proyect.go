package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Proyect struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Nombre    string `json:"nombre"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetProyects(userSession Session) (proyects []Proyect, err error) {
	userJSON, err := json.Marshal(userSession)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"GET",
		"https://quiet-fortress-17520.herokuapp.com/api/v1/proyectos",
		bytes.NewBuffer(userJSON),
	)
	if err != nil {
		return
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", userSession.Type+" "+userSession.Token)

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

	if response.StatusCode != http.StatusOK {
		return
	}

	err = json.Unmarshal(content, &proyects)
	if err != nil {
		return
	}

	return
}

func EditProyect(userSession Session, id string, r *http.Request) (err error) {
	proyectJSON, err := json.Marshal(
		Proyect{
			Nombre: r.FormValue("nombreProyect"),
		})
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"PATCH",
		"https://quiet-fortress-17520.herokuapp.com/api/v1/proyectos/"+id,
		bytes.NewBuffer(proyectJSON),
	)
	if err != nil {
		return
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", userSession.Type+" "+userSession.Token)

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

func DeleteProyect(userSession Session, id string) (err error) {
	request, err := http.NewRequest(
		"DELETE",
		"https://quiet-fortress-17520.herokuapp.com/api/v1/proyectos/"+id,
		nil,
	)
	if err != nil {
		return
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", userSession.Type+" "+userSession.Token)

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

func AddProyect(userSession Session, r *http.Request) (err error) {
	proyectJSON, err := json.Marshal(
		Proyect{
			Nombre: r.FormValue("nombreProyect"),
		})
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		"https://quiet-fortress-17520.herokuapp.com/api/v1/proyectos",
		bytes.NewBuffer(proyectJSON),
	)
	if err != nil {
		return
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", userSession.Type+" "+userSession.Token)

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
