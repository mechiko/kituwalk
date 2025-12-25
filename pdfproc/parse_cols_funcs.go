package pdfproc

import (
	"fmt"
	"kituwalk/domain"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

func (p *pdfProc) addBarColumn(col core.Col, colTempl *domain.RowPrimitive, code *domain.Pallete) error {
	switch colTempl.Bar {
	case "code128":
		img, err := bar128Img(code.Code, colTempl.RowHeight)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		col.Add(
			image.NewFromBytes(img, extension.Png, colTempl.PropsRect()),
		)
	}
	return nil
}

func (p *pdfProc) addJpgColumn(col core.Col, colTempl *domain.RowPrimitive, _ *domain.Pallete) error {
	if p.assets != nil {
		img, err := p.assets.Jpg(colTempl.Image)
		if err != nil {
			return fmt.Errorf("page image assets %w", err)
		}
		if len(img) == 0 {
			return fmt.Errorf("page image assets empty for %q", colTempl.Image)
		}
		col.Add(
			image.NewFromBytes(img, colTempl.ImageExt, colTempl.PropsRect()),
		)
	} else {
		return fmt.Errorf("page image assets not available (assets is nil) for %q", colTempl.Image)
	}
	return nil
}

func (p *pdfProc) addStringColumn(col core.Col, colTempl *domain.RowPrimitive, code *domain.Pallete) error {
	value := strings.ReplaceAll(colTempl.Value, "@code", code.Code)
	value = strings.ReplaceAll(value, "@box", code.Box)
	value = strings.ReplaceAll(value, "@name", code.Name)
	value = strings.ReplaceAll(value, "@vol", code.Volume)
	col.Add(text.New(value, colTempl.PropsText()))
	return nil
}

func (p *pdfProc) addArrayStringColumn(col core.Col, colTempl *domain.RowPrimitive, code *domain.Pallete) error {
	comps := make([]core.Component, 0)
	for _, val := range colTempl.Values {
		value := ""
		if val.Value != "" {
			value = strings.ReplaceAll(val.Value, "@code", code.Code)
			value = strings.ReplaceAll(value, "@box", code.Box)
			value = strings.ReplaceAll(value, "@name", code.Name)
			value = strings.ReplaceAll(value, "@vol", code.Volume)
			comps = append(comps, text.New(value, val.PropsText()))
		}
		if val.Bar != "" {
			switch val.Bar {
			}
		}
	}
	col.Add(comps...)
	return nil
}
