package gui

import (
	"fmt"
	"kituwalk/dbscan"
	"kituwalk/domain"

	"github.com/mechiko/walk"
)

type gui struct {
	domain.Apper
	dbs                *dbscan.Dbs
	dlg                *walk.Dialog
	page               *walk.Composite
	acceptPB, cancelPB *walk.PushButton
	selectFile         *walk.PushButton
	textRich           *RichEdit
	fileName           *walk.Label
	perPallet          *walk.NumberEdit
	startPallet        *walk.NumberEdit
	order              *walk.NumberEdit
	ssccPrefix         *walk.TextEdit
	entirely           *walk.CheckBox
	statusKM           *walk.CheckBox
}

func New(app domain.Apper, dbs *dbscan.Dbs) (*gui, error) {
	if dbs == nil {
		return nil, fmt.Errorf("dbs is nil")
	}
	g := &gui{
		Apper: app,
		dbs:   dbs,
	}
	return g, nil
}
