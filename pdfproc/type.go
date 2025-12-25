package pdfproc

import (
	"encoding/json"
	"fmt"
	"kituwalk/assets"
	"kituwalk/domain"

	"github.com/johnfercher/maroto/v2/pkg/core"
)

type pdfProc struct {
	domain.Apper
	maroto   core.Maroto
	assets   *assets.Assets
	document core.Document
	debug    bool
	height   float64
	width    float64
	tmplt    *domain.MarkTemplate
}

func New(app domain.Apper) (*pdfProc, error) {
	if app == nil {
		return nil, fmt.Errorf("app is nil")
	}
	asts, err := assets.New("assets")
	if err != nil {
		return nil, fmt.Errorf("Error assets: %w", err)
	}
	tmplDatamatrixJson, err := asts.Json("datamatrix")
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %w", err)
	}
	tmplDatamatrix := &domain.MarkTemplate{}
	err = json.Unmarshal(tmplDatamatrixJson, tmplDatamatrix)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshal datamatrix file: %w", err)
	}
	p := &pdfProc{
		Apper:  app,
		assets: asts,
		tmplt:  tmplDatamatrix,
	}
	return p, nil
}
