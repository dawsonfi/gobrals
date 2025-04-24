package main

import (
	"github.com/dawsonfi/gobrals/internal/address"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Loading Address Database")
	addressService, err := address.NewService(address.NewJsonDatabaseLoader(
		"data/br_address_database.json",
	))
	if err != nil {
		slog.Error("Failed to open repository file", "error", err.Error())
		return
	}

	slog.Info("Creating Handlers")
	handler := address.NewHandler(addressService)

	slog.Info("Starting HTTP Server")
	sm := http.NewServeMux()
	sm.HandleFunc("/address/{zip_code}", handler.HandleGet)

	err = http.ListenAndServe("localhost:4000", sm)
	if err != nil {

		return
	}
}
