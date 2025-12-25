package process

import (
	"fmt"
	"kituwalk/reductor"
	"kituwalk/repo/znakdb"

	"kituwalk/utility"
)

func (k *Process) ReadOrder() (err error) {
	model := reductor.Instance().Model("")
	if order, err := k.dbZnak.FindOrder(model.Order); err != nil {
		return err
	} else {
		model.Name = order.ProductName
		model.Gtin = order.Gtin
	}
	if production, err := k.dbZnak.FindOrderProductionDate(model.Order); err == nil {
		model.ProductionDate = production
	}
	var arr znakdb.SliceOrderSerialNumbers
	if model.StatusKM {
		if arr, err = k.dbZnak.OrderSerialNumbersApply(model.Order); err != nil {
			return err
		}
	} else {
		if arr, err = k.dbZnak.OrderSerialNumbers(model.Order); err != nil {
			return err
		}
	}
	// Build a local slice, then assign once
	cisParsed := make([]*utility.CisInfo, 0, len(arr))
	for _, cis := range arr {
		item, err := utility.ParseCisInfo(cis.Code)
		if err != nil {
			return fmt.Errorf("parse cis code:%s %w", cis.Code, err)
		}
		cisParsed = append(cisParsed, item)
	}
	k.Cis = cisParsed
	_ = reductor.Instance().SetModel("", model)
	return nil
}
