package ucexcel

import (
	"kituwalk/ucexcel/address"

	"github.com/xuri/excelize/v2"
)

type ucexcel struct {
	layout             string
	template           string
	name               string
	address            address.ExcelAddress
	sheet              string
	file               *excelize.File
	celStyleDefault    int
	celStyleLeft       int
	celStyleVCenter    int
	celStyleBold       int
	celStyleBoldGreen  int
	celStyleBoldBlue   int
	celStyleBoldRight  int
	celStyleBoldCenter int
	celStyle           excelize.Style
}

func New(name string) *ucexcel {
	excel := &ucexcel{
		name:    name,
		address: address.New(14, 0),
	}
	excel.celStyle = excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold:   false,
			Family: "Arial",
			Color:  "000000",
			Size:   9,
		},
	}

	return excel
}
