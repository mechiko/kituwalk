package pdfproc

import (
	"fmt"
	"kituwalk/domain"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

// ciss КМ передаются в строку колонок или одна или несколько
func (p *pdfProc) parseColsRow(pg core.Page, colsTempl []*domain.RowPrimitive, code *domain.Pallete) error {
	cols := make([]core.Col, 0)
	// строки с колонками
	for i, colSingle := range colsTempl {
		if colSingle.ColWidth == 0 {
			continue
		}
		colNew := col.New(colSingle.ColWidth)
		cols = append(cols, colNew)

		switch {
		case colSingle.Bar != "":
			if err := p.addBarColumn(colNew, colSingle, code); err != nil {
				return fmt.Errorf("generate col %d %w", i, err)
			}
		case colSingle.Image != "":
			if err := p.addJpgColumn(colNew, colSingle, code); err != nil {
				return fmt.Errorf("generate col %d %w", i, err)
			}
		case colSingle.Value != "":
			if err := p.addStringColumn(colNew, colSingle, code); err != nil {
				return fmt.Errorf("generate col %d %w", i, err)
			}
		case len(colSingle.Values) > 0:
			if err := p.addArrayStringColumn(colNew, colSingle, code); err != nil {
				return fmt.Errorf("generate col %d %w", i, err)
			}
		}
	}
	height := 0.0
	if len(colsTempl) > 0 {
		height = colsTempl[0].RowHeight
	}
	pg.Add(
		row.New(height).Add(cols...),
	)
	return nil
}
