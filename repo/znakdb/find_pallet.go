package znakdb

// возвращает map[string]interface{}
// возвращает ошибку когда не найдено и когда реально ошибка
// если что, такая ошибка может быть errors.Is(err, db.ErrNoMoreRows)
func (z *DbZnak) FindPallet(number string) (pallet map[string]interface{}, err error) {
	sess := z.dbSession
	pallet = make(map[string]interface{})
	res := sess.Collection("order_mark_aggregation_codes").Find("unit_serial_number", number)
	if err := res.One(&pallet); err != nil {
		// if errors.Is(err, db.ErrNoMoreRows) {
		// }
		return pallet, err
	}
	return pallet, nil
}
