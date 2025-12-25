package pdfproc

import (
	"fmt"
	"kituwalk/domain"
	"slices"

	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

// в станицу передаются только cis для нее
func (p *pdfProc) Page(t *domain.MarkTemplate, code *domain.Pallete) (core.Page, error) {
	pg := page.New()
	rowKeys := make([]string, 0, len(t.Rows))
	for k := range t.Rows {
		rowKeys = append(rowKeys, k)
	}
	slices.Sort(rowKeys)
	for _, rowKey := range rowKeys {
		rowTempl := t.Rows[rowKey]
		switch {
		case len(rowTempl) == 0:
		case len(rowTempl) == 1:
			row1 := rowTempl[0]
			// одна строка автороу
			if err := p.parseSingleRow(pg, row1, code); err != nil {
				return nil, fmt.Errorf("parse single row error %w", err)
			}
		case len(rowTempl) > 1:
			if err := p.parseColsRow(pg, rowTempl, code); err != nil {
				return nil, fmt.Errorf("parse cols row error %w", err)
			}
		default:
		}
	}
	return pg, nil
}
