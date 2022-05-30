package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/sandjuarezg/neptuno-api/models"
)

var userSession models.Session
var idProyect string

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/proyectos", proyectos)
	http.HandleFunc("/tareas", tareas)

	fmt.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func signin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}
	defer fmt.Printf("Response from %s\n", r.URL.RequestURI())

	if err := r.ParseForm(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	t, err := template.ParseFiles("./public/index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = models.Signin(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func proyectos(w http.ResponseWriter, r *http.Request) {
	defer fmt.Printf("Response from %s\n", r.URL.RequestURI())

	// edit
	if strings.Contains(r.URL.String(), "/proyectos?editar=") {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		err := models.EditProyect(
			userSession,
			strings.Trim(r.URL.String(), "/proyectos?editar="),
			r,
		)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	} else

	// delete
	if strings.Contains(r.URL.String(), "/proyectos?eliminar=") {
		err := models.DeleteProyect(
			userSession,
			strings.Trim(r.URL.String(), "/proyectos?eliminar="),
		)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	} else

	// add
	if strings.Contains(r.URL.String(), "/proyectos?agregar") {
		err := models.AddProyect(userSession, r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	} else {
		// login
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		var err error

		userSession, err = models.Login(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	proyects, err := models.GetProyects(userSession)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	t, err := template.ParseFiles("./public/html/proyectos.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = t.Execute(w, proyects)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func tareas(w http.ResponseWriter, r *http.Request) {
	defer fmt.Printf("Response from %s\n", r.URL.RequestURI())

	// edit
	if strings.Contains(r.URL.String(), "/tareas?editar=") {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		err := models.EditTask(
			userSession,
			strings.Trim(r.URL.String(), "/tareas?editar="),
			r,
		)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	} else

	// delete
	if strings.Contains(r.URL.String(), "/tareas?eliminar=") {
		err := models.DeleteTask(
			userSession,
			strings.Trim(r.URL.String(), "/tareas?eliminar="),
		)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	} else {
		// add
		if strings.Contains(r.URL.String(), "/tareas?agregar") {
			err := models.AddTask(userSession, idProyect, r)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)

				return
			}
		}
	}

	if strings.Contains(r.URL.String(), "/tareas?id=") {
		idProyect = strings.Trim(r.URL.String(), "/tareas?id=")
	}

	tasks, err := models.GetTasks(userSession, idProyect)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	t, err := template.ParseFiles("./public/html/tareas.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = t.Execute(w, tasks)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
