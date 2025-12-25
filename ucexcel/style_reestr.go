package ucexcel

import (
	"github.com/xuri/excelize/v2"
)

var (
	sizeFontStyle               float64 = 8
	celStyleReport              int
	celStyleReportJust          int
	celStyleReportCenter        int
	celStyleReportBold          int
	celStyleReportBoldBlue      int
	celStyleReportRight         int
	celStyleReportRightTotal    int
	celStyleReportBoldRight     int
	celStyleReportRightNum      int
	celStyleReportBoldRightNum  int
	celStyleReportRightSum      int
	celStyleReportBoldRightSum  int
	celStyleReportRightNumTotal int
	celStyleReportRightSumTotal int
)

func (ue *ucexcel) styleReestrReport() {
	expNum := "0.000"
	expSum := "0.00"
	decPlace := 4
	celStyleReport, _ = ue.file.NewStyle(&excelize.Style{
		NumFmt:        0,
		DecimalPlaces: &decPlace,
		Alignment: &excelize.Alignment{
			Horizontal:      "left",
			Vertical:        "center",
			JustifyLastLine: false,
			ShrinkToFit:     false,
			WrapText:        true,
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
			Size:   sizeFontStyle,
		},
	})
	celStyleReportJust, _ = ue.file.NewStyle(&excelize.Style{
		NumFmt:        0,
		DecimalPlaces: &decPlace,
		Alignment: &excelize.Alignment{
			Horizontal:      "left",
			Vertical:        "center",
			JustifyLastLine: true,
			ShrinkToFit:     true,
			WrapText:        true,
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
			Size:   sizeFontStyle,
		},
	})
	celStyleReportBold, _ = ue.file.NewStyle(&excelize.Style{
		NumFmt:        0,
		DecimalPlaces: &decPlace,
		Alignment: &excelize.Alignment{
			Horizontal:      "left",
			Vertical:        "center",
			JustifyLastLine: false,
			ShrinkToFit:     false,
			WrapText:        true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold:   true,
			Family: "Arial",
			Size:   sizeFontStyle,
		},
	})
	celStyleReportRight, _ = ue.file.NewStyle(&excelize.Style{
		NumFmt:        0,
		DecimalPlaces: &decPlace,
		Alignment: &excelize.Alignment{
			Horizontal:      "right",
			Vertical:        "center",
			JustifyLastLine: true,
			ShrinkToFit:     true,
			WrapText:        true,
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
			Size:   sizeFontStyle,
		},
	})
	celStyleReportBoldRight, _ = ue.file.NewStyle(&excelize.Style{
		NumFmt:        0,
		DecimalPlaces: &decPlace,
		Alignment: &excelize.Alignment{
			Horizontal:      "right",
			Vertical:        "center",
			JustifyLastLine: true,
			ShrinkToFit:     true,
			WrapText:        true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold:   true,
			Family: "Arial",
			Size:   sizeFontStyle,
		},
	})
	celStyleReportRightNum, _ = ue.file.NewStyle(&excelize.Style{
		CustomNumFmt: &expNum,
		Alignment: &excelize.Alignment{
			Horizontal:      "right",
			Vertical:        "center",
			JustifyLastLine: true,
			ShrinkToFit:     true,
			WrapText:        true,
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
			Size:   sizeFontStyle,
		},
	})
	celStyleReportBoldRightNum, _ = ue.file.NewStyle(&excelize.Style{
		CustomNumFmt: &expNum,
		Alignment: &excelize.Alignment{
			Horizontal:      "right",
			Vertical:        "center",
			JustifyLastLine: true,
			ShrinkToFit:     true,
			WrapText:        true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold:   true,
			Family: "Arial",
			Size:   sizeFontStyle,
		},
	})
	celStyleReportRightSum, _ = ue.file.NewStyle(&excelize.Style{
		CustomNumFmt: &expSum,
		Alignment: &excelize.Alignment{
			Horizontal:      "right",
			Vertical:        "center",
			JustifyLastLine: true,
			ShrinkToFit:     true,
			WrapText:        true,
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
			Size:   sizeFontStyle,
		},
	})
	celStyleReportBoldRightSum, _ = ue.file.NewStyle(&excelize.Style{
		CustomNumFmt: &expSum,
		Alignment: &excelize.Alignment{
			Horizontal:      "right",
			Vertical:        "center",
			JustifyLastLine: true,
			ShrinkToFit:     true,
			WrapText:        true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold:   true,
			Family: "Arial",
			Size:   sizeFontStyle,
		},
	})
	celStyleReportBoldBlue, _ = ue.file.NewStyle(&excelize.Style{
		NumFmt:        0,
		DecimalPlaces: &decPlace,
		Alignment: &excelize.Alignment{
			Horizontal:      "left",
			Vertical:        "center",
			JustifyLastLine: false,
			ShrinkToFit:     false,
			WrapText:        true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold:   true,
			Family: "Arial",
			Size:   sizeFontStyle,
			Color:  "056c91",
		},
	})
	celStyleReportCenter, _ = ue.file.NewStyle(&excelize.Style{
		NumFmt:        0,
		DecimalPlaces: &decPlace,
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Vertical:        "center",
			JustifyLastLine: false,
			ShrinkToFit:     false,
			WrapText:        true,
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
			Size:   sizeFontStyle,
		},
	})
	celStyleReportRightTotal, _ = ue.file.NewStyle(&excelize.Style{
		NumFmt:        0,
		DecimalPlaces: &decPlace,
		Alignment: &excelize.Alignment{
			Horizontal:      "right",
			Vertical:        "center",
			JustifyLastLine: true,
			ShrinkToFit:     true,
			WrapText:        true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold:   false,
			Italic: true,
			Family: "Arial",
			Size:   sizeFontStyle,
		},
	})
	celStyleReportRightSumTotal, _ = ue.file.NewStyle(&excelize.Style{
		CustomNumFmt: &expSum,
		Alignment: &excelize.Alignment{
			Horizontal:      "right",
			Vertical:        "center",
			JustifyLastLine: true,
			ShrinkToFit:     true,
			WrapText:        true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold:   false,
			Italic: true,
			Family: "Arial",
			Size:   sizeFontStyle,
		},
	})
	celStyleReportRightNumTotal, _ = ue.file.NewStyle(&excelize.Style{
		CustomNumFmt: &expNum,
		Alignment: &excelize.Alignment{
			Horizontal:      "right",
			Vertical:        "center",
			JustifyLastLine: true,
			ShrinkToFit:     true,
			WrapText:        true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold:   false,
			Italic: true,
			Family: "Arial",
			Size:   sizeFontStyle,
		},
	})

}
