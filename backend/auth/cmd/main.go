package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	psql "github.com/DexScen/SuSuSport/backend/auth/internal/repository/psql"
	"github.com/DexScen/SuSuSport/backend/auth/internal/service"
	"github.com/DexScen/SuSuSport/backend/auth/internal/transport/rest"
	"github.com/DexScen/SuSuSport/backend/auth/pkg/database"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	usersRepo := psql.NewUsers(db)
	usersService := service.NewUsers(usersRepo)
	handler := rest.NewUsers(usersService)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: handler.InitRouter(),
	}

	log.Println("Server started at:", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
