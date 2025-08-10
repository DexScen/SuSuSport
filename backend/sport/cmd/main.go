package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	psql "github.com/DexScen/SuSuSport/backend/sport/internal/repository/psql"
	"github.com/DexScen/SuSuSport/backend/sport/internal/service"
	"github.com/DexScen/SuSuSport/backend/sport/internal/transport/rest"
	"github.com/DexScen/SuSuSport/backend/sport/pkg/database"
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

	sportRepo := psql.NewSport(db)
	sportService := service.NewSport(sportRepo)
	handler := rest.NewSport(sportService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	log.Println("Server started at:", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
