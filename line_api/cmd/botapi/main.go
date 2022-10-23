package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/tmkshy1908/LINE-bot/infrastructure"
	db "github.com/tmkshy1908/LINE-bot/infrastructure/db"
)

func main() {

	db, err := db.NewHandler()
	if err != nil {
		fmt.Println(err)
	}

	handler := infrastructure.NewServer(db)

	// Port Conf
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	srv := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: 180 * time.Second,
		ReadTimeout:  180 * time.Second,
		IdleTimeout:  300 * time.Second,
	}

	// tag := logger.NewTag("host", srv.Addr)
	// logger.Info("00", "Starting Serverer Listening", tag)

	if err := srv.ListenAndServe(); err != nil {
		// logger.Fatal(
		// logger.GetApplicationError(err).
		// AddMessage("Server Down..."),
		// )
	}
}
