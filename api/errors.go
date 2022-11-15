package api

import (
	"net/http"
	"net/rpc"
)

type envelope map[string]interface{}

func ErrorResponse(w http.ResponseWriter, r *rpc.Response, status int, message interface{}) {
	env := envelope{"error": message}
}
