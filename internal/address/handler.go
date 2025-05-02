package address

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Fetcher interface {
	Fetch(zipCode string) (*Address, bool)
}

type Handler struct {
	fetcher *Service
}

func NewHandler(fetcher *Service) *Handler {
	return &Handler{fetcher: fetcher}
}

func (handler *Handler) HandleGet(response http.ResponseWriter, request *http.Request) {
	zipCode := request.PathValue("zip_code")
	foundAddress, ok := handler.fetcher.Fetch(zipCode)
	if !ok {
		response.WriteHeader(http.StatusNotFound)
		_, err := fmt.Fprintf(response, "ZipCode: %s not found", zipCode)
		if err != nil {
			slog.Error("Failed to Write response", "error", err.Error())
		}
		return
	}

	jsonAddress, err := foundAddress.ToJson()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(response, "Failed to convert Address to json with error: %s", err.Error())
		if err != nil {
			slog.Error("Failed to Write response", "error", err.Error())
		}
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(response, jsonAddress)
	if err != nil {
		slog.Error("Failed to Write response", "error", err.Error())
		return
	}
}
