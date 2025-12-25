package dbscan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testsParse = []struct {
	name   string
	err    bool
	info   DbInfo
	result DbInfo
}{
	// the table itself
	{"test 0", true, DbInfo{File: "config.db", Driver: "sqlite"}, DbInfo{File: "config.db", Driver: "sqlite", Path: "cmd", Exists: true}},
	{"test 0", false, DbInfo{File: "config.db", Driver: "sqlite", Path: "cmd"}, DbInfo{File: "config.db", Driver: "sqlite", Path: "cmd", Exists: true}},
	{"test 0", true, DbInfo{File: "cmd/config.db", Driver: "sqlite"}, DbInfo{}},
	{"test 1", true, DbInfo{File: "", Driver: "sqlite"}, DbInfo{}},
	{"test 2", true, DbInfo{File: "config.db", Driver: ""}, DbInfo{}},
	{"test 4", true, DbInfo{File: "cm/config.db", Driver: "sqlite", Path: "cmd"}, DbInfo{}},
	{"test 4", true, DbInfo{File: "030000679428.db", Driver: "mssql"}, DbInfo{}},
	{"test 5", true, DbInfo{File: "030000679428.db", Name: "030000679428.db", Driver: "mssql"}, DbInfo{}},
	{"test 6", false, DbInfo{File: "030000679428.db", Name: "030000679428", Driver: "mssql"}, DbInfo{File: "030000679428.db", Name: "030000679428", Driver: "mssql", Exists: true}},
	{"test 7", false, DbInfo{File: "", Name: "030000679428", Driver: "mssql"}, DbInfo{File: "", Name: "030000679428", Driver: "mssql", Exists: true}},
}

func TestParse(t *testing.T) {
	// The execution loop
	for _, tt := range testsParse {
		t.Run(tt.name, func(t *testing.T) {
			dbi, err := ParseDbInfo(&tt.info)
			if tt.err {
				assert.NotNil(t, err, "ожидаем ошибку")
			} else {
				// ожидаем отсутствие ошибки
				out := *dbi
				assert.Nil(t, err)
				assert.Equal(t, out, tt.result, "ожидаемое значение")
			}
		})
	}
}
