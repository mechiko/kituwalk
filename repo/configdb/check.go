package configdb

import (
	"fmt"
)

// вызывается каждый раз при создании объекта
func (r *DbConfig) Check() (err error) {
	if r == nil {
		return fmt.Errorf("%s check: receiver is nil", modError)
	}
	if r.dbInfo == nil {
		return fmt.Errorf("%s dbInfo is nil", modError)
	}
	if !r.dbInfo.Exists {
		return fmt.Errorf("dbConfig dbInfo.Exists false")
	}
	r.dbSession, err = r.dbInfo.Connect()
	if err != nil {
		return fmt.Errorf("%s check ошибка подключения к БД %w", modError, err)
	}
	return nil
}
