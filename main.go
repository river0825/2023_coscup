package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.io/river0825/2023_coscup/module/backpack/port/controller"
	"github.io/river0825/2023_coscup/module/backpack/port/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// setup graceful shutdown for gin server

	// load configuration if there is any

	// create repository
	repo := repository.NewInMemRepository("backpack")

	// create controller
	ginController := controller.NewGinController(repo)

	// setup gin
	router := gin.Default()
	v1 := router.Group("api/v1")
	backpack := v1.Group("backpack")
	backpack.POST("/putitem", ginController.PutItem)
	backpack.POST("/takeout", ginController.TakeOut)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
