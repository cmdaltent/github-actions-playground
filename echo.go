package main

import (
	"io"
	"net/http"
)

func handler(response http.ResponseWriter, request *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
	}(request.Body)

	if request.Method != http.MethodPost {
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = response.Write(body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}
