package dbscan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindA3Name(t *testing.T) {
	// существует каталог и бд
	name := findA3DbName("cmd")
	assert.NotEqual(t, name, "", "db found")
	// существует каталог и нет бд
	name = findA3DbName(".")
	assert.Equal(t, name, "", "db found")
	// не существует каталог
	name = findA3DbName("")
	assert.Equal(t, name, "", "db not found")
	// не существует каталог
	name = findA3DbName("ccc")
	assert.Equal(t, name, "", "path not found")
}

func TestFind4zName(t *testing.T) {
	// существует каталог и бд
	name := find4zDbName("cmd")
	assert.NotEqual(t, name, "", "db not found")
	// существует каталог но нет бд
	name = find4zDbName(".")
	assert.NotEqual(t, name, "", "db not found")
	// не существует каталог
	name = find4zDbName("ccc")
	assert.Equal(t, name, "", "path not found")
	// не существует каталог
	name = find4zDbName("")
	assert.Equal(t, name, "", "path not found")
}
