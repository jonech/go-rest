package controllers


import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
//	"io/ioutil"
//	"io"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/jonech/travel-bug-backend/models"
)


type (
	UserController struct{
		db *sql.DB
	}
)

func NewUserController(db *sql.DB) *UserController {
	return &UserController{db}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var u models.User

	// decode received json
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	// prepare database
	stmt, err := uc.db.Prepare("INSERT INTO users(name, age) VALUES ($1,$2)")
	if err != nil {
		panic(err)
	}

	// execute 
	_, err = stmt.Exec(u.Name, u.Age)
	if err != nil {
		panic(err)
	}

	// return something to client
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

func (uc UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {

	u := models.User {
		Id: "23",
		Name: "John",
		Age: 43,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

func (uc UserController) FindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars[ "userId" ]

	query := fmt.Sprintf("SELECT * FROM users WHERE id=%s", userId)
	rows, err := uc.db.Query(query)

	var user models.User
	for rows.Next() {
		
		var id string
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}

		user.Id = id
		user.Name = name
		user.Age = age
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}

}
