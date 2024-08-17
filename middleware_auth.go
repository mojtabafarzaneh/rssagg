package main

import (
	"fmt"
	"net/http"

	"github.com/mojtabafarzaneh/rssagg/auth"
	"github.com/mojtabafarzaneh/rssagg/internal/database"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiconf *apiConfig) MiddlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithErr(w, 404, fmt.Sprintf("please provide the right apikey: %s", err))
			return

		}
		user, err := apiconf.DB.GetUser(r.Context(), apikey)
		if err != nil {
			responseWithErr(w, 404, fmt.Sprintf("couldn't find the user: %s", err))
			return
		}

		handler(w, r, user)
	}
}
