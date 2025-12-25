package ucexcel

import (
	"fmt"
	"kituwalk/domain"
	"kituwalk/process"
	"kituwalk/ucexcel/address"
)

func (ue *ucexcel) ReportList(report *process.Process) error {
	ue.sheet = "Sheet1"
	// countRow = 1
	ue.address = address.New(1, 0)
	for _, palet := range report.PalletOrder {
		arrRec := report.Pallet[palet]
		if err := ue.templatePalet(palet, arrRec); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (ue *ucexcel) templatePalet(pallet string, s []*domain.Record) error {
	for _, rec := range s {
		if err := ue.templateLine(pallet, rec); err != nil {
			return fmt.Errorf("excel error %w", err)
		}
	}
	return nil
}

func (ue *ucexcel) templateLine(palet string, s *domain.Record) error {
	addr := ue.address.Address()

	if err := ue.file.SetCellStr(ue.sheet, addr, s.Cis.Code); err != nil {
		return fmt.Errorf("excel error %w", err)
	}

	if err := ue.file.SetCellStr(ue.sheet, ue.address.NextCol(), s.Serial); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	if err := ue.file.SetCellStr(ue.sheet, ue.address.NextCol(), s.Box); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	if err := ue.file.SetCellStr(ue.sheet, ue.address.NextCol(), s.Gtin); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	if err := ue.file.SetCellStr(ue.sheet, ue.address.NextCol(), s.Name); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	if err := ue.file.SetCellStr(ue.sheet, ue.address.NextCol(), palet); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextRow()
	return nil
}
