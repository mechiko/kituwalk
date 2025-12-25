//go:build windows
// +build windows

package utility

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows"
)

func OpenFileInShell(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	// Convert backslashes to forward slashes for file URLs
	urlPath := filepath.ToSlash(absPath)
	// Prepend slash for absolute paths on Windows
	if !strings.HasPrefix(urlPath, "/") {
		urlPath = "/" + urlPath
	}
	return startShell("file://" + urlPath)
}

// открыть ссылку в браузере
func OpenHttpLinkInShell(urlRaw string) error {
	if urlRaw == "" {
		return fmt.Errorf("empty URL provided")
	}
	// Parse URL without scheme first
	parsedURL, err := url.Parse(urlRaw)
	if err != nil {
		return err
	}
	parsedURL.Scheme = "http"
	return startShell(parsedURL.String())
}

func OpenHttpsLinkInShell(urlRaw string) error {
	if urlRaw == "" {
		return fmt.Errorf("empty URL provided")
	}
	// Parse URL without scheme first
	parsedURL, err := url.Parse(urlRaw)
	if err != nil {
		return err
	}
	parsedURL.Scheme = "https"
	return startShell(parsedURL.String())
}

func startShell(url string) error {
	if url == "" {
		return fmt.Errorf("empty URL provided to startShell")
	}
	err := windows.ShellExecute(0, nil, windows.StringToUTF16Ptr(url), nil, nil, windows.SW_SHOWNORMAL)
	if err != nil {
		return fmt.Errorf("failed to execute shell command for URL '%s': %w", url, err)
	}
	return nil
}

type Browser string

// типы зарезервированных БД
const (
	Default Browser = ""
	Chrome  Browser = "chrome"
	Edge    Browser = "msedge"
	Firefox Browser = "firefox"
	Yandex  Browser = "yandex"
	Opera   Browser = "opera"
)

func isValidBrowse(b string) bool {
	switch Browser(b) {
	case Default, Chrome, Edge, Firefox:
		return true
	default:
		return false
	}
}

func OpenHttpBrowser(urlRaw string, browser Browser) error {
	if urlRaw == "" {
		return fmt.Errorf("empty URL provided")
	}
	// Parse URL without scheme first
	parsedURL, err := url.Parse(urlRaw)
	if err != nil {
		return err
	}
	parsedURL.Scheme = "http"
	return StartBrowser(browser, parsedURL.String())
}

func OpenHttpsBrowser(urlRaw string, browser Browser) error {
	if urlRaw == "" {
		return fmt.Errorf("empty URL provided")
	}
	// Parse URL without scheme first
	parsedURL, err := url.Parse(urlRaw)
	if err != nil {
		return err
	}
	parsedURL.Scheme = "https"
	return StartBrowser(browser, parsedURL.String())
}
