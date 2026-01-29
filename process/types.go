package process

import (
	"fmt"
	"kituwalk/domain"
	"kituwalk/repo/configdb"
	"kituwalk/repo/znakdb"

	"kituwalk/utility"

	"kituwalk/dbscan"
)

// const startSSCC = "1462709225" // gs1 rus id zapivkom для памяти запивком

type Process struct {
	domain.Apper
	dbZnak   *znakdb.DbZnak
	inn      string
	Sscc     []string
	Cis      []*utility.CisInfo
	Pallet   map[string][]*utility.CisInfo
	warnings []string
	errors   []string
}

func New(app domain.Apper, dbs *dbscan.Dbs) (*Process, error) {
	cfgInfo := dbs.Info(dbscan.Config)
	if cfgInfo == nil {
		return nil, fmt.Errorf("бд конфиг.дб не найдена")
	}
	dbCfg, err := configdb.New(cfgInfo)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	inn, err := dbCfg.Key("inn")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	_ = dbCfg.Close()
	if inn == "" {
		return nil, fmt.Errorf("inn empty")
	}
	znakInfo := dbs.Info(dbscan.TrueZnak)
	if znakInfo == nil {
		return nil, fmt.Errorf("бд znak.дб не найдена")
	}
	dbZnak, err := znakdb.New(znakInfo)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	k := &Process{
		Apper:    app,
		dbZnak:   dbZnak,
		inn:      inn,
		Pallet:   make(map[string][]*utility.CisInfo),
		warnings: make([]string, 0),
		errors:   make([]string, 0),
		Sscc:     make([]string, 0),
		Cis:      make([]*utility.CisInfo, 0),
	}
	return k, nil
}

func (k *Process) AddWarn(warn string) {
	k.warnings = append(k.warnings, warn)
}

func (k *Process) Warnings() []string {
	return k.warnings
}

func (k *Process) AddError(err string) {
	k.errors = append(k.errors, err)
}

func (k *Process) Errors() []string {
	return k.errors
}

func (k *Process) ResetPalletMap() {
	for key := range k.Pallet {
		delete(k.Pallet, key)
	}
}

func (k *Process) Reset() {
	k.ResetPalletMap()
	k.Cis = make([]*utility.CisInfo, 0)
	k.Sscc = make([]string, 0)
	k.errors = make([]string, 0)
	k.warnings = make([]string, 0)
}

func (k *Process) CloseZnakDB() error {
	return k.dbZnak.Close()
}
