package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	Server struct {
		DB        *mongo.Database
		Router    *mux.Router
		AppConfig *AppConfig
		Context   context.Context
	}

	AppConfig struct {
		AppName string
		AppURL  string
		AppEnv  string
		AppPort string
	}

	DBConfig struct {
		DBDriver string
		DBHost   string
		DBPort   string
		DBUser   string
		DBPass   string
		DBName   string
	}
)

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.Context = context.Background()
}
