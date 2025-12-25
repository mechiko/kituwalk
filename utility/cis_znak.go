package utility

import (
	"fmt"
	"strings"
)

const (
	serialStartPos = 19
	gtinStartPos   = 2
	gtinEndPos     = 16
	gtinMinLength  = gtinEndPos - 1 // 15
)

// Вот подробное описание структуры кода маркировки:
// GTIN (Global Trade Item Number):
// 14 цифр, уникальный код товара, идентифицирующий конкретный вид продукции.
// Серийный номер:
// Уникальный идентификатор, присваиваемый каждой конкретной единице товара, может содержать буквы и цифры.
// Ключ проверки:
// Используется для проверки подлинности кода.
// Идентификатор применения (AI):
// Двухзначные коды, такие как 01, 21, 91, 92, 93, которые указывают на тип данных, содержащихся в следующей группе символов. Например, AI=01 указывает на GTIN, AI=21 на серийный номер.
// Пример структуры кода для обуви:
// 01 (GTIN) + 14 цифр (код товара) + 21 (серийный номер) + 13 символов (серийный номер) + 91 (идентификатор ключа проверки) + 4 символа (ключ проверки) + 92 + 88 символов (код проверки).
// Пример структуры кода для пивной продукции:
// 01 (GTIN) + 14 цифр (код товара) + 21 (серийный номер) + 7 символов (серийный номер) + 93 (идентификатор кода проверки) + 4 символа (код проверки).

// KМ полный
// 0    2               16   19      25 26
// [01][04810014011833][215][000001][0x1d][93][cRX2]
// [A2][B14           ][C3 ][D6    ]   [E6     ]
// В - GTIN
// D - Serial огрнаничивается 25 байтом (код 29 в таблице символов ASCII)
// E - crypto 93 плюс 4 символа кода проверки
// С - идентификатор применения (21) и идентификатор государства, в котором этот код был эмитирован (5 - Российская Федерация)

// структура КМ для А3 ЧЗ
type CisInfo struct {
	Cis    string // полный код без криптохвоста
	Code   string // полный код КМ выдаваемый ЧЗ
	Serial string // серийный номер КМ без префикса страны (без первого символа)
	Gtin   string // GTIN
}

// парсим полный код без FNC1 в поля
// марки без GS но с криптохвостом!
// 0108934627096868215fN_nfd934Plc
// отрезаем хвост 6 символов
const lengthCrypto = 6

func ParseCisInfo(code string) (*CisInfo, error) {
	i := &CisInfo{
		Code: code,
	}
	index := strings.IndexByte(code, '\x1D')
	if index > 0 {
		i.Cis = code[:index]
	} else {
		i.Cis = code[:len(code)-lengthCrypto]
		return nil, fmt.Errorf("марка без GS символа не пригодна для печати")
	}
	err := i.parseSerial()
	if err != nil {
		return nil, err
	}
	err = i.parseGtin()
	return i, err
}

func (i *CisInfo) parseSerial() error {
	if len(i.Cis) > serialStartPos {
		i.Serial = i.Cis[serialStartPos:]
		return nil
	} else {
		return fmt.Errorf("код КМ не полный для серийного номера")
	}
}

func (i *CisInfo) parseGtin() error {
	if len(i.Cis) > gtinMinLength {
		i.Gtin = i.Cis[gtinStartPos:gtinEndPos]
		return nil
	} else {
		return fmt.Errorf("код КМ не полный для GTIN")
	}
}

func (i *CisInfo) FNC1() string {
	return fmt.Sprintf("\xe8%s", i.Code)
}
