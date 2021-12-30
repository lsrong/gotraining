package service

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/lsrong/kit/web"
)

func Start() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	api := web.NewServer(shutdown)
	api.Handle("POST", "", "/search", searchHandler)
	api.Handle("GET", "", "/search", searchHandler)
	api.Handle("GET", "", "/static/*filepath", staticHandler)
	api.Handle("GET", "", "/", indexHandler)

	host := "localhost:4000"
	readTimeout := 10 * time.Second
	writeTimeout := 30 * time.Second
	idleTimeout := 60 * time.Second
	svc := http.Server{
		Addr:           host,
		Handler:        api,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
		IdleTimeout:    idleTimeout,
	}

	errCh := make(chan error)

	go func() {
		log.Printf("Service Started, Listening on http://%s \n", host)
		errCh <- svc.ListenAndServe()
	}()

	select {
	case <-shutdown:
		log.Println("Starting shutdown...")
		svc.Close()
	case err := <-errCh:
		log.Printf("Api Server Error: %v \n", err)
	}
}
