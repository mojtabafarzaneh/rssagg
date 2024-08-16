package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mojtabafarzaneh/rssagg/internal/database"
)

func (apiconf *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type Params struct {
		Name string `json:"name"`
	}
	param := Params{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&param)
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("err parsing json: %s", err))
		return
	}

	user, err := apiconf.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      param.Name,
	})

	if err != nil {
		responseWithErr(w, 400, fmt.Sprint("couldn't create user: ", err))
		return
	}

	responseWithJSON(w, 200, DatabaseUserToUser(user))
}
