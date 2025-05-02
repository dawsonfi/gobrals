package main

import (
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Creating Handlers")
	handler, err := InitializeAddressHandler("data/br_address_database.json")
	if err != nil {
		slog.Error("Failed to create handler", "error", err.Error())
		return
	}

	slog.Info("Starting HTTP Server")
	sm := http.NewServeMux()
	sm.HandleFunc("/address/{zip_code}", handler.HandleGet)

	err = http.ListenAndServe("localhost:4000", sm)
	if err != nil {
		return
	}
}
