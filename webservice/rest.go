package webservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Note      string    `json:"note,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

var users []User = []User{{
	ID:        1,
	FirstName: "John",
	LastName:  "Example",
	Note:      "",
	CreatedAt: time.Now(),
}}

func RunREST() {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello Follow The Pattern!")
		})
		r.Route("/users", func(r chi.Router) {
			r.Get("/", GetUser)
			r.Post("/", CreateUser)
		})
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	fmt.Println("Server Started!")
	log.Fatal(server.ListenAndServe())
}

func Parse(r *http.Request, T interface{}) error {
	err := json.NewDecoder(r.Body).Decode(T)
	if err != nil {
		return err
	}

	return nil
}

func Write(w http.ResponseWriter, responseStatus int, T interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(responseStatus)
	jsonObj, err := json.Marshal(T)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonObj)
}

func Success(w http.ResponseWriter, T interface{}) {
	Write(w, http.StatusOK, T)
}

func BadRequest(w http.ResponseWriter, T interface{}) {
	Write(w, http.StatusBadRequest, T)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		BadRequest(w, "failed to decode request Body")
		return
	}
	user.CreatedAt = time.Now()
	users = append(users, user)

	Success(w, "User created!")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	Success(w, users)
}
