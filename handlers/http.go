package handlers

import (
	"go-practice/domain"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	service domain.Service
}

func NewHandler(service domain.Service) Handler {
	return &handler{service}
}

func (h *handler) Get(rw http.ResponseWriter, r *http.Request) {
	name := getMovieName(r)
	metadata, err := h.service.Find(name)
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = domain.ToJSON(metadata, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (h *handler) Post(rw http.ResponseWriter, r *http.Request) {

}

func getMovieName(r *http.Request) string {
	vars := mux.Vars(r)

	name := vars["name"]

	return name
}
