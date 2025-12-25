package ucexcel

import (
	"github.com/xuri/excelize/v2"
)

func (ue *ucexcel) Open() error {
	// создаем чистый файл без стилей
	ue.file = excelize.NewFile()
	return nil
}
