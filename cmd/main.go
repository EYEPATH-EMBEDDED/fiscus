package main

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"os"
)

var (
	log    = logging.MustGetLogger("usage")
	format = logging.MustStringFormatter(
		`%{color}[%{level:.4s}] %{time:2006/01/02 - 15:04:05}%{color:reset} ▶ %{message}`,
	)
)

func initLogger() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatted := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(formatted)
}

func main() {
	initLogger()

	router := gin.Default()
	RegisterRoutes(router) // api 등록

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
