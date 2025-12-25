package address

import (
	"strconv"
)

type UseCaseExcel interface {
	Open() error
	Save() error
	SaveSimple() error
	// BytesBuffer() (*bytes.Buffer, error)
	// ExcelFileName(prefix string) string
	// ExcelFileNameDownload(prefix string) string
	// // ToFile() error
	// UtilizedReportAll(report IZnakUtilizedReport) error
	// UtilizedReportGtin(report IZnakUtilizedReport) error
	// UtilizationReportList(report []string) error
}

type ExcelAddress interface {
	Address() string
	NextCol() string
	NextRow() string
	MoveTo(row int, col int) string
	ShiftCol(col int) string
	Range(row int, col int) string
	AddCol(int) string
	AddRow(int) string
	Row() int
	Col() int
}

type address struct {
	row int
	col int
}

var _ ExcelAddress = (*address)(nil)

func New(row int, col int) *address {
	return &address{
		row: row,
		col: col,
	}
}

func (a *address) Address() string {
	colChar := string(a.col + 65)
	out := colChar + strconv.Itoa(a.row)
	return out
}

func (a *address) NextCol() string {
	a.col += 1
	out := a.Address()
	return out
}

func (a *address) NextRow() string {
	a.row += 1
	a.col = 0
	out := a.Address()
	return out
}

func (a *address) MoveTo(row int, col int) string {
	a.row = row
	a.col = col
	out := a.Address()
	return out
}
func (a *address) AddCol(col int) string {
	a.col += col
	out := a.Address()
	return out
}
func (a *address) AddRow(row int) string {
	a.row += row
	out := a.Address()
	return out
}

func (a *address) ShiftCol(col int) string {
	colChar := string(a.col + col + 65)
	out := colChar + strconv.Itoa(a.row)
	return out
}

func (a *address) Range(row int, col int) string {
	colChar := string(a.col + col + 65)
	out := colChar + strconv.Itoa(a.row+row)
	return out
}

func (a *address) Row() int {
	return a.row
}

func (a *address) Col() int {
	return a.col
}
