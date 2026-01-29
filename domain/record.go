package domain

import (
	"fmt"
	"kituwalk/utility"
	"strings"
	"time"
)

type Record struct {
	Cis      *utility.CisInfo
	Korob    string
	Palet    string
	Produced time.Time
	Expired  time.Time
	Order    int64
	Serial   string
}

type Pallete struct {
	Code   string
	Volume string
	Name   string
	Box    string // порядковый номер коробки по этому заказу только по шаблону Sprintf("%s_%03d", boxPrefix, box)
}

// вставляем GS в файлах от криницы, там он опущен...
// это только для пива!!!
func NewRecord(row []string) (*Record, error) {
	if len(row) != 5 {
		return nil, fmt.Errorf("записей не равно 5")
	}
	s := row[0]
	// Insert GS only when absent and length is sufficient
	if !strings.ContainsRune(s, rune(29)) {
		if len(s) < 26 {
			return nil, fmt.Errorf("некорректная длина КМ %d (<26)", len(s))
		}
		s = s[:25] + "\x1D" + s[25:]
	}
	cis, err := utility.ParseCisInfo(s)
	if err != nil {
		return nil, fmt.Errorf("получение КМ %w", err)
	}
	produced, err := parseDate(row[3])
	if err != nil {
		return nil, fmt.Errorf("ошибка даты производства %s%w", row[3], err)
	}
	expired, err := parseDate(row[4])
	if err != nil {
		return nil, fmt.Errorf("ошибка даты срока годности %s%w", row[4], err)
	}

	r := &Record{
		Cis:      cis,
		Korob:    row[1],
		Palet:    row[2],
		Produced: produced,
		Expired:  expired,
	}
	return r, nil
}

// полная строка 11 ячеек
func IsRecord(row []string) bool {
	return len(row) == 5 && strings.HasPrefix(row[0], "01")
}

func parseDate(s string) (time.Time, error) {
	layout := "02.01.2006" // Corresponds to DD.MM.YYYY
	parsedTime, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
