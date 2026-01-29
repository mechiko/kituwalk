package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"kituwalk/app"
	"kituwalk/config"
	"kituwalk/dbscan"
	"kituwalk/gui"
	"kituwalk/licenser"
	"kituwalk/reductor"
	"kituwalk/zaplog"

	"kituwalk/utility"

	"go.uber.org/zap"
)

// если local true то папка создается локально
var local = flag.Bool("local", false, "")

func ifErrorExit(loger *zap.SugaredLogger, title string, err error) {
	if err != nil {
		if loger != nil {
			loger.Errorf("%s %v", title, err)
		}
		utility.MessageBox(title, err.Error())
		os.Exit(-1)
	}
}

func main() {
	cfg, err := config.New("", !*local)
	ifErrorExit(nil, "ошибка конфигурации", err)

	var logsOutConfig = map[string][]string{
		"logger": {"stdout", filepath.Join(cfg.LogPath(), config.Name)},
	}
	zl, err := zaplog.New(logsOutConfig, true)
	ifErrorExit(nil, "ошибка создания логера", err)

	lg, err := zl.GetLogger("logger")
	ifErrorExit(nil, "ошибка получения логера", err)

	loger := lg.Sugar()
	loger.Debug("zaplog started")
	loger.Infof("mode = %s", config.Mode)
	if cfg.Warning() != "" {
		loger.Infof("pkg:config warning %s", cfg.Warning())
	}
	_, err = licenser.New(licenser.MAC, "")
	ifErrorExit(loger, "ошибка лицензии", err)

	// создаем приложение с опциями из конфига и логером основным
	app := app.New(cfg, loger, ".")
	err = app.CreatePath()
	ifErrorExit(loger, "ошибка создания папки приложения", err)
	// инициализируем REPO
	// TODO изменить получение путей из конфига
	listDbs := make(dbscan.ListDbInfoForScan)
	listDbs[dbscan.Config] = dbscan.NewDbInfo(dbscan.Config)
	listDbs[dbscan.TrueZnak] = dbscan.NewDbInfo(dbscan.TrueZnak)

	dbs, err := dbscan.New(listDbs, ".")
	ifErrorExit(loger, "Ошибки запуска сканирования баз данных", err)
	if dbs == nil {
		ifErrorExit(loger, "Ошибки запуска сканирования баз данных", fmt.Errorf("ошибка dbs nil"))
	}
	model := reductor.Model{}
	model.Read(app)
	if model.StartNumberSSCC < 0 {
		model.StartNumberSSCC = 0
		err := model.Sync(app)
		ifErrorExit(loger, "Ошибки записи файла конфигурации", err)
	}
	if model.PerPallet < 0 {
		model.PerPallet = 1
		err := model.Sync(app)
		ifErrorExit(loger, "Ошибки записи файла конфигурации", err)
	}

	// создаем редуктор с новой моделью
	reductor.New(model, app.Logger())
	guimain, err := gui.New(app, dbs)
	ifErrorExit(loger, "Ошибки запуска оболочки", err)
	_, _ = guimain.StartDialog()
}
