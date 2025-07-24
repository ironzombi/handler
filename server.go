package main

import (
	"log"
	"net"
	"net/http"
	"network_go/ch9/handler/handlers"
	"sync"
	"time"
)

/* http handler */
func main() {
	var wg sync.WaitGroup
	srv := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: http.TimeoutHandler(handlers.DefaultHandler(), 2*time.Minute, ""),
		//Handler:           http.TimeoutHandler(handlers.DefaultMethodsHandler(), 2*time.Minute, ""),
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	l, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go func() {

		err := srv.Serve(l)
		if err != http.ErrServerClosed {
			error.Error(err)
		}
	}()

	wg.Wait()

}
