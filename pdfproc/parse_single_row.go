package pdfproc

import (
	"fmt"
	"kituwalk/domain"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

// если не указывается высота то вставляется авто роу не важно на остальное
// и текст из value
func (p *pdfProc) parseSingleRow(pg core.Page, row1 *domain.RowPrimitive, code *domain.Pallete) error {
	if row1.Value == "" && row1.Bar == "" {
		// пустая строка с высотой
		pg.Add(
			row.New(row1.RowHeight).Add(),
		)
	} else {
		if row1.RowHeight == 0 {
			value := strings.ReplaceAll(row1.Value, "@code", code.Code)
			value = strings.ReplaceAll(value, "@box", code.Box)
			value = strings.ReplaceAll(value, "@name", code.Name)
			value = strings.ReplaceAll(value, "@vol", code.Volume)
			pg.Add(
				text.NewAutoRow(value, row1.PropsText()),
			)
		} else {
			colNew := col.New(12)
			if row1.Bar != "" {
				img, err := bar128Img(code.Code, row1.RowHeight)
				if err != nil {
					return fmt.Errorf("%w", err)
				}
				colNew.Add(image.NewFromBytes(img, extension.Png, row1.PropsRect()))
				pg.Add(
					row.New(row1.RowHeight).Add(
						colNew,
					),
				)
			} else {
				value := strings.ReplaceAll(row1.Value, "@code", code.Code)
				value = strings.ReplaceAll(value, "@box", code.Box)
				value = strings.ReplaceAll(value, "@name", code.Name)
				value = strings.ReplaceAll(value, "@vol", code.Volume)
				pg.Add(
					row.New(row1.RowHeight).Add(
						text.NewCol(12, value, row1.PropsText()),
					),
				)
			}
		}
	}
	return nil
}
