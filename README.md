# Gotenberg Client

Go client for [Gotenberg](https://gotenberg.dev/) — document conversion service supporting Chromium, LibreOffice, and PDF manipulation engines.

```bash
go get github.com/nativebpm/gotenberg
```

## Features

- **Chromium**: Convert URLs, HTML, and Markdown to PDF
- **LibreOffice**: Convert Office documents (Word, Excel, PowerPoint) to PDF
- **PDF Engines**: Merge, split, and manipulate PDFs
- **Webhook support**: Async conversions with callback URLs
- **Stream-first**: Built on `httpstream` for efficient multipart uploads

## Conversion Engines

### Chromium

Convert web content to PDF:
- URL → PDF
- HTML file → PDF
- Markdown → PDF

Supports:
- Custom page properties (size, margins, orientation)
- Headers & footers with page numbers
- Wait strategies (delay, JavaScript expression)
- Cookies & custom HTTP headers
- Emulated media types (screen/print)

### LibreOffice

Convert Office documents:
- Word (.docx, .doc) → PDF
- Excel (.xlsx, .xls) → PDF
- PowerPoint (.pptx, .ppt) → PDF
- OpenDocument formats → PDF

### PDF Engines

PDF operations:
- Merge multiple PDFs
- Split pages
- Convert images to PDF

## Webhook Mode

Async conversions with callbacks:
- Returns `204 No Content` immediately
- Uploads result to webhook URL in background
- Separate error callback URL
- Custom HTTP headers for callbacks

See [Gotenberg webhook docs](https://gotenberg.dev/docs/webhook) for details.

## Examples

- [Chromium: URL to PDF](examples/cmd/chromium/converturl)
- [Chromium: Hello World](examples/cmd/chromium/helloworld)
- [Chromium: Markdown to PDF](examples/cmd/chromium/markdown)
- [Chromium: Timeout handling](examples/cmd/chromium/timeout)
- [Chromium: Trace header](examples/cmd/chromium/trace)
- [Chromium: Webhook async](examples/cmd/chromium/webhook)
- [LibreOffice: Document conversion](examples/cmd/libreoffice/convert)
- [PDF Engines: Merge PDFs](examples/cmd/pdfengines/merge)
- [Health check](examples/cmd/health)

## License

MIT — see [`LICENSE`](../LICENSE).
