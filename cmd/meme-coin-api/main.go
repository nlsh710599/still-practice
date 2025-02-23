package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nlsh710599/still-practice/internal/config"
	"github.com/nlsh710599/still-practice/internal/database"
	"github.com/nlsh710599/still-practice/internal/route"

	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := config.SetVar()
	if err != nil {
		fmt.Printf("error setting configuration variables: %v", err)
		return
	}

	r := gin.New()

	pg, err := database.NewPostgres(config.PG_DSN)
	if err != nil {
		fmt.Printf("error creating postgres client: %v", err)
	}

	err = route.Setup(r, pg)
	if err != nil {
		fmt.Printf("error setting up routes: %v", err)
		return
	}

	srv := &http.Server{
		Addr:    config.ServiceAddr,
		Handler: r,
	}

	nCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server failed to start: %v", err)
		}
	}()

	<-nCtx.Done()
	stop()

	fmt.Println("Shutting down server...")
	tCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(tCtx); err != nil {
		fmt.Printf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exiting")
}
