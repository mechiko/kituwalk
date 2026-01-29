package dbscan

import (
	"path/filepath"
	"regexp"
	"strings"
)

var reA3DbName = regexp.MustCompile(`^0[0-9]{11}\.db$`)

// find a3 db in dir
func findA3Name(dir string) string {
	findName := findA3DbName(dir)
	if findName == "" {
		return ""
	}
	base := filepath.Base(findName)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}

func findA3DbName(dir string) string {
	if dir == "" {
		dir = "."
	}
	if files, err := FilteredDirectory(reA3DbName, dir); err != nil {
		return ""
	} else {
		if len(files) == 0 {
			return ""
		}
		return files[0]
	}
}
