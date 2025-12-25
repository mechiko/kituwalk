package gui

import (
	"fmt"
	"kituwalk/process"
	"kituwalk/process/protocol"
	"kituwalk/reductor"
	"kituwalk/utility"
	"os"
	"path/filepath"
	"time"
)

func (g *gui) generate() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	model := reductor.Instance().Model("")
	if model.Order == 0 {
		return fmt.Errorf("заказ не может быть 0")
	}

	if model.Order <= 0 {
		return fmt.Errorf("номер заказа должен быть больше 0")
	}
	if model.PerPallet <= 0 {
		return fmt.Errorf("количество единиц в упаковке должно быть больше 0")
	}
	if model.StartNumberSSCC <= 0 {
		return fmt.Errorf("начальный номер упаковки должен быть больше 0")
	}

	prc, err := process.New(g, g.dbs)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer prc.CloseZnakDB()

	if err := prc.ReadOrder(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if len(prc.Cis) == 0 {
		return fmt.Errorf("выбрано 0 марок. агрегация невозможна")
	}
	logMessage(fmt.Sprintf("отобрано %d КМ", len(prc.Cis)), g.textRich)
	if err := prc.GeneratePalletOrder(); err != nil {
		return fmt.Errorf("%w", err)
	}
	logMessage(fmt.Sprintf("сгенерировано %d палет по %d шт.", len(prc.Pallet), model.PerPallet), g.textRich)

	if err := prc.WritePaletsForce(); err != nil {
		return fmt.Errorf("%w", err)
	}
	logMessage("агрегация записано в АлкоХелп, обновите значком Поиск вид Агрегация в программе АлкоХелп", g.textRich)

	model = reductor.Instance().Model("")
	// запишем значения следующей SSCC в конфиг
	model.StartNumberSSCC = model.LastSSCC + 1
	if err := model.Sync(g); err != nil {
		return fmt.Errorf("%w", err)
	}
	model.Date = time.Now().Format("2006-01-02")
	// запоминаем модель уже когда палеты внесены в бд
	reductor.Instance().SetModel("", model)

	// если ошибка формирования протокола, то обращаться ко мне...
	if fileTxt, err := protocol.PrintKrinicaProtocol(prc); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		fileName := fmt.Sprintf("Агрегация_Заказ_%d_%s_%s.html", model.Order, time.Now().Format("2006-01-02"), utility.String(6))
		dir := g.Output()
		if !utility.PathOrFileExists(dir) {
			dir = "."
		}
		url := filepath.Join(dir, fileName)
		// if err := os.MkdirAll(dir, 0o755); err != nil {
		// 	return fmt.Errorf("%w", err)
		// }
		if err := os.WriteFile(url, fileTxt, 0o644); err != nil {
			return fmt.Errorf("%w", err)
		}
		if err := OpenFile(url); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}
