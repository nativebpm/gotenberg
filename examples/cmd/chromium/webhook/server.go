package main

import (
	"log"
	"log/slog"
	"net/http"
)

func StartServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/success", webhookHandler("success"))
	mux.HandleFunc("/error", webhookHandler("error"))

	srv := &http.Server{Addr: addr, Handler: mux}
	go func() {
		slog.Info("starting local webhook server", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("webhook server error: %v", err)
		}
	}()

	return srv
}
