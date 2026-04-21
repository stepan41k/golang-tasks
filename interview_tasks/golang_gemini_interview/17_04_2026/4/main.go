package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	server := http.Server{Addr: ":8081"}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("failed to listen and server http server")
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		runBackGroundTask(ctx)
	}()

	<-ctx.Done()
	log.Println("Shutdown signal received")
	
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("failed to graceful shutdown server")
	}
	
	wg.Wait()
	log.Println("Exit successfully")
}

func runBackGroundTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Worker: stopping...")
			time.Sleep(2 * time.Second)
			log.Println("Worker: stopped")
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}
