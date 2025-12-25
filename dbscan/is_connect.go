package dbscan

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mssql"
	"github.com/upper/db/v4/adapter/sqlite"
)

// func (d *dbs) IsConnected(info *DbInfo) (err error) {
func (d *DbInfo) IsConnected() (err error) {
	var dbSess db.Session
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	switch d.Driver {
	case "sqlite":
		if d.File == "" {
			return fmt.Errorf("%s отсутствуют имя файла базы данных для sqlite", modError)
		}
		resultFilePath := filepath.Join(d.Path, d.File)
		if err := pathCheck(resultFilePath); err != nil {
			return fmt.Errorf("connect %w", err)
		}
		// если указан не файл а путь к каталогу
		if st, statErr := os.Stat(resultFilePath); statErr != nil || !st.Mode().IsRegular() {
			return fmt.Errorf("%s путь %s не является файлом sqlite", modError, resultFilePath)
		}
		uri := d.SqliteUri(resultFilePath)
		dbSess, err = sqlite.Open(uri)
		if err != nil {
			return fmt.Errorf("%s ошибка подключения %v", modError, err)
		}
		defer func() {
			if errSess := dbSess.Close(); errSess != nil {
				// Go 1.20+: joins parse error (if any) with close error
				err = errors.Join(err, fmt.Errorf("close session %w", errSess))
			}
		}()
	case "mssql":
		if d.Name == "" {
			return fmt.Errorf("%s отсутствуют имя базы данных для Other", modError)
		}
		uri := d.MssqlUri()
		dbSess, err = mssql.Open(uri)
		if err != nil {
			return fmt.Errorf("%s %s", modError, err.Error())
		}
		defer func() {
			if errSess := dbSess.Close(); errSess != nil {
				// Go 1.20+: joins parse error (if any) with close error
				err = errors.Join(err, fmt.Errorf("close session %w", errSess))
			}
		}()
	default:
		return fmt.Errorf("%s ошибка driver %v", modError, d.Driver)
	}
	err = dbSess.Ping()
	if err != nil {
		return fmt.Errorf("%s ошибка ping %v", modError, err)
	}
	// пинг успешен
	return nil
}
