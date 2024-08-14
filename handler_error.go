package main

import "net/http"

func HandleErrResponse(w http.ResponseWriter, r *http.Request) {
	responseWithErr(w, 400, "something went wrong")
}
