package rest

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/DexScen/SuSuSport/backend/auth/internal/domain"
	e "github.com/DexScen/SuSuSport/backend/auth/internal/errors"
	"github.com/gorilla/mux"
)

type Users interface {
	LogIn(ctx context.Context, login, password string) (*domain.User, error)
}

type Handler struct {
	usersService Users
}

func NewUsers(users Users) *Handler {
	return &Handler{
		usersService: users,
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

	links := r.PathPrefix("/users").Subrouter()
	{
		links.HandleFunc("/login", h.LogIn).Methods(http.MethodPost)
		links.HandleFunc("/login", h.OptionsHandler).Methods(http.MethodOptions)
	}
	return r
}

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	var info domain.LoginInfo
	result := &domain.User{}
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Login error:", err)
		return
	}
	var err error
	result, err = h.usersService.LogIn(context.TODO(), info.Login, info.Password)
	if err != nil {
		if errors.Is(err, e.ErrUserNotFound) {
			result = &domain.User{Role: "unauthorized by user"}
			log.Println("Login error2:", err)
		} else if errors.Is(err, e.ErrWrongPassword) {
			result = &domain.User{Role: "unauthorized by password"}
			log.Println("Login error3:", err)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Login error1:", err, e.ErrUserNotFound)
			return
		}
	}

	if jsonResp, err := json.Marshal(*result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Login error:", err)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
}
