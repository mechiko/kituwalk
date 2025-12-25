package dbscan

import (
	"fmt"
	"path/filepath"
)

// name driver должны быть заполнены
// file если пусто будет вычислен из name добавлением .db
// проверяет возможность подключения
func ParseDbInfo(info *DbInfo) (dbi *DbInfo, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	filePathDir := ""
	if info.Driver == "" {
		return nil, fmt.Errorf("%s отсутствует имя драйвера", modError)
	}
	if info.Driver == "sqlite" {
		if info.File == "" {
			if info.Name == "" {
				return nil, fmt.Errorf("%s отсутствует имя базы данных для sqlite", modError)
			}
			info.File = info.Name + ".db"
		}
		filePathDir = filepath.Dir(info.File)
		if filePathDir == "." {
			filePathDir = ""
		}
		if filePathDir != "" {
			return nil, fmt.Errorf("%s имя файла базы данных для sqlite не должно содержать путь %s", modError, info.File)
		}
	}
	dbi = &DbInfo{
		Type:   info.Type,
		File:   info.File,
		Path:   info.Path,
		Driver: info.Driver,
		Name:   info.Name,
		Host:   info.Host,
		User:   info.User,
		Pass:   info.Pass,
		Exists: false,
	}
	if err := dbi.IsConnected(); err != nil {
		return nil, fmt.Errorf("ошибка %w", err)
	} else {
		dbi.Exists = true
	}

	return dbi, nil
}
