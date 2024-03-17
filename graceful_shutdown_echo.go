package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	echoInstance := echo.New()
	echoInstance.Logger.SetLevel(log.INFO)
	echoInstance.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
	go func() {
		if err := echoInstance.Start(":1323"); err != nil && err != http.ErrServerClosed { // Start server
			echoInstance.Logger.Fatal("shutting down the server")
		}
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := echoInstance.Shutdown(ctx); err != nil {
		echoInstance.Logger.Fatal(err)
	}
}
