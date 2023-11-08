package main

import (
	"log"
	"os"

	"ipr-savelichev/internal/config"
	"ipr-savelichev/internal/jwt"
	"ipr-savelichev/internal/logger"
	"ipr-savelichev/internal/router"
	"ipr-savelichev/internal/server"
	"ipr-savelichev/internal/store/postgredb"
)

func main() {
	config, err := config.LoadConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	logger, err := logger.NewLogger(config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	db, err := postgredb.NewGorm(config, logger)
	if err != nil {
		log.Fatalln(err.Error())
	}

	store, err := postgredb.NewPostgre(db, logger)
	if err != nil {
		log.Fatalln(err.Error())
	}

	jwt, err := jwt.NewJWTController(config, logger)
	if err != nil {
		log.Fatalln(err.Error())
	}

	router := router.NewRouter(logger, store, jwt)
	server := server.NewServer(router, store, config)

	err = server.StartServer()
	if err != nil {
		log.Fatalln(err.Error())
	}

}
