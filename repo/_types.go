package repo

import (
	"fmt"
	"kituwalk/dbscan"
	"sync"

	"go.uber.org/zap"
)

const modError = "pkg:repo"

var Version int64

type Repository struct {
	logger *zap.SugaredLogger
	dbs    *dbscan.Dbs
	mutex  sync.Mutex
}

func New(logger *zap.SugaredLogger, dbs *dbscan.Dbs) (rp *Repository, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("repo panic %v", r)
		}
	}()
	rp = &Repository{
		logger: logger,
		dbs:    dbs,
	}
	return rp, nil
}

// возвращаем DbInfo или nil
func (r *Repository) Info(t dbscan.DbInfoType) *dbscan.DbInfo {
	if di := r.dbs.Info(t); di != nil {
		return di
	}
	return nil
}

func (r *Repository) ListDb() (out []dbscan.DbInfo) {
	out = nil
	if r != nil && r.dbs != nil {
		out = r.dbs.List()
	}
	return out
}
