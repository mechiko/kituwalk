package dbscan

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func pathCheck(path string) error {
	// Check path existence
	info, err := os.Stat(path)
	if err != nil {
		// Handle different error scenarios
		switch {
		case os.IsNotExist(err):
			return fmt.Errorf("path does not exist: %s", path)
		case os.IsPermission(err):
			return fmt.Errorf("permission denied for path: %s", path)
		default:
			return err
		}
	}

	// Additional checks
	if info.IsDir() {
		// Check directory readability
		_, readErr := os.ReadDir(path)
		if readErr != nil {
			return fmt.Errorf("cannot read directory: %s", path)
		}
	}

	return nil
}

// EX: re, err := regexp.Compile(`^0[0-9]{11}\.db$`)
//
//	files, err := FilteredDirectory(re, ""); err != nil {
//
// "^[a-zA-Z0-9].*\\.db$"
// `^[a-zA-Z0-9].*\.db$`
// глубина поиска 0 только в указанном каталоге
func FilteredDirectory(re *regexp.Regexp, dir string) ([]string, error) {
	if err := pathCheck(dir); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if !filepath.IsAbs(dir) {
		dir, _ = filepath.Abs(dir)
	}
	files := []string{}
	base := filepath.Base(dir)
	walk := func(path string, d fs.DirEntry, err error) error {
		// каталоги не равные base пропускаем
		if d.IsDir() && (d.Name() != base) {
			return fs.SkipDir
		}
		if d.IsDir() {
			return nil
		}
		if !re.MatchString(d.Name()) {
			return nil
		}
		files = append(files, path)
		return nil
	}
	err := filepath.WalkDir(dir, walk)
	return files, err
}
