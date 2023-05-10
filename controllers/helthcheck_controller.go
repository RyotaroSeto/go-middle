package controllers

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func HelthCheckHandler(w http.ResponseWriter, req *http.Request) {
	health := Health{
		Status: http.StatusOK,
		Result: "success",
	}

	json.NewEncoder(w).Encode(health)
}
