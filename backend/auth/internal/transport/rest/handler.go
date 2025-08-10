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
	LogIn(ctx context.Context, login, password string) (string, error)
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
		links.HandleFunc("", h.OptionsHandler).Methods(http.MethodOptions)
	}
	return r
}

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	var info domain.LoginInfo
	var roleInfo domain.RoleInfo

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Login error:", err)
		return
	}

	role, err := h.usersService.LogIn(context.TODO(), info.Login, info.Password)
	if err != nil {
		if errors.Is(err, e.ErrUserNotFound) {
			role = "unauthorized by user"
			log.Println("Login error2:", err)
		} else if errors.Is(err, e.ErrWrongPassword) {
			role = "unauthorized by password"
			log.Println("Login error3:", err)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Login error1:", err, e.ErrUserNotFound)
			return
		}
	}
	roleInfo.Role = role
	if jsonResp, err := json.Marshal(roleInfo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Login error:", err)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
}