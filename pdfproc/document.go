package pdfproc

import (
	"fmt"
	"kituwalk/domain"
	"kituwalk/embeded"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/repository"

	"github.com/johnfercher/maroto/v2/pkg/props"
)

func (p *pdfProc) BuildPages(codes []*domain.Pallete) error {
	for _, cd := range codes {
		if err := p.addPageByTemplate(p.tmplt, cd); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}

func (p *pdfProc) PdfDocumentSave(fileName string) (err error) {
	if p.document == nil {
		return fmt.Errorf("save document: document is nil")
	}
	err = p.document.Save(fileName)
	if err != nil {
		return fmt.Errorf("save document: %w", err)
	}
	return nil
}

func (p *pdfProc) PdfDocumentReportSave(fileName string) (err error) {
	// запись отчета генерации
	if p.document == nil {
		return fmt.Errorf("save document report: document is nil")
	}
	err = p.document.GetReport().Save(fileName)
	if err != nil {
		return fmt.Errorf("save document report: %w", err)
	}
	return nil
}

func (p *pdfProc) BuildMaroto() (err error) {
	customFont := "roboto"
	customFonts, err := repository.New().
		AddUTF8FontFromBytes(customFont, fontstyle.Normal, embeded.Regular).
		AddUTF8FontFromBytes(customFont, fontstyle.Italic, embeded.Italic).
		AddUTF8FontFromBytes(customFont, fontstyle.Bold, embeded.Bold).
		AddUTF8FontFromBytes(customFont, fontstyle.BoldItalic, embeded.BoldItalic).
		Load()
	if err != nil {
		return err
	}
	// …custom font setup…
	builder := config.NewBuilder().WithCustomFonts(customFonts).WithCompression(true)
	cfg := builder.WithDefaultFont(&props.Font{Family: customFont}).Build()
	cfg.Dimensions.Height = p.tmplt.PageHeight
	cfg.Dimensions.Width = p.tmplt.PageWidth
	cfg.Margins.Bottom = p.tmplt.Margin.Bottom
	cfg.Margins.Top = p.tmplt.Margin.Top
	cfg.Margins.Left = p.tmplt.Margin.Left
	cfg.Margins.Right = p.tmplt.Margin.Right
	cfg.DefaultFont.Size = 4
	p.maroto = maroto.New(cfg)
	p.maroto = maroto.NewMetricsDecorator(p.maroto)
	return nil
}

func (p *pdfProc) DocumentGenerate() (err error) {
	if p.maroto == nil {
		return fmt.Errorf("document generate: maroto is not initialized (call BuildMaroto first)")
	}
	doc, err := p.maroto.Generate()
	if err != nil {
		return fmt.Errorf("document generate: %w", err)
	}
	p.document = doc
	return nil
}

func (p *pdfProc) addPageByTemplate(tmpl *domain.MarkTemplate, cis *domain.Pallete) error {
	if tmpl == nil {
		return fmt.Errorf("add page: template is nil")
	}
	if p.maroto == nil {
		return fmt.Errorf("add page: maroto is not initialized (call BuildMaroto first)")
	}
	pgNew, err := p.Page(tmpl, cis)
	if err != nil {
		return fmt.Errorf("build page: %w", err)
	}
	if pgNew == nil {
		return fmt.Errorf("build page: result is nil")
	}
	p.maroto.AddPages(pgNew)
	return nil
}
