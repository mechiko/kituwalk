package pdfproc

import "fmt"

// значения переменных для шаблона
// индекс страницы
// партия и другие
type Vars struct {
	value map[string]string
}

func NewVars() *Vars {
	out := &Vars{
		value: map[string]string{},
	}
	return out
}

func (v *Vars) Add(name string, value string) error {
	if name == "" {
		return fmt.Errorf("is empty key")
	}
	v.value[name] = value
	return nil
}

func (v *Vars) Get(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("is empty key")
	}
	if val, exist := v.value[name]; exist {
		return val, nil
	}
	return "", fmt.Errorf("not exist key %s", name)
}
