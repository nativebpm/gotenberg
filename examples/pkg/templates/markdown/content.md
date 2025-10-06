# Welcome to Gotenberg Markdown Example

This is a demonstration of converting **Markdown** to *PDF* using Gotenberg.

## Features

The Gotenberg client supports:

- Converting Markdown files to PDF
- Custom HTML templates
- Styling with CSS
- Multiple markdown files

## Code Example

Here's how to use the client:

```go
client, err := gotenberg.NewClient(httpClient, "http://localhost:3000")
if err != nil {
    log.Fatal(err)
}
ctx := context.Background()
request := client.Chromium().
  ConvertMarkdown(ctx, bytes.NewReader(indexHTML)).
  File("content.md", bytes.NewReader(markdownContent))
```

## Math Support

Gotenberg supports MathJax for mathematical expressions:

$$E = mc^2$$

And inline math: $x = \frac{-b \pm \sqrt{b^2 - 4ac}}{2a}$

## Lists

### Unordered List
- Item 1
- Item 2
  - Nested item
  - Another nested item

### Ordered List
1. First item
2. Second item
3. Third item

## Blockquote

> This is a blockquote. It can span multiple lines and provides
> a way to highlight important information or quotes.

## Table Example

| Feature | Supported | Notes |
|---------|-----------|-------|
| Headers | ✅ | H1-H6 |
| Lists | ✅ | Ordered and unordered |
| Code | ✅ | Inline and blocks |
| Math | ✅ | MathJax support |

## Conclusion

This example demonstrates the power and flexibility of the Gotenberg
markdown conversion feature.