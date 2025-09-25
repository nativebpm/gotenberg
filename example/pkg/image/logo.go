package image

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
)

func LogoPNG() []byte {
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
