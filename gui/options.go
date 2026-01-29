package gui

import (
	"fmt"
	"kituwalk/reductor"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (g *gui) optionsDialog() (out int, err error) {
	var acceptPB, cancelPB *walk.PushButton
	var dlg *walk.Dialog
	var perPallet *walk.NumberEdit
	var startPallet *walk.NumberEdit
	var ssccPrefix *walk.TextEdit

	icon, err := walk.Resources.Icon("3")
	if err != nil {
		return -1, fmt.Errorf("%w", err)
	}
	model := reductor.Instance().Model("")
	if err := (dcl.Dialog{
		AssignTo: &dlg,
		Title:    "Настройка",
		// Size:          dcl.Size{Width: 700, Height: 400},
		Icon:          icon,
		Layout:        dcl.VBox{Spacing: 10, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		Children: []dcl.Widget{
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
				Border: false,
				Children: []dcl.Widget{
					dcl.Label{
						Text: "SSCC Префикс:",
					},
					dcl.TextEdit{
						AssignTo: &ssccPrefix,
						Text:     model.PrefixSSCC,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
				Border: false,
				Children: []dcl.Widget{
					dcl.Label{
						Text: "В коробке:",
					},
					dcl.NumberEdit{
						AssignTo: &perPallet,
						Value:    0,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
				Border: false,
				Children: []dcl.Widget{
					dcl.Label{
						Text: "Начальный № палеты :",
					},
					dcl.NumberEdit{
						AssignTo: &startPallet,
						Value:    0,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Border: false,
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: true, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 0}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo: &acceptPB,
						Text:     "Сохранить",
						OnClicked: func() {
							model := reductor.Instance().Model("")
							model.PerPallet = int(perPallet.Value())
							model.StartNumberSSCC = int(startPallet.Value())
							model.PrefixSSCC = ssccPrefix.Text()
							reductor.Instance().SetModel("", model)
							model.Sync(g)
							dlg.Accept()
						},
					},
					dcl.PushButton{
						AssignTo: &cancelPB,
						Text:     "Cancel",
						OnClicked: func() {
							dlg.Cancel()
						},
					},
					dcl.HSpacer{},
				},
			},
			dcl.VSpacer{},
		},
	}).Create(nil); err != nil {
		return -1, fmt.Errorf("%w", err)
	}
	rect := g.dlg.Bounds()
	dlg.SetBounds(walk.Rectangle{
		X: rect.X + 50,
		Y: rect.Y + 50,
		// Width: 400,
		// Height: 200,
	})
	perPallet.SetValue(float64(model.PerPallet))
	startPallet.SetValue(float64(model.StartNumberSSCC))
	ssccPrefix.SetText(model.PrefixSSCC)
	ret := dlg.Run()
	g.Logger().Debugf("dialog options return %d", ret)
	return out, nil
}
