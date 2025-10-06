package main

import (
	"io"
	"log"
	"log/slog"
	"mime"
	"net/http"
	"os"
	"sync"
)

var fileMutex sync.Mutex

func webhookHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if name != "success" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				slog.Error("webhook", "error", "failed to read body", "read_error", err)
			} else {
				slog.Error("webhook", "error", string(body))
			}
			http.Error(w, "error", http.StatusBadRequest)
			return
		}

		filename := filename(r.Header)

		fileMutex.Lock()
		defer fileMutex.Unlock()

		outFile, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()

		n, err := io.Copy(outFile, r.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()

		slog.Info("webhook",
			"gotenberg-trace", r.Header.Get("Gotenberg-Trace"),
			"received", name,
			"method", r.Method,
			"path", r.URL.Path,
			"content length", n,
			"x-custom-header", r.Header.Get("X-Custom-Header"),
			"x-custom-header2", r.Header.Get("X-Custom-Header2"),
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}
}

func filename(headers http.Header) string {
	if v := headers.Get("Content-Disposition"); v != "" {
		_, params, err := mime.ParseMediaType(v)
		if err == nil {
			if fn, ok := params["filename"]; ok {
				return fn
			}
		}
	}
	return "file.pdf"
}
