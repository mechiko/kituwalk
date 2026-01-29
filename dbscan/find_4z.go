package dbscan

import (
	"path/filepath"
	"regexp"
	"strings"
)

// using [8|9|aA|bB], which (a) includes the literal | and (b) is verbose. This will accept invalid characters and miss the intent. Use [89abAB] for the RFC 4122 variant and prefer a raw string
// var trueRegex = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
// var trueRegex = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
var trueRegex = regexp.MustCompile(`^[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}\.db$`)

func find4zDbName(dir string) string {
	if dir == "" {
		dir = "."
	}
	if files, err := FilteredDirectory(trueRegex, dir); err != nil {
		return ""
	} else {
		if len(files) == 0 {
			return ""
		}
		return files[0]
	}
}

func find4zName(dir string) string {
	// discover a 4z file under the current directory
	findName := find4zDbName(dir)
	if findName == "" {
		return ""
	}
	// strip off any directory components, then drop the extension
	base := filepath.Base(findName)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}
