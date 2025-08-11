package rest

import (
	"context"
	"encoding/json"

	//"errors"
	"log"
	"net/http"

	//e "github.com/DexScen/SuSuSport/backend/sport/internal/errors"
	"github.com/DexScen/SuSuSport/backend/sport/internal/domain"
	"github.com/gorilla/mux"
)

type Sport interface {
	GetSections(ctx context.Context) (*[]string, error)
	GetSectionInfoByName(ctx context.Context, name string) (*domain.Section, error)
}

type Handler struct {
	sportService Sport
}

func NewSport(sport Sport) *Handler {
	return &Handler{
		sportService: sport,
	}
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (h *Handler) OptionsHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(loggingMiddleware)

	links := r.PathPrefix("/sport").Subrouter()
	{
		links.HandleFunc("/sections", h.GetSections).Methods(http.MethodGet)
		links.HandleFunc("/sections/{name}", h.GetSectionInfoByName).Methods(http.MethodGet)
		links.HandleFunc("/sections", h.OptionsHandler).Methods(http.MethodOptions)
		links.HandleFunc("/sections/{name}", h.OptionsHandler).Methods(http.MethodOptions)
	}
	return r
}

func (h *Handler) GetSections(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	names, err := h.sportService.GetSections(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("GetSections error:", err)
	}

	if jsonResp, err := json.Marshal(*names); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("GetSections error:", err)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
}

func (h *Handler) GetSectionInfoByName(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	vars := mux.Vars(r)
	name := vars["name"]
	log.Println(name)
	info, err := h.sportService.GetSectionInfoByName(context.TODO(), name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("GetSectionInfoByName error:", err)
		return
	}

	if jsonResp, err := json.Marshal(*info); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("GetSectionInfoByName error:", err)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
}
