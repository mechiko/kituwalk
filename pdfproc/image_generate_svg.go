package pdfproc

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/juliankoehn/barcode"
)

// генератор из SVG
func svgImg(code string) ([]byte, error) {
	var b bytes.Buffer
	width := 2
	height := 35
	color := "black"
	showCode := false
	which := "EAN13"
	extension := ".jpg"
	if extension == ".png" || extension == ".jpg" || extension == ".jpeg" || extension == ".gif" {
		var img *image.RGBA
		if extension == ".png" {
			_, img = barcode.GetBarcodeFile(code, which, width, height, color, showCode, false, true)
		} else {
			_, img = barcode.GetBarcodeFile(code, which, width, height, color, showCode, false, false)
		}
		switch extension {
		case ".png":
			png.Encode(&b, img)
		case ".jpg", ".jpeg":
			jpeg.Encode(&b, img, &jpeg.Options{Quality: 100})
		case ".gif":
			gif.Encode(&b, img, nil)
		}
	}
	return b.Bytes(), nil
}
