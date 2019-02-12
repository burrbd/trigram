package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/burrbd/trigram"
	"github.com/burrbd/trigram/cmd/web/handlers"
)

const TrigramGetLimit = 1000

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	gen := trigram.NewLanguageGenerator(trigram.NewMapStore(TrigramGetLimit))

	mux := http.NewServeMux()
	mux.Handle("/generate", handlers.GenerateHandler(gen))
	mux.Handle("/learn", handlers.LearnHandler(gen))

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
		Handler:      mux,
	}
	go func() {
		log.Println("listen on port 8080")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	<-stop
	log.Println("shutting down...")
	server.Shutdown(context.Background())
	log.Println("server gracefully stopped")
}
