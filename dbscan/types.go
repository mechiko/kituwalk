package dbscan

import (
	"errors"
	"fmt"
	"strings"
)

const modError = "dbscan"

// type Apper interface {
// 	Logger() *zap.SugaredLogger
// 	Pwd() string
// 	ConfigPath() string
// 	DbPath() string
// }

type Dbs struct {
	// Apper
	path  string
	infos ListDbInfoForScan
}

// dbPath путь сканирования БД A3 и 4Z
// listInfo описатели для всех БД
// пустой путь перезаписывается на текущий "."
func New(listInfo ListDbInfoForScan, dbPath string) (d *Dbs, err error) {
	d = &Dbs{
		// Apper: app,
		infos: make(ListDbInfoForScan),
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.Join(err, fmt.Errorf("panic %s: %v", modError, r))
		}
	}()
	if dbPath != "" {
		if err := pathCheck(dbPath); err != nil {
			return nil, fmt.Errorf("dbPath %s %w", dbPath, err)
		}
	}
	// для поиска файлов надо точку
	if dbPath == "" {
		dbPath = "."
	}
	d.path = dbPath
	fsrarId := ""
	file4z := ""
	dbType := "sqlite"
	host := ""
	user := ""
	pass := ""
	// если Config есть в списке
	if config, ok := listInfo[Config]; ok && config != nil {
		if config.Path == "" {
			config.Path = dbPath
		}
		// only sqlite config.db
		config.Driver = "sqlite"
		config.File = "config.db"
		// проверяем структуру и пробуем коннект
		configParsedInfo, err := ParseDbInfo(config)
		if err != nil {
			return nil, fmt.Errorf("dbscan parse config info %w", err)
		}
		d.infos[Config] = configParsedInfo
		file4z, err = d.fromConfig(config, "oms_id")
		if err != nil {
			return nil, fmt.Errorf("dbscan fromConfig %w", err)
		}
		fsrarId, err = d.fromConfig(config, "fsrar_id")
		if err != nil {
			return nil, fmt.Errorf("dbscan fromConfig %w", err)
		}
		dbType, err = d.fromConfig(config, "db_type")
		if err != nil {
			return nil, fmt.Errorf("dbscan fromConfig %w", err)
		}
		dbType = strings.ToLower(dbType)
		host, err = d.fromConfig(config, "db_server")
		if err != nil {
			return nil, fmt.Errorf("dbscan fromConfig %w", err)
		}
		user, err = d.fromConfig(config, "db_user")
		if err != nil {
			return nil, fmt.Errorf("dbscan fromConfig %w", err)
		}
		pass, err = d.fromConfig(config, "db_pass")
		if err != nil {
			return nil, fmt.Errorf("dbscan fromConfig %w", err)
		}
	}
	if fsrarId == "" {
		fsrarId = findA3Name(dbPath)
	}
	if file4z == "" {
		file4z = find4zName(dbPath)
	}
	// Other: ignore parse/connect errors; save a defensive copy so callers can inspect it later
	if other, ok := listInfo[Other]; ok && other != nil {
		if other.Path == "" {
			other.Path = dbPath
		}
		otherParsedInfo, err := ParseDbInfo(other)
		if err != nil {
			// Defensive copy to avoid aliasing with the caller's map entry.
			cp := *other
			// Make the state explicit: not validated/connected.
			cp.Exists = false
			d.infos[Other] = &cp
		} else {
			d.infos[Other] = otherParsedInfo
		}
	}
	if a3, ok := listInfo[A3]; ok && a3 != nil {
		if a3.Path == "" {
			a3.Path = dbPath
		}
		if a3.Driver == "" {
			a3.Driver = dbType
		}
		if a3.Name == "" {
			a3.Name = fsrarId
		}
		if a3.Host == "" {
			a3.Host = host
		}
		if a3.User == "" {
			a3.User = user
		}
		if a3.Pass == "" {
			a3.Pass = pass
		}
		a3ParsedInfo, err := ParseDbInfo(a3)
		if err != nil {
			return nil, fmt.Errorf("dbscan parse a3 info %w", err)
		}
		d.infos[A3] = a3ParsedInfo
	}
	if trueZnak, ok := listInfo[TrueZnak]; ok && trueZnak != nil {
		if trueZnak.Path == "" {
			trueZnak.Path = dbPath
		}
		if trueZnak.Driver == "" {
			trueZnak.Driver = dbType
		}
		if trueZnak.Name == "" {
			trueZnak.Name = file4z
		}
		if trueZnak.Host == "" {
			trueZnak.Host = host
		}
		if trueZnak.User == "" {
			trueZnak.User = user
		}
		if trueZnak.Pass == "" {
			trueZnak.Pass = pass
		}
		trueParsedInfo, err := ParseDbInfo(trueZnak)
		if err != nil {
			return nil, fmt.Errorf("dbscan parse trueznak info %w", err)
		}
		d.infos[TrueZnak] = trueParsedInfo
	}
	return d, nil
}

func (d *Dbs) Info(t DbInfoType) *DbInfo {
	if d == nil {
		return nil
	}
	if dbi, ok := d.infos[t]; ok {
		return dbi
	}
	return nil
}

func (d *Dbs) List() []DbInfo {
	out := make([]DbInfo, 0)
	for _, i := range d.infos {
		info := *i
		out = append(out, info)
	}
	return out
}
