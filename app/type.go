package app

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"kituwalk/config"
	"kituwalk/domain"
	"kituwalk/utility"
	"os"
	"path/filepath"

	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// type IApp interface {
// 	Options() *config.Configuration
// 	SaveOptions(string, any) error
// 	Logger() *zap.SugaredLogger
// }

type app struct {
	ctx           context.Context
	uuid          string // идентификатор для уникальности формы
	config        *config.Config
	options       *config.Configuration // копия config.Configuration
	loger         *zap.SugaredLogger
	pwd           string
	startTime     time.Time
	endTime       time.Time
	output        string
	dbSelfPath    string
	defaultDbPath string
}

var _ domain.Apper = (*app)(nil)

// const modError = "app"

func New(cfg *config.Config, logger *zap.SugaredLogger, pwd string) *app {
	newApp := &app{}
	newApp.pwd = pwd
	newApp.loger = logger
	newApp.config = cfg
	newApp.options = cfg.Configuration()
	newApp.uuid = uuid.New().String()
	newApp.initDateMn()
	return newApp
}

func (a *app) initDateMn() {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		a.loger.Warnf("failed to load timezone, using UTC: %v", err)
		loc = time.UTC
	}
	t := time.Now().In(loc)
	year, month, _ := t.Date()
	a.startTime = time.Date(year, month, 1, 0, 0, 0, 0, loc)
	a.endTime = time.Date(year, month+1, 1, 0, 0, 0, 0, loc).Add(-time.Nanosecond)
}

func (a *app) NowDateString() string {
	n := time.Now()
	return fmt.Sprintf("%4d.%02d.%02d %02d:%02d:%02d", n.Local().Year(), n.Local().Month(), n.Local().Day(), n.Local().Hour(), n.Local().Minute(), n.Local().Second())
}

func (a *app) StartDateString() string {
	return fmt.Sprintf("%4d.%02d.%02d", a.startTime.Local().Year(), a.startTime.Local().Month(), a.startTime.Local().Day())
}

func (a *app) EndDateString() string {
	return fmt.Sprintf("%4d.%02d.%02d", a.endTime.Local().Year(), a.endTime.Local().Month(), a.endTime.Local().Day())
}

func (a *app) SetStartDate(d time.Time) {
	a.startTime = d
}

func (a *app) SetEndDate(d time.Time) {
	a.endTime = d
}

func (a *app) StartDate() time.Time {
	return a.startTime
}

func (a *app) EndDate() time.Time {
	return a.endTime
}

func (a *app) SetFsrarID(id string) {
	a.Config().SetInConfig("application.fsrarid", id)
}

func (a *app) Pwd() string {
	return a.pwd
}

func (a *app) Output() string {
	return a.output
}

func (a *app) Config() *config.Config {
	return a.config
}

func (a *app) Logger() *zap.SugaredLogger {
	return a.loger
}

func (a *app) Ctx() context.Context {
	return a.ctx
}

// выдаем адрес структуры опций программы чтобы править по месту
func (a *app) Options() *config.Configuration {
	return a.options
}

// записываем ключ и его значение только в пакет config
// изменения не записываются в файл конфигурации
func (a *app) SaveOptions(key string, value any) error {
	a.config.SetInConfig(key, value)
	if err := a.config.Save(); err != nil {
		return fmt.Errorf("save in config error %w", err)
	}
	return nil
}

// изменения записываются в файл конфигурации
func (a *app) SaveAllOptions() error {
	if err := a.config.Save(); err != nil {
		return fmt.Errorf("save all in config error %w", err)
	}
	return nil
}

// создаем по необходимости пути программы
func (a *app) CreatePath() error {
	// создаем папку вывода если не пустое значение
	// в папке запуска программы только или если она задана абсолютным значением пути
	dir := filepath.Join(utility.UserHomeDir(), "Documents")
	if utility.PathOrFileExists(dir) {
		nevaKod := filepath.Join(dir, "НеваКОД")
		if !utility.PathOrFileExists(nevaKod) {
			if err := pathCreate(nevaKod); err != nil {
				a.output = "."
				return fmt.Errorf("%w", err)
			}
		}
		a.output = nevaKod
		return nil
	}
	a.output = "."
	return nil
}

// создаем путь в каталоге программы или home
func createPath(path string, home string) (string, error) {
	fullPath := filepath.Join(home, path)
	if filepath.IsAbs(path) {
		fullPath = path
	}
	if err := pathCreate(fullPath); err != nil && !errors.Is(err, fs.ErrExist) {
		return "", fmt.Errorf("cannot create path %s: %w", fullPath, err)
	}
	return filepath.Abs(fullPath)
}

func pathCreate(path string) error {
	if path != "" {
		if err := os.MkdirAll(path, 0o755); err != nil {
			return err
		}
	}
	return nil
}

func (a *app) ConfigPath() string {
	if a.config != nil {
		return a.config.ConfigPath()
	}
	return ""
}

func (a *app) DefaultDbPath() string {
	return a.defaultDbPath
}

func (a *app) SetDefaultDbPath(path string) {
	a.defaultDbPath = path
}

func (a *app) LogPath() string {
	if a.config != nil {
		return a.config.LogPath()
	}
	return ""
}

func (a *app) BaseUrl() string {
	host := a.options.Hostname
	port := a.options.HostPort
	if port == "" || port == "auto" {
		return fmt.Sprintf("http://%s", host)
	}
	return fmt.Sprintf("http://%s:%s", host, port)
}

func (a *app) DbSelfPath() string {
	return a.dbSelfPath
}

func (a *app) SetDbSelfPath(path string) {
	a.dbSelfPath = path
}
