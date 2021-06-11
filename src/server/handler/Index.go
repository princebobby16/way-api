package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"way/src/server/index"
)

func Index(w http.ResponseWriter, r *http.Request){
	// get environment
	environment, valid := os.LookupEnv("ENVIRONMENT")
	if !valid {
		environment = "Development"
	}
	w.WriteHeader(http.StatusOK)

	indexResponse:= index.Index{
		Alive:       true,
		Author:      "Benjy",
		Maintainers: []string{},
		Email:       "aknanayaw77@gmail.com",
		System:      "WAY",
		Version:     "0",
		Environment: environment,
	}

	_ = json.NewEncoder(w).Encode(indexResponse)
	return
}
