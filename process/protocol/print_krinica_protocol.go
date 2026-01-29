package protocol

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"kituwalk/process"
	"kituwalk/reductor"

	"kituwalk/utility"
)

//go:embed tmplKrinicaProtocol.html
var tmplKrinicaProtocol string

type Data struct {
	*process.Process
	Model reductor.Model
}

func PrintKrinicaProtocol(k *process.Process) (file []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	var buf bytes.Buffer

	model := reductor.Instance().Model("")
	data := Data{
		Process: k,
		Model:   model,
	}
	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"lastSerial": func(arr []*utility.CisInfo) string {
			if len(arr) > 0 {
				if arr[len(arr)-1] != nil {
					return arr[len(arr)-1].Serial
				}
			}
			return ""
		},
		"firstSerial": func(arr []*utility.CisInfo) string {
			if len(arr) > 0 {
				if arr[0] != nil {
					return arr[0].Serial
				}
			}
			return ""
		},
		"inc": func(i int) int {
			return i + 1
		},
		"noescape": func(str string) template.HTML {
			return template.HTML(str)
		},
		"pallet": func(s string) string {
			if len(s) != 41 {
				return s
			}
			return fmt.Sprintf("(%s)%s(%s)%s(%s)%s(%s)%s(%s)%s", s[0:2], s[2:16], s[16:18], s[18:24], s[24:26], s[26:30], s[30:32], s[32:34], s[34:36], s[36:])
		},
	}
	t, err := template.New("protokol").Funcs(funcMap).Parse(tmplKrinicaProtocol)
	if err != nil {
		return nil, fmt.Errorf("parse template: %w", err)
	}
	err = t.ExecuteTemplate(&buf, "protokol", data)
	if err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}
