package ucexcel

import (
	"fmt"
	"os"
	_ "time/tzdata"
)

func (ue *ucexcel) ToFile() error {
	fname := ue.ExcelFileName(ue.name)
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if _, err := ue.file.WriteTo(file); err != nil {
		file.Close()
		return fmt.Errorf("%w", err)
	}

	return file.Close()
}

func (ue *ucexcel) ToFileSimple() error {
	fname := ue.ExcelFileNameSimple(ue.name)
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err := ue.file.WriteTo(file); err != nil {
		file.Close()
		return fmt.Errorf("%w", err)
	}

	return file.Close()
}

func (ue *ucexcel) ToFileName(fn string) error {
	if fn == "" {
		return fmt.Errorf("пустое имя файла")
	}
	file, err := os.OpenFile(fn, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err := ue.file.WriteTo(file); err != nil {
		file.Close()
		return fmt.Errorf("%w", err)
	}

	return file.Close()
}
