package dbscan

import (
	"github.com/upper/db/v4/adapter/mssql"
	"github.com/upper/db/v4/adapter/sqlite"
)

// host
// user
// password
// name имя файла без пути
// file полное имя файла и путь
// name имя БД
// driver драйвер
// Exists файл существует и только
type DbInfo struct {
	Type       DbInfoType
	Host       string
	User       string
	Pass       string
	File       string
	Name       string
	Driver     string
	Connection string
	Exists     bool // только для sqlite делает поиск файла
	Path       string
}

type ListDbInfoForScan map[DbInfoType]*DbInfo

func IsValidDbInfoType(s string) bool {
	switch DbInfoType(s) {
	case Config, A3, TrueZnak, Other:
		return true
	}
	return false
}

func NewDbInfo(t DbInfoType) *DbInfo {
	return &DbInfo{Type: t}
}

func (d ListDbInfoForScan) Info(t DbInfoType) *DbInfo {
	if dbi, ok := d[t]; ok {
		return dbi
	}
	return nil
}

type DbInfoType string

// типы зарезервированных БД
const (
	Config   DbInfoType = "config"
	A3       DbInfoType = "a3"
	TrueZnak DbInfoType = "trueznak"
	Other    DbInfoType = "other"
)

func (d *DbInfo) MssqlUri() *mssql.ConnectionURL {
	uri := &mssql.ConnectionURL{
		User:     d.User,
		Password: d.Pass,
		Host:     d.Host,
		Database: d.Name,
		Options: map[string]string{
			"encrypt": "disable",
		},
	}
	if uri.Host == "" {
		uri.Host = "127.0.0.1:1433"
	}
	return uri
}

func (d *DbInfo) SqliteUri(file string) *sqlite.ConnectionURL {
	uri := &sqlite.ConnectionURL{
		Database: file,
		Options: map[string]string{
			"mode":          "rw",
			"_journal_mode": "DELETE",
		},
	}
	return uri
}
