package main

import (
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
)

const (
	host = ""
	port = "8080"
)

// Run app:
//		make run
// Run tests (with test coverage):
//		make test
// Send requests with:
//		curl -F 'file=@./data/matrix.csv' "localhost:8080/echo"
//		curl -F 'file=@./data/matrix.csv' "localhost:8080/invert"
//		curl -F 'file=@./data/matrix.csv' "localhost:8080/flatten"
//		curl -F 'file=@./data/matrix.csv' "localhost:8080/sum"
//		curl -F 'file=@./data/matrix.csv' "localhost:8080/multiply"

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Panic("error while building zapLogger", err)
	}
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Fatalf("error while sync logger %s", err.Error())
		}
	}()

	router := NewRouter(logger)
	router.InitRoutes()

	logger.Info("Server started", zap.String("port", port))
	err = http.ListenAndServe(net.JoinHostPort(host, port), router)
	if err != nil {
		logger.Error("ListenAndServe failed", zap.Error(err))
		return
	}

}
