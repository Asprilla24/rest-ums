package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Asprilla24/rest-ums/repository"
)

type User struct {
	Repository repository.User
}

//GetUser : is a function to get user by email or username
func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	var response Response
	user, err := u.Repository.GetUser("Asprilla24")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		response.Result = struct{}{}
	} else {
		w.WriteHeader(http.StatusOK)
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Result = user
	}

	json.NewEncoder(w).Encode(response)
}
