package ucexcel

import (
	_ "time/tzdata"
)

func (ue *ucexcel) Save() error {
	return ue.ToFile()
}

func (ue *ucexcel) SaveSimple() error {
	return ue.ToFileSimple()
}
