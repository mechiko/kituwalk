package dbscan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name  string
	err   bool
	input DbInfo
}{
	// the table itself
	{"test 0", false, DbInfo{File: "cmd/config.db", Driver: "sqlite"}},
	{"test 1", false, DbInfo{File: "config.db", Path: "cmd", Driver: "sqlite"}},
	{"test 2", false, DbInfo{File: "4zupper.db", Path: "cmd/.nevakod/4zupper", Driver: "sqlite"}},
	{"test 3", true, DbInfo{File: "030000679428.db", Driver: "mssql"}},
	{"правильное подключение к mssql", false, DbInfo{Name: "030000679428", Driver: "mssql"}},
}

func TestIsConnect(t *testing.T) {
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.IsConnected()
			if tt.err {
				assert.NotNil(t, err)
			} else {
				// ожидаем отсутствие ошибки
				assert.Nil(t, err)
			}
		})
	}
}

func TestConnect(t *testing.T) {
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sess, err := tt.input.Connect()
			if tt.err {
				assert.NotNil(t, err)
			} else {
				// ожидаем отсутствие ошибки
				assert.Nil(t, err)
				pingErr := sess.Ping()
				assert.Nil(t, pingErr)
			}
		})
	}
}
