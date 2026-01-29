package reductor

import (
	"sync"

	"go.uber.org/zap"
)

type Reductor struct {
	mutex      sync.Mutex
	logger     *zap.SugaredLogger
	modelPages ModelList
}

var once sync.Once
var instance *Reductor

// по умолчанию страница пустая строка
// дальше придумаю как добавлять страницы методом
func New(model Model, logger *zap.SugaredLogger) *Reductor {
	once.Do(func() {
		instance = &Reductor{
			logger:     logger,
			modelPages: make(ModelList),
		}
		instance.modelPages[""] = model
	})
	return instance
}

func Instance() *Reductor {
	return instance
}

func (rdc *Reductor) Model(page string) (mdl Model) {
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()

	// разыменовываем ссылку получаем копию объекта и возвращаем
	if pageModel, ok := rdc.modelPages[page]; ok {
		return pageModel
	}
	mdl = Model{}
	return mdl
}

// записываем модель,  возвращаем предыдущую
func (rdc *Reductor) SetModel(page string, model Model) (prevModel Model) {
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()
	prevModel = rdc.modelPages[page]
	rdc.modelPages[page] = model
	return prevModel
}

func (rdc *Reductor) Logger() *zap.SugaredLogger {
	return rdc.logger
}
