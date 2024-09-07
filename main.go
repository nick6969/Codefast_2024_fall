package main

import (
	"codefast_2024/app"
	"codefast_2024/route"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := app.NewApp()
	server := route.NewServer(app)
	server.Start()

	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("Unexpected error: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	server.Shutdown()
}
