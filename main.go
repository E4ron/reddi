package main

import (
	"log"
	"reddi/pkg/handler"
	"reddi/pkg/repository"
	"reddi/pkg/service"
)

func main() {
	serverInstance := new(Server)

	postgresConfig := repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "123",
		DBName:   "reddit",
		SSLMode:  "disable",
	}

	database, err := repository.NewPostgresDB(postgresConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	repos := repository.NewRepository(database)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	runServer(serverInstance, handlers)
}

func runServer(serverInstance *Server, handlerLayer *handler.Handler) {
	port := "8080"
	router := handlerLayer.GetRouter()

	if err := serverInstance.Run(port, router); err != nil {
		log.Fatal(err.Error())
	}
	log.Print("server started successfully")
}
