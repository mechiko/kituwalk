package process

import (
	"fmt"
	"kituwalk/reductor"
	"strings"
	"time"

	"kituwalk/utility"

	"github.com/upper/db/v4"
)

type Aggregation struct {
	Id                      int64  `db:"id,omitempty"`
	CreateDate              string `db:"create_date"`
	Inn                     string `db:"inn"`
	UnitSerialNumber        string `db:"unit_serial_number"`
	AggregationUnitCapacity string `db:"aggregation_unit_capacity"`
	AggregatedItemsCount    string `db:"aggregated_items_count"`
	AggregationType         string `db:"aggregation_type"`
	Gtin                    string `db:"gtin"`
	Note                    string `db:"note"`
	Version                 string `db:"version"`
	State                   string `db:"state"`
	Status                  string `db:"status"`
	OrderId                 string `db:"order_id"`
	Archive                 int    `db:"archive"`
	Json                    string `db:"json"`
}

type AggregationCode struct {
	Id                      int64  `db:"id,omitempty"`
	IdOrderMarkAggregation  int    `db:"id_order_mark_aggregation"`
	SerialNumber            string `db:"serial_number"`
	Code                    string `db:"code"`
	UnitSerialNumber        string `db:"unit_serial_number"`
	AggregationUnitCapacity string `db:"aggregation_unit_capacity"`
	AggregatedItemsCount    string `db:"aggregated_items_count"`
	Status                  string `db:"status"`
}

// opts map[inn] [order] [count]
// запись в обход функции модуля repo
func (k *Process) WritePaletsForce() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	model := reductor.Instance().Model("")

	sess := k.dbZnak.Sess()
	note := strings.Join(k.Sscc, " ")
	err = sess.Tx(func(tx db.Session) error {
		agr := &Aggregation{
			CreateDate:      time.Now().Local().Format("2006.01.02 15:04:05"),
			Inn:             model.Inn,
			OrderId:         fmt.Sprintf("%d", model.Order),
			Note:            note,
			Gtin:            model.Gtin,
			AggregationType: "Упаковка",
			Version:         "1",
			State:           "Создан",
			Status:          "Не проведён",
			Archive:         0,
		}
		if err := tx.Collection("order_mark_aggregation").InsertReturning(agr); err != nil {
			return err
		} else {
			for _, paletProducer := range k.Sscc {
				cis := k.Pallet[paletProducer]
				if err := k.writePaletForce(tx, agr, paletProducer, cis); err != nil {
					return err
				}
			}
			return nil
		}
	})
	return err
}

func (k *Process) writePaletForce(tx db.Session, agr *Aggregation, palet string, cises []*utility.CisInfo) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	for i := range cises {
		cis := &AggregationCode{
			IdOrderMarkAggregation:  int(agr.Id),
			SerialNumber:            cises[i].Serial,
			Code:                    cises[i].Cis,
			Status:                  "Импортирован",
			UnitSerialNumber:        palet,
			AggregationUnitCapacity: fmt.Sprintf("%d", len(cises)),
			AggregatedItemsCount:    fmt.Sprintf("%d", len(cises)),
		}
		if _, err := tx.Collection("order_mark_aggregation_codes").Insert(cis); err != nil {
			return err
		}
	}
	return nil
}
