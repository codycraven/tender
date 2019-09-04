package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codycraven/tender/pkg/liveprobe"
	"github.com/ghodss/yaml"
)

// A Tender provides attachment of a handler to http.
type Tender interface {
	DeployTender(path, route string, mux *http.ServeMux) error
}

var (
	configFile    = flag.String("config-file", "config.yml", "Path to config file")
	liveCheck     = flag.Bool("livecheck", false, "check if the service is alive")
	liveCheckFile = "./isalive"
	healthy       = true
)

func main() {
	flag.Parse()

	if *liveCheck {
		if liveprobe.CheckLiveliness(liveCheckFile) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	log.Println("Loading config")
	b, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	var cfg config
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	cfg.Mux.Handle("/healthz", healthCheck())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: cfg.Mux,
	}

	idleConnsClosed := make(chan struct{})
	// Listen for shutdown signal.
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		// We received an interrupt or SIGTERM signal, shut down.
		log.Println("graceful shutdown initiated")
		healthy = false
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		// Close and return the channel.
		close(idleConnsClosed)
	}()

	liveprobe.MakeAlive(liveCheckFile)
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}

	// When this channel returns, the server has gracefully shut down the web server.
	<-idleConnsClosed
	liveprobe.MakeDead(liveCheckFile)
}

func healthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if healthy {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "OK")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Not Healthy")
		}
	})
}
