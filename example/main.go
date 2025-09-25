package main

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/nativebpm/gotenberg-client"
	"github.com/nativebpm/gotenberg-client/example/model"
	"github.com/nativebpm/gotenberg-client/example/templates/invoice"
)

func main() {
	gotenbergURL := `http://localhost:3000`

	httpClient := &http.Client{
		Timeout: 90 * time.Second,
	}

	client := gotenberg.NewClient(httpClient, gotenbergURL)

	err := convertURLToPDF(client)
	if err != nil {
		slog.Error("Failed to convert URL to PDF", "error", err)
		return
	}

	data := model.InvoiceData
	html := makeHtml(data)

	err = convertHTMLToPDF(client, html)
	if err != nil {
		slog.Error("Failed to convert HTML to PDF", "error", err)
		return
	}

	err = convertHTMLToPDFWithWebhook(client, html)
	if err != nil {
		slog.Error("Failed to convert HTML to PDF with webhook", "error", err)
		return
	}

	err = convertHTMLMinimal(client, html)
	if err != nil {
		slog.Error("Failed minimal HTML conversion", "error", err)
		return
	}
}

func convertHTMLToPDF(client *gotenberg.Client, htmlDoc *bytes.Buffer) error {
	slog.Info("Converting HTML to PDF with options...")

	logoPNG := generateLogoPNG()
	files := map[string][]byte{"logo.png": logoPNG}

	resp, err := client.ConvertHTMLToPDF(htmlDoc.Bytes(),
		gotenberg.WithPrintBackground(true),
		gotenberg.WithLandscape(false),
		gotenberg.WithScale(1.0),
		gotenberg.WithOutputFilename("invoice.pdf"),
		gotenberg.WithPaperSizeA4(),
		gotenberg.WithMargins(1.0, 1.0, 1.0, 1.0),
		gotenberg.WithHTMLAdditionalFiles(files),
	)
	if err != nil {
		return err
	}

	err = os.WriteFile("./invoice_new.pdf", resp.PDF, 0644)
	if err != nil {
		return err
	}

	slog.Info("HTML to PDF conversion completed",
		"pdf_size", len(resp.PDF),
		"content_type", resp.ContentType,
		"trace", resp.Trace)

	return nil
}

func convertURLToPDF(client *gotenberg.Client) error {
	slog.Info("Converting URL to PDF...")

	resp, err := client.ConvertURLToPDF("https://example.com",
		gotenberg.WithPrintBackground(true),
		gotenberg.WithOutputFilename("example.pdf"),
		gotenberg.WithPaperSizeLetter(),
	)
	if err != nil {
		return err
	}

	err = os.WriteFile("./example.pdf", resp.PDF, 0644)
	if err != nil {
		return err
	}

	slog.Info("URL to PDF conversion completed",
		"pdf_size", len(resp.PDF),
		"trace", resp.Trace)

	return nil
}

func convertHTMLToPDFWithWebhook(client *gotenberg.Client, htmlDoc *bytes.Buffer) error {
	slog.Info("Converting HTML to PDF with webhook (async)...")

	logo := generateLogoPNG()
	files := map[string][]byte{"logo.png": logo}

	resp, err := client.ConvertHTMLToPDF(htmlDoc.Bytes(),
		gotenberg.WithPrintBackground(true),
		gotenberg.WithOutputFilename("invoice_async.pdf"),
		gotenberg.WithWebhook(
			"https://your-webhook-url.com/success",
			"https://your-webhook-url.com/error",
		),
		gotenberg.WithWebhookMethods("POST", "POST"),
		gotenberg.WithWebhookExtraHeaders(map[string]string{
			"Authorization":   "Bearer your-token",
			"X-Custom-Header": "custom-value",
		}),
		gotenberg.WithHTMLAdditionalFiles(files),
	)
	if err != nil {
		return err
	}

	slog.Info("Async HTML to PDF conversion started",
		"trace", resp.Trace,
		"pdf_returned", resp.PDF != nil)

	return nil
}

func convertHTMLMinimal(client *gotenberg.Client, htmlDoc *bytes.Buffer) error {
	slog.Info("Converting HTML to PDF (minimal, no options)...")

	logoImage := generateLogoPNG()
	files := map[string][]byte{"logo.png": logoImage}

	resp, err := client.ConvertHTMLToPDF(htmlDoc.Bytes(),
		gotenberg.WithHTMLAdditionalFiles(files),
	)
	if err != nil {
		return err
	}

	err = os.WriteFile("./invoice_minimal.pdf", resp.PDF, 0644)
	if err != nil {
		return err
	}

	slog.Info("Minimal HTML to PDF conversion completed",
		"pdf_size", len(resp.PDF),
		"trace", resp.Trace)

	return nil
}

func makeHtml(data model.Invoice) *bytes.Buffer {
	buf := bytes.NewBuffer(nil)
	invoice.Template.Execute(buf, data)
	return buf
}

func generateLogoPNG() []byte {
	const w, h = 300, 80
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)

	blue := color.RGBA{R: 10, G: 102, B: 194, A: 255}
	rect := image.Rect(10, 10, 90, h-10)
	draw.Draw(img, rect, &image.Uniform{C: blue}, image.Point{}, draw.Src)

	gray := color.RGBA{R: 200, G: 200, B: 200, A: 255}
	draw.Draw(img, image.Rect(110, 20, w-10, 36), &image.Uniform{C: gray}, image.Point{}, draw.Src)
	draw.Draw(img, image.Rect(110, 44, w-60, 60), &image.Uniform{C: gray}, image.Point{}, draw.Src)

	var buf bytes.Buffer
	png.Encode(&buf, img)
	return buf.Bytes()
}
