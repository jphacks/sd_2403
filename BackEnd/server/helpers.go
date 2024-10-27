/*

 */

package server

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func makeHttpServer(addr string, r http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func handleGracefulShutdown(server *http.Server, closers []io.Closer) {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("receive interrupt signal")
	for _, c := range closers {
		err := c.Close()
		if err != nil {
			log.Println("error during closing: ", err)
		}
	}
	if err := server.Close(); err != nil {
		log.Fatal("Server Close:", err)
	}
}
