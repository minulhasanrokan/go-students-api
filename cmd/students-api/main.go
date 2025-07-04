package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/minulhasanrokan/students-api/internal/config"
	"github.com/minulhasanrokan/students-api/internal/http/handlers/student"
)

func main() {

	cfg := config.MustLoad()

	//fmt.Println(cfg)

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	fmt.Println("start server")

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("faild to statrt server")
		}
	}()

	<-done

	slog.Info("Sutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("fails to shutdown sfferver")
	}

	slog.Info("Server shutdown successfully")

	fmt.Println("start server")
}
