package gui

import (
	"fmt"
	"kituwalk/utility"
)

// открываем в браузере ссылку
func Open(url string) error {
	if url == "" {
		return fmt.Errorf("пустой url")
	}
	if err := utility.OpenHttpLinkInShell(url); err != nil {
		return err
	}
	return nil
}

// открываем в эксплорере текущую папку программы
func OpenDir(dir string) (err error) {
	if !utility.PathOrFileExists(dir) {
		return fmt.Errorf("path not found: %s", dir)
	}
	return utility.OpenFileInShell(dir)
}

// открываем в эксплорере файл по имени
func OpenFile(file string) (err error) {
	if !utility.PathOrFileExists(file) {
		return fmt.Errorf("file not found: %s", file)
	}
	return utility.OpenFileInShell(file)
}
