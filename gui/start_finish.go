package gui

import "kituwalk/reductor"

func (g *gui) start() {
	model := reductor.Instance().Model("")
	model.IsProcess = true
	reductor.Instance().SetModel("", model)
	g.dlg.Synchronize(func() {
		g.acceptPB.SetEnabled(false)
		g.cancelPB.SetEnabled(false)
	})
}

func (g *gui) finish() {
	modelFinal := reductor.Instance().Model("")
	modelFinal.IsProcess = false
	reductor.Instance().SetModel("", modelFinal)
	g.dlg.Synchronize(func() {
		g.order.SetValue(float64(modelFinal.Order))
		g.entirely.SetChecked(modelFinal.Entirely)
		g.statusKM.SetChecked(modelFinal.StatusKM)
		g.acceptPB.SetEnabled(true)
		g.cancelPB.SetEnabled(true)
	})
}
