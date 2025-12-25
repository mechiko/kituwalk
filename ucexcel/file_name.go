package ucexcel

import (
	"fmt"
	"kituwalk/utility"
	"path/filepath"
	"time"
	_ "time/tzdata"
)

const ext = "xlsx"

func (ue *ucexcel) ExcelFileName(prefix string) string {
	ds := time.Now()
	startDate := fmt.Sprintf("%02d.%02d.%4d", ds.Local().Day(), ds.Local().Month(), ds.Local().Year())
	endDate := fmt.Sprintf("%02d.%02d.%4d", ds.Local().Day(), ds.Local().Month(), ds.Local().Year())
	guid := utility.String(8)
	name := prefix + "_" + startDate + "_" + endDate + "_" + guid
	name += "." + ext
	return filepath.Join("output", name)
}

func (ue *ucexcel) ExcelFileNameDownload(prefix string) string {
	ds := time.Now()
	startDate := fmt.Sprintf("%02d.%02d.%4d", ds.Local().Day(), ds.Local().Month(), ds.Local().Year())
	endDate := fmt.Sprintf("%02d.%02d.%4d", ds.Local().Day(), ds.Local().Month(), ds.Local().Year())
	name := prefix + "_" + startDate + "_" + endDate
	name += "." + ext
	return name
}

func (ue *ucexcel) ExcelFileNameSimple(prefix string) string {
	name := prefix + "." + ext
	return name
}
