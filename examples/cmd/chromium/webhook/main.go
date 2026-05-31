package main

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/nativebpm/gotenberg"
	"github.com/nativebpm/gotenberg/examples/model"
	"github.com/nativebpm/gotenberg/examples/pkg/image"
	"github.com/nativebpm/gotenberg/examples/pkg/templates/invoice"
)

// cleanupPDFFiles removes all PDF files from the current directory
func cleanupPDFFiles() {
	entries, err := os.ReadDir(".")
	if err != nil {
		slog.Warn("failed to read current directory", "err", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if strings.HasSuffix(strings.ToLower(name), ".pdf") {
			if err := os.Remove(name); err != nil {
				slog.Warn("failed to remove PDF file", "file", name, "err", err)
			} else {
				slog.Info("removed PDF file", "file", name)
			}
		}
	}
}

func main() {
	cleanupPDFFiles()

	srv := StartServer(":28080")

	gotenbergURL := `http://localhost:3000`

	httpClient := http.Client{
		Timeout: 90 * time.Second,
	}

	logo := image.LogoPNG()

	client, err := gotenberg.NewClient(httpClient, gotenbergURL)
	if err != nil {
		slog.Error("failed to create Gotenberg client", "err", err)
		return
	}

	go func() { // Example #1:
		data := model.InvoiceData
		html := bytes.NewBuffer(nil)
		if err := invoice.Template.Execute(html, data); err != nil {
			slog.Error("failed to execute invoice template", "err", err)
			return
		}

		resp, err := client.Chromium().
			ConvertHTML(context.Background(), html).
			File("logo.png", bytes.NewReader(logo)).
			PrintBackground().
			WebhookURL("http://host.docker.internal:28080/success", http.MethodPost).
			WebhookErrorURL("http://host.docker.internal:28080/error", http.MethodPost).
			WebhookHeader("X-Custom-Header", "MyValue").
			OutputFilename("invoice_async").
			Send()

		if err != nil {
			slog.Error("failed to convert HTML to PDF", "err", err)
			return
		}
		defer resp.Body.Close()

		slog.Info("Async HTML to PDF conversion started",
			"gotenberg-trace", resp.GotenbergTrace)
	}()

	go func() { // Example #2:
		time.Sleep(1 * time.Second)
		data := model.InvoiceData
		html := bytes.NewBuffer(nil)
		if err := invoice.Template.Execute(html, data); err != nil {
			slog.Error("failed to execute invoice template", "err", err)
			return
		}

		resp, err := client.Chromium().
			ConvertHTML(context.Background(), html).
			File("logo.png", bytes.NewReader(logo)).
			PrintBackground().
			WebhookURL("http://host.docker.internal:28080/success", http.MethodPost).
			WebhookErrorURL("http://host.docker.internal:28080/error", http.MethodPost).
			WebhookHeader("X-Custom-Header", "MyValue").
			WebhookHeader("X-Custom-Header2", "MyValue2").
			Margins(1.0, 1.5, 1.0, 1.5).
			OutputFilename("invoice_async2").
			Send()

		if err != nil {
			slog.Error("failed to convert HTML to PDF", "err", err)
			return
		}
		defer resp.Body.Close()

		slog.Info("Async HTML to PDF conversion started",
			"gotenberg-trace", resp.GotenbergTrace)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	slog.Info("waiting for webhook callbacks; press Ctrl+C to exit")
	<-sigCh

	slog.Info("shutting down webhook server")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Info("error shutting down server", "err", err)
	}
}
