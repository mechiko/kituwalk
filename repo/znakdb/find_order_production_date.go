package znakdb

type Utilisation struct {
	Id               int64  `db:"id"`
	IdOrderMarkCodes int    `db:"id_order_mark_codes"`
	CreateDate       string `db:"create_date"`
	ProductionDate   string `db:"production_date"`
	ExpirationDate   string `db:"expiration_date"`
	UsageType        string `db:"usage_type"`
	Inn              string `db:"inn"`
	Kpp              string `db:"kpp"`
	Version          string `db:"version"`
	State            string `db:"state"`
	Status           string `db:"status"`
	ReportId         string `db:"report_id"`
	Archive          int    `db:"archive"`
	Json             string `db:"json"`
	Quantity         string `db:"quantity"`
	PrimaryDocNumber string `db:"primary_doc_number"`
	PrimaryDocDate   string `db:"primary_doc_date"`
	AlcVolume        string `db:"alc_volume"`
}

func (z *DbZnak) FindOrderProductionDate(number int64) (out string, err error) {
	sess := z.dbSession
	order := &Utilisation{}
	res := sess.Collection("order_mark_utilisation").Find("id_order_mark_codes = ? and status <> ?", number, "Отклонён")
	if err := res.One(order); err != nil {
		return "", err
	}

	return order.ProductionDate, nil
}
