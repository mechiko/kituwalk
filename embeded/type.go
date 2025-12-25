package embeded

import (
	"embed"
)

//go:embed RobotoCondensed-Regular.ttf
var Regular []byte

//go:embed RobotoCondensed-Bold.ttf
var Bold []byte

//go:embed RobotoCondensed-Italic.ttf
var Italic []byte

//go:embed RobotoCondensed-BoldItalic.ttf
var BoldItalic []byte

//go:embed assets
var EmbeddedAssets embed.FS
