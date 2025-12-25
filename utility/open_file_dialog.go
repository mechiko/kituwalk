package utility

import (
	"errors"
	"fmt"
	"os"
)

type FileType string

const (
	Excel FileType = "xlsx"
	Txt   FileType = "txt"
	Pdf   FileType = "pdf"
	Csv   FileType = "csv"
	DB    FileType = "db"
	All   FileType = "all"
)

func DialogOpenFile(fileType []FileType, name string, wd string) (string, error) {
	if wd == "." {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		wd = dir
	}
	dlg := File().SetStartDir(wd).SetStartFile(name)
	for _, t := range fileType {
		switch t {
		case Excel:
			dlg.Filter("Excel", "xlsx")
		case Csv:
			dlg.Filter("CSV", "csv")
		case Txt:
			dlg.Filter("TXT", "txt")
		case Pdf:
			dlg.Filter("PDF", "pdf")
		case DB:
			dlg.Filter("DB", "db")
		case All:
			dlg.Filter("ALL", "*")
		}
	}
	result, err := dlg.Load()
	if errors.Is(err, ErrCancelled) {
		return "", fmt.Errorf("диалог выбора отменен")
	}
	if err != nil {
		return "", err
	}
	return result, nil
}

func DialogSaveFile(fileType FileType, name string, wd string) (string, error) {
	if wd == "." {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		wd = dir
	}
	dlg := File().SetStartDir(wd).SetStartFile(name)
	switch fileType {
	case Excel:
		dlg.Filter("Excel", "xlsx")
	case Csv:
		dlg.Filter("CSV", "csv")
	case Txt:
		dlg.Filter("Txt", "txt")
	case DB:
		dlg.Filter("Database", "db")
	case All:
		dlg.Filter("All", "*")
	}
	result, err := dlg.Save()
	if errors.Is(err, ErrCancelled) {
		return "", fmt.Errorf("диалог выбора отменен")
	}
	if err != nil {
		return "", err
	}

	return result, nil
}

func MessageBox(title, msg string) {
	if msg == "" {
		return
	}
	Message("%s", msg).Title(title).Info()
}

func DialogSelectDir(wd string) (string, error) {
	if wd == "." {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		wd = dir
	}
	directory, err := Directory().SetStartDir(wd).Title("Выберите папку с файлами данных").Browse()
	if err != nil {
		return "", err
	}
	if directory == "Cancelled" {
		return "", errors.New("прерван диалог выбора каталога")
	}
	return directory, nil
}
