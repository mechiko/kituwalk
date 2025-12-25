package reductor

import (
	"errors"
	"fmt"
	"kituwalk/domain"
)

type Model struct {
	StartNumberSSCC int
	PerPallet       int
	Inn             string
	PrefixSSCC      string
	Name            string
	Gtin            string
	LastSSCC        int
	Date            string
	ProductionDate  string
	Order           int64
	Entirely        bool
	StatusKM        bool
	IsProcess       bool
}

type ModelList map[string]Model

// формат сообщения
type Message struct {
	Sender string
	Page   string
	Model  Model
}

func (m *Model) Sync(cfg domain.Apper) error {
	var agg error
	if err := cfg.SaveOptions("entirely", m.Entirely); err != nil {
		agg = errors.Join(agg, err)
	}
	if err := cfg.SaveOptions("statuskm", m.StatusKM); err != nil {
		agg = errors.Join(agg, err)
	}
	// if err := cfg.SaveOptions("inn", m.Inn); err != nil { agg = errors.Join(agg, err) }
	if err := cfg.SaveOptions("ssccprefix", m.PrefixSSCC); err != nil {
		agg = errors.Join(agg, err)
	}
	if err := cfg.SaveOptions("ssccstartnumber", m.StartNumberSSCC); err != nil {
		agg = errors.Join(agg, err)
	}
	if err := cfg.SaveOptions("perpallet", m.PerPallet); err != nil {
		agg = errors.Join(agg, err)
	}
	if agg != nil {
		return fmt.Errorf("save options: %w", agg)
	}
	return cfg.SaveAllOptions()
}

func (m *Model) Read(cfg domain.Apper) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	m.Entirely = cfg.Options().Entirely
	m.StatusKM = cfg.Options().StatusKM
	// m.Inn = inn
	m.PrefixSSCC = cfg.Options().SsccPrefix
	m.StartNumberSSCC = cfg.Options().SsccStartNumber
	m.PerPallet = cfg.Options().PerPallet
	return nil
}
