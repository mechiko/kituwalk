package znakdb

type SerialNumbers struct {
	Id               int64  `db:"id"`
	IdOrderMarkCodes int    `db:"id_order_mark_codes"`
	Gtin             string `db:"gtin"`
	SerialNumber     string `db:"serial_number"`
	Code             string `db:"code"`
	BlockId          string `db:"block_id"`
	Status           string `db:"status"`
}

type SliceOrderSerialNumbers []*SerialNumbers

func (z *DbZnak) OrderSerialNumbers(number int64) (out SliceOrderSerialNumbers, err error) {
	out = make(SliceOrderSerialNumbers, 0)
	sess := z.Sess()
	col := sess.Collection("order_mark_codes_serial_numbers")
	res := col.Find("id_order_mark_codes", number)
	if err := res.All(&out); err != nil {
		return out, err
	}
	return out, nil
}

// находим КМ со статусом Нанесён
func (z *DbZnak) OrderSerialNumbersApply(number int64) (out SliceOrderSerialNumbers, err error) {
	out = make(SliceOrderSerialNumbers, 0)
	sess := z.Sess()
	col := sess.Collection("order_mark_codes_serial_numbers")
	res := col.Find("id_order_mark_codes = ?", number).And("status = ?", "Нанесён")
	if err := res.All(&out); err != nil {
		return out, err
	}
	return out, nil
}
