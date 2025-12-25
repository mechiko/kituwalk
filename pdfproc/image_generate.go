package pdfproc

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func forceTo8BitPNG(data []byte) ([]byte, error) {
	src, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("decode png: %w", err)
	}
	b := src.Bounds()
	dst := image.NewNRGBA(b)
	draw.Draw(dst, b, src, b.Min, draw.Src)
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, dst); err != nil {
		return nil, fmt.Errorf("encode 8-bit png: %w", err)
	}
	return buf.Bytes(), nil
}

// генератор code128
func bar128Img(code string, h float64) ([]byte, error) {
	bcImg, err := code128.Encode(code)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	dx := bcImg.Bounds().Dx()
	dy := int(h * 3.528) // точек в мм при дпи 72 (254/72)
	scaled, err := barcode.Scale(bcImg, dx, dy)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	var b bytes.Buffer
	png.Encode(&b, scaled)
	return forceTo8BitPNG(b.Bytes())
}
