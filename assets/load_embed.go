package assets

import (
	"fmt"
	"io"
	"kituwalk/domain"
	"kituwalk/embeded"
	"path"
	"path/filepath"
	"strings"
)

func (a *Assets) loadEmbed() (err error) {
	entries, err := embeded.EmbeddedAssets.ReadDir("assets")
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	for _, file := range entries {
		if !file.IsDir() {
			name := strings.ToLower(file.Name())
			ext := filepath.Ext(name)
			if len(ext) > 1 {
				ext = ext[1:]
			}
			base := name[:len(name)-len(filepath.Ext(name))]
			switch ext {
			case "jpg":
				contentBytes, err := readContentFile(file.Name())
				if err != nil {
					return fmt.Errorf("Error reading file: %v", err)
				}
				// Convert the byte slice to a string
				a.jpg[base] = contentBytes
			case "png":
				contentBytes, err := readContentFile(file.Name())
				if err != nil {
					return fmt.Errorf("Error reading file: %v", err)
				}
				// Convert the byte slice to a string
				a.png[base] = contentBytes
			case "json":
				contentBytes, err := readContentFile(file.Name())
				if err != nil {
					return fmt.Errorf("Error reading file: %v", err)
				}
				// Convert the byte slice to a string
				a.json[base] = contentBytes
				out, err := domain.NewMarkTemplate(contentBytes)
				if out == nil || err != nil {
					return fmt.Errorf("new marktemplate error %w", err)
				}
				if out.Name == "" {
					return fmt.Errorf("new marktemplate name empty")
				}
				if _, ok := a.templateNames[out.Name]; ok {
					return fmt.Errorf("marktemplate %s already present", out.Name)
				}
				a.templateNames[out.Name] = base
			}
		}
	}
	return nil
}

func readContentFile(filePath string) ([]byte, error) {
	fpath := path.Join("assets", filePath)
	if file, err := embeded.EmbeddedAssets.Open(fpath); err != nil {
		return nil, fmt.Errorf("%w", err)
	} else {
		if out, err := io.ReadAll(file); err != nil {
			return nil, fmt.Errorf("%w", err)
		} else {
			return out, nil
		}
	}
}
