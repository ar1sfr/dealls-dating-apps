package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"dealls-dating-apps/handlers"
	"dealls-dating-apps/utils"
)

func main() {

	utils.ConnectDB()

	handlers.Initialize(utils.Client)

	router := mux.NewRouter()
	router.HandleFunc("/", StatusHandler).Methods("GET")
	router.HandleFunc("/signup", handlers.SignUpHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	srv := &http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Server is running on port 9000")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not listen on port 9000: %v\n", err)
		}
	}()

	<-quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	if err := utils.Client.Disconnect(ctx); err != nil {
		log.Fatalf("could not disconnect from MongoDB: %v", err)
	}

	log.Println("server exiting")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "server is running"})
}
