package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/danilluk1/test-task-6/config"
	db "github.com/danilluk1/test-task-6/db/sqlc"
	"github.com/danilluk1/test-task-6/internal/app/api"
	router "github.com/danilluk1/test-task-6/internal/app/api/router"
	loggerimpl "github.com/danilluk1/test-task-6/internal/services/logger/impl"
)

func main() {
	cfg, err := config.New(true)
	if err != nil {
		panic(err)
	}
	conn, err := sql.Open("postgres", cfg.DbConn)
	if err != nil {
		panic(err)
	}

	store := db.NewStore(conn)

	app := &api.App{
		Store:  store,
		Logger: loggerimpl.NewLogger(),
	}

	router := router.Setup(app)

	srv := &http.Server{
		Addr:         "0.0.0.0:3002",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
