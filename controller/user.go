package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Asprilla24/rest-ums/repository"
)

type User struct {
	Repository repository.User
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	user := u.Repository.GetUser("alde", "password")

	json.NewEncoder(w).Encode(user)
}
