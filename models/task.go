package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Task struct {
	ID          int    `json:"id"`
	ProjectoID  int    `json:"projecto_id"`
	Descripcion string `json:"descripcion"`
	Completada  bool   `json:"completada"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func GetTasks(userSession Session, id string) (tasks []Task, err error) {
	request, err := http.NewRequest(
		"GET",
		"https://quiet-fortress-17520.herokuapp.com/api/v1/proyectos/"+id+"/tareas",
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

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		return
	}

	err = json.Unmarshal(content, &tasks)
	if err != nil {
		return
	}

	return
}

func EditTask(userSession Session, id string, r *http.Request) (err error) {
	taskJSON, err := json.Marshal(
		Task{
			Descripcion: r.FormValue("nombreTarea"),
		})
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"PATCH",
		"https://quiet-fortress-17520.herokuapp.com/api/v1/tareas/"+id,
		bytes.NewBuffer(taskJSON),
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

func DeleteTask(userSession Session, id string) (err error) {
	request, err := http.NewRequest(
		"DELETE",
		"https://quiet-fortress-17520.herokuapp.com/api/v1/tareas/"+id,
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

func AddTask(userSession Session, id string, r *http.Request) (err error) {
	taskJSON, err := json.Marshal(
		Task{
			Descripcion: r.FormValue("nombreTarea"),
		})
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		"https://quiet-fortress-17520.herokuapp.com/api/v1/proyectos/"+id+"/tareas",
		bytes.NewBuffer(taskJSON),
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
