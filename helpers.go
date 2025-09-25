package gotenberg

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

// pageProperties internal type for page parameters
type pageProperties struct {
	SinglePage              *bool
	PaperWidth              *float64
	PaperHeight             *float64
	MarginTop               *float64
	MarginBottom            *float64
	MarginLeft              *float64
	MarginRight             *float64
	PreferCSSPageSize       *bool
	GenerateDocumentOutline *bool
	GenerateTaggedPDF       *bool
	PrintBackground         *bool
	OmitBackground          *bool
	Landscape               *bool
	Scale                   *float64
	NativePageRanges        *string
}

// webhookOptions internal type for webhook parameters
type webhookOptions struct {
	URL          *string
	ErrorURL     *string
	Method       *string
	ErrorMethod  *string
	ExtraHeaders map[string]string
}

// addFileField adds file to multipart form
func (c *Client) addFileField(writer *multipart.Writer, fieldName, filename string, content []byte) error {
	part, err := writer.CreateFormFile(fieldName, filename)
	if err != nil {
		return err
	}

	_, err = part.Write(content)
	return err
}

// addPageProperties adds page properties to form
func (c *Client) addPageProperties(writer *multipart.Writer, props pageProperties) error {
	if props.SinglePage != nil {
		if err := writer.WriteField("singlePage", strconv.FormatBool(*props.SinglePage)); err != nil {
			return err
		}
	}

	if props.PaperWidth != nil {
		if err := writer.WriteField("paperWidth", strconv.FormatFloat(*props.PaperWidth, 'f', -1, 64)); err != nil {
			return err
		}
	}

	if props.PaperHeight != nil {
		if err := writer.WriteField("paperHeight", strconv.FormatFloat(*props.PaperHeight, 'f', -1, 64)); err != nil {
			return err
		}
	}

	if props.MarginTop != nil {
		if err := writer.WriteField("marginTop", strconv.FormatFloat(*props.MarginTop, 'f', -1, 64)); err != nil {
			return err
		}
	}

	if props.MarginBottom != nil {
		if err := writer.WriteField("marginBottom", strconv.FormatFloat(*props.MarginBottom, 'f', -1, 64)); err != nil {
			return err
		}
	}

	if props.MarginLeft != nil {
		if err := writer.WriteField("marginLeft", strconv.FormatFloat(*props.MarginLeft, 'f', -1, 64)); err != nil {
			return err
		}
	}

	if props.MarginRight != nil {
		if err := writer.WriteField("marginRight", strconv.FormatFloat(*props.MarginRight, 'f', -1, 64)); err != nil {
			return err
		}
	}

	if props.PreferCSSPageSize != nil {
		if err := writer.WriteField("preferCssPageSize", strconv.FormatBool(*props.PreferCSSPageSize)); err != nil {
			return err
		}
	}

	if props.GenerateDocumentOutline != nil {
		if err := writer.WriteField("generateDocumentOutline", strconv.FormatBool(*props.GenerateDocumentOutline)); err != nil {
			return err
		}
	}

	if props.GenerateTaggedPDF != nil {
		if err := writer.WriteField("generateTaggedPdf", strconv.FormatBool(*props.GenerateTaggedPDF)); err != nil {
			return err
		}
	}

	if props.PrintBackground != nil {
		if err := writer.WriteField("printBackground", strconv.FormatBool(*props.PrintBackground)); err != nil {
			return err
		}
	}

	if props.OmitBackground != nil {
		if err := writer.WriteField("omitBackground", strconv.FormatBool(*props.OmitBackground)); err != nil {
			return err
		}
	}

	if props.Landscape != nil {
		if err := writer.WriteField("landscape", strconv.FormatBool(*props.Landscape)); err != nil {
			return err
		}
	}

	if props.Scale != nil {
		if err := writer.WriteField("scale", strconv.FormatFloat(*props.Scale, 'f', -1, 64)); err != nil {
			return err
		}
	}

	if props.NativePageRanges != nil {
		if err := writer.WriteField("nativePageRanges", *props.NativePageRanges); err != nil {
			return err
		}
	}

	return nil
}

// addWebhookHeaders adds headers for webhook
func (c *Client) addWebhookHeaders(req *http.Request, opts webhookOptions) {
	if opts.URL != nil {
		req.Header.Set("Gotenberg-Webhook-Url", *opts.URL)
	}

	if opts.ErrorURL != nil {
		req.Header.Set("Gotenberg-Webhook-Error-Url", *opts.ErrorURL)
	}

	if opts.Method != nil {
		req.Header.Set("Gotenberg-Webhook-Method", strings.ToUpper(*opts.Method))
	}

	if opts.ErrorMethod != nil {
		req.Header.Set("Gotenberg-Webhook-Error-Method", strings.ToUpper(*opts.ErrorMethod))
	}

	if len(opts.ExtraHeaders) > 0 {
		extraHeaders, err := json.Marshal(opts.ExtraHeaders)
		if err == nil {
			req.Header.Set("Gotenberg-Webhook-Extra-Http-Headers", string(extraHeaders))
		}
	}
}
