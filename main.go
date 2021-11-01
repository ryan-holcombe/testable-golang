package main

//go:generate go run github.com/vektra/mockery/v2@v2.9.0 -r --inpackage --dir ./api --testonly --all --case=underscore --disable-version-string
//go:generate go run github.com/vektra/mockery/v2@v2.9.0 -r --inpackage --dir ./service --testonly --all --case=underscore --disable-version-string
//go:generate go run github.com/vektra/mockery/v2@v2.9.0 -r --inpackage --dir ./dao --testonly --all --case=underscore --disable-version-string
//go:generate go run github.com/vektra/mockery/v2@v2.9.0 -r --inpackage --dir ./client --testonly --all --case=underscore --disable-version-string

import (
	"context"
	"flag"
	"fmt"
	"github.com/ryan-holcombe/testable-golang/api"
	"github.com/ryan-holcombe/testable-golang/client"
	"github.com/ryan-holcombe/testable-golang/dao"
	"github.com/ryan-holcombe/testable-golang/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	port             = flag.Int("port", 8888, "http port to listen on")
	inventoryBaseURL = flag.String("inventoryBaseUrl", "http://inventory", "base URL of the inventory service")
	dbConn           = flag.String("dbConn", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=user sslmode=disable", "DB connection string")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	inventoryClient, err := client.NewInventoryClient(*inventoryBaseURL)
	if err != nil {
		log.Fatalf("unable to initialize inventory client")
	}
	sqlxDB := dao.NewPostgresDB(*dbConn)
	userDAO := dao.NewUserDAO(sqlxDB)
	userTicketsService := service.NewUserTicketsService(userDAO, inventoryClient)
	api.RegisterRoutes(mux, userTicketsService)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: mux,
	}

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	go func() {
		log.Printf("HTTP server listening on port %d", *port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	<-signalChan

	log.Println("HTTP server shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP server unable to shut down gracefully: %v", err)
	}

	log.Println("HTTP server shut down")
}
