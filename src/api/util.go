package api

import (
	"os"
	"log"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/tifmoe/go-fishbowl/src/errors"
)

// GetEnv is utility function to get default environment variable if not defined
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetIntEnv is utility function to get default int environment variable if not defined
func GetIntEnv(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		ret, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		return ret
	}
	return fallback
}

func serveResponse(w http.ResponseWriter, res *apiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)

	err := json.NewEncoder(w).Encode(res.Data)
	if err != nil {
		log.Fatalf("Error serving response %+v: %+v", res.Data, err)
	}
}

func buildResponse(game Game, er *errors.ErrorInternal, message string) *apiResponse {
	var res *Response
	status := 200

	if er.IsEmpty() {
		res = buildResult(game, message)
	} else {
		res = buildError(er.Error)
		status = er.Status
	}

	return &apiResponse{
		Status: status,
		Data: res,
	}
}

func buildResult(g Game, m string) *Response {
	return &Response{
		Result: []Game{g},
		Success: true,
		Error: []errors.Error{},
		Message: m,
	}
}

func buildError(e *errors.Error) *Response {
	return &Response{
		Result: []Game{},
		Success: false,
		Error: []errors.Error{*e},
		Message: "",
	}
}
