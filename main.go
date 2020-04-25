package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Libro struct {
	ID          string `json:"id,omitempty"`
	Nombre      string `json:"nombre,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
	Autor       string `json:"autor,omitempty"`
	Editorial   string `json:"editorial,omitempty"`
}

var libros []Libro

func GetLibros(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(libros)
}

func GetLibro(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range libros {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Libro{})
}

func CreateLibro(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var libroo Libro
	_ = json.NewDecoder(req.Body).Decode(&libroo)
	libroo.ID = params["id"]
	libros = append(libros, libroo)
	json.NewEncoder(w).Encode(libros)

}
func DeleteLibro(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range libros {
		if item.ID == params["id"] {
			libros = append(libros[:index], libros[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(libros)

}
func UpdateLibro(w http.ResponseWriter, req *http.Request) {

}
func main() {
	router := mux.NewRouter()

	libros = append(libros, Libro{ID: "1", Nombre: "Pina", Descripcion: "novela", Autor: "la paty", Editorial: "tepache"})
	libros = append(libros, Libro{ID: "2", Nombre: "Paty", Descripcion: "novela", Autor: "la paty", Editorial: "tepache"})
	//endpoints
	router.HandleFunc("/libro", GetLibros).Methods("GET")
	router.HandleFunc("/libros/{id}", GetLibro).Methods("GET")
	router.HandleFunc("/libros/{id}", CreateLibro).Methods("POST")
	router.HandleFunc("/libros/{id}", DeleteLibro).Methods("DELETE")
	router.HandleFunc("/libros/{id}", UpdateLibro).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
