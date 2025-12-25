package ucexcel

import (
	"fmt"
)

func (ue *ucexcel) templateStatFloatBold(pos float64) {
	if err := ue.file.SetCellFloat(ue.sheet, ue.address.Address(), pos, 3, 64); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), celStyleReportBoldRightNum); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}
func (ue *ucexcel) templateStatIntBold(pos int64) {
	if err := ue.file.SetCellInt(ue.sheet, ue.address.Address(), pos); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), celStyleReportBoldRightNum); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}

func (ue *ucexcel) templateStatFloat(pos float64) {
	if err := ue.file.SetCellFloat(ue.sheet, ue.address.Address(), pos, 3, 64); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), celStyleReportRightNum); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}

func (ue *ucexcel) templateStatInt(pos int64) {
	if err := ue.file.SetCellInt(ue.sheet, ue.address.Address(), pos); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), celStyleReportRightNum); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}
func (ue *ucexcel) templateStatIntCenter(pos int64) {
	if err := ue.file.SetCellInt(ue.sheet, ue.address.Address(), pos); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), celStyleReportCenter); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}

func (ue *ucexcel) templateStatFloatSumBold(pos float64) {
	if err := ue.file.SetCellFloat(ue.sheet, ue.address.Address(), pos, 3, 64); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), celStyleReportBoldRightSum); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}
func (ue *ucexcel) templateStatFloatSum(pos float64) {
	if err := ue.file.SetCellFloat(ue.sheet, ue.address.Address(), pos, 3, 64); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), celStyleReportRightSum); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}
func (ue *ucexcel) templateStatStr(pos string) {
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), pos); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), celStyleReport); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}
func (ue *ucexcel) templateStatStrJust(pos string) {
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), pos); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), celStyleReportJust); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}
func (ue *ucexcel) templateStatStrCenter(pos string) {
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), pos); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), celStyleReportCenter); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}
func (ue *ucexcel) templateStatStrRight(pos string) {
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), pos); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), celStyleReportRight); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}

func (ue *ucexcel) templateFloat(pos float64, style int) {
	if err := ue.file.SetCellFloat(ue.sheet, ue.address.Address(), pos, 3, 64); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), style); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}
func (ue *ucexcel) templateInt(pos int64, style int) {
	if err := ue.file.SetCellInt(ue.sheet, ue.address.Address(), pos); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), style); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}
func (ue *ucexcel) templateStr(pos string, style int) {
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), pos); err != nil {
		fmt.Printf("SetCellStr err %s\n", err.Error())
	}
	if err := ue.file.SetCellStyle(ue.sheet, ue.address.Address(), ue.address.Address(), style); err != nil {
		fmt.Printf("SetCellStyle err %s\n", err.Error())
	}
	ue.address.NextCol()
}
