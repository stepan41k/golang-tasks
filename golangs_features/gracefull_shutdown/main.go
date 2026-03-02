package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr: ":8080",
		Handler: nil,
	}

	go func() {
		fmt.Println("server started on :8080")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and server error: %v", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("shutting down gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("forced shutdown: %v", err)
	}

	fmt.Println("sever exited gracefully")
}

// func main() {
//     srv := &http.Server{Addr: ":8080"}

// 	go func() {
// 		err := srv.ListenAndServe()
// 		if err != nil {
// 			if errors.Is(err, http.ErrServerClosed) {
// 				fmt.Printf("server closed: %s\n", err.Error())
// 				return
// 			}
			
// 		}
// 	}()

// 	ch := make(chan os.Signal, 1)
// 	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	
// 	_ = <- ch

// 	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
// 	defer cancel()

// 	fmt.Println("received interrup signal, shutting down")

// 	err := srv.Shutdown(ctx)
// 	if err != nil {
// 		fmt.Printf("error with shutting down server: %s\n", err.Error())
// 		return
// 	}

// 	fmt.Println("successfull shutdown")
// }

