package ucexcel

type UseCaseExcel interface {
	Open() error
	// Save() error
	SaveSimple() error
	// BytesBuffer() (*bytes.Buffer, error)
	// ExcelFileName(prefix string) string
	// ExcelFileNameDownload(prefix string) string
	// // ToFile() error
	// UtilizedReportAll(report IZnakUtilizedReport) error
	// UtilizedReportGtin(report IZnakUtilizedReport) error
	UtilizationReportList(report [][]string) error
	UtilizationTxt(report []string) error
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
