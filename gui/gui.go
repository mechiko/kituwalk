package gui

import (
	"fmt"
	"kituwalk/reductor"
	"kituwalk/utility"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (g *gui) StartDialog() (out string, err error) {
	icon, err := walk.Resources.Icon("3")
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	model := reductor.Instance().Model("")
	if err := (dcl.Dialog{
		AssignTo:      &g.dlg,
		Title:         "Агрегация для AlcoHelp 3",
		Size:          dcl.Size{Width: 700, Height: 400},
		Icon:          icon,
		Layout:        dcl.VBox{Spacing: 10, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
		DefaultButton: &g.acceptPB,
		CancelButton:  &g.cancelPB,
		Children: []dcl.Widget{
			dcl.Composite{
				Layout:   dcl.VBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
				Border:   true,
				AssignTo: &g.page,
				MinSize:  dcl.Size{Height: 200},
				Children: []dcl.Widget{},
			},
			dcl.Composite{
				Layout:  dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
				Border:  false,
				MinSize: dcl.Size{Width: 500},
				Children: []dcl.Widget{
					dcl.Label{
						Text: "Заказ :",
					},
					dcl.NumberEdit{
						AssignTo: &g.order,
						Value:    model.Order,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout:  dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
				Border:  false,
				MinSize: dcl.Size{Width: 500},
				Children: []dcl.Widget{
					dcl.CheckBox{
						AssignTo: &g.entirely,
						Checked:  model.Entirely,
						Text:     "агрегировать целиком в одну упаковку",
					},
					// dcl.Label{
					// 	Text: "Агрегировать целиком в одну упаковку :",
					// },
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout:  dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
				Border:  false,
				MinSize: dcl.Size{Width: 500},
				Children: []dcl.Widget{
					dcl.CheckBox{
						AssignTo: &g.statusKM,
						Checked:  model.StatusKM,
						Text:     "проверять статус КМ",
					},
					// dcl.Label{
					// 	Text: "Агрегировать целиком в одну упаковку :",
					// },
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Border: false,
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: true, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 0}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo: &g.acceptPB,
						Text:     "Агрегировать",
						OnClicked: func() {
							model := reductor.Instance().Model("")
							model.Order = int64(g.order.Value())
							model.Entirely = g.entirely.Checked()
							model.StatusKM = g.statusKM.Checked()
							reductor.Instance().SetModel("", model)
							model.Sync(g)
							g.start()
							go func() {
								if err := g.generate(); err != nil {
									logErrMessage(err.Error(), g.textRich)
								}
								g.finish()
							}()
							// dlg.Accept()
						},
					},
					dcl.PushButton{
						AssignTo: &g.cancelPB,
						Text:     "Настройка",
						OnClicked: func() {
							// ret = 0 кнопка ОК
							if _, err := g.optionsDialog(); err != nil {
								g.Logger().Errorf("%w", err)
							}
						},
					},
					dcl.PushButton{
						AssignTo: &g.cancelPB,
						Text:     "Выход",
						OnClicked: func() {
							model := reductor.Instance().Model("")
							if model.IsProcess {
								utility.MessageBox("ошибка", "запущен процесс обработки. дождитесь его завершения")
								return
							}
							g.dlg.Cancel()
						},
					},
					dcl.HSpacer{},
				},
			},
			dcl.VSpacer{},
		},
	}).Create(nil); err != nil {
		return "", fmt.Errorf("%w", err)
	}
	g.dlg.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		model := reductor.Instance().Model("")
		if model.IsProcess {
			*canceled = true
			utility.MessageBox("ошибка", "запущен процесс обработки. дождитесь его завершения")
			return
		}
		*canceled = false
	})
	g.textRich, err = NewRichEdit(g.page)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	if g.textRich == nil {
		return "", fmt.Errorf("nil richtext")
	}
	g.dlg.SetBounds(walk.Rectangle{
		X:     300,
		Y:     300,
		Width: 400,
		// Height: 200,
	})
	g.order.SetValue(float64(model.Order))
	g.entirely.SetChecked(model.Entirely)
	g.statusKM.SetChecked(model.StatusKM)
	logMessage("введите ид заказа в поле", g.textRich)
	logMessage("определитесь целиком или иначе агрегировать", g.textRich)
	logMessage("задайте проверять ли статус КМ", g.textRich)
	logMessage("для ПИВА требуется статус Нанесен", g.textRich)
	logMessage("для остальных требуется статус Введен в оборот", g.textRich)
	logMessage("программа может проверить только статус Нанесен!!!", g.textRich)
	if ret := g.dlg.Run(); ret != 1 {
		return "", fmt.Errorf("dialog return %d", ret)
	}
	return out, nil
}
