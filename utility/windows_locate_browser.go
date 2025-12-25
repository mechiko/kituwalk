//go:build windows
// +build windows

package utility

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ChromeExecutable returns a string which points to the preferred Chrome
// executable file.
var ChromeExecutable = LocateChrome

// LocateChrome returns a path to the Chrome binary, or an empty string if
// Chrome installation is not found.
func LocateChrome() string {
	pathsChrome := []string{
		os.Getenv("LocalAppData") + "/Google/Chrome/Application/chrome.exe",
		os.Getenv("ProgramFiles") + "/Google/Chrome/Application/chrome.exe",
		os.Getenv("ProgramFiles(x86)") + "/Google/Chrome/Application/chrome.exe",
		os.Getenv("LocalAppData") + "/Chromium/Application/chrome.exe",
		os.Getenv("ProgramFiles") + "/Chromium/Application/chrome.exe",
		os.Getenv("ProgramFiles(x86)") + "/Chromium/Application/chrome.exe",
	}
	for _, path := range pathsChrome {
		if path == "" {
			continue
		}
		if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
			return path
		}
	}
	return ""
}

func LocateYandex() string {
	pathsChrome := []string{
		os.Getenv("LocalAppData") + "/Yandex/YandexBrowser/Application/browser.exe",
	}
	for _, path := range pathsChrome {
		if path == "" {
			continue
		}
		if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
			return path
		}
	}
	return ""
}

func LocateFox() string {
	pathsChrome := []string{
		os.Getenv("ProgramFiles") + "/Firefox Developer Edition/firefox.exe",
		os.Getenv("ProgramFiles(x86)") + "/Firefox Developer Edition/firefox.exe",
		os.Getenv("ProgramFiles") + "/Mozilla Firefox/firefox.exe",
		os.Getenv("ProgramFiles(x86)") + "/Mozilla Firefox/firefox.exe",
	}
	for _, path := range pathsChrome {
		if path == "" {
			continue
		}
		if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
			return path
		}
	}
	return ""
}

func LocateEdge() string {
	pathsChrome := []string{
		os.Getenv("ProgramFiles(x86)") + "/Microsoft/Edge/Application/msedge.exe",
	}
	for _, path := range pathsChrome {
		if path == "" {
			continue
		}
		if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
			return path
		}
	}
	return ""
}

func StartBrowser(browser Browser, url string) (err error) {
	switch browser {
	case Chrome:
		args := []string{}
		appUrl := "--app=" + url
		args = append(args, appUrl)
		bin := ChromeExecutable()
		if bin == "" {
			return fmt.Errorf("chrome executable not found")
		}
		err = exec.Command(bin, args...).Start()
	case Firefox:
		argsFox := []string{}
		appUrlFox := url
		argsFox = append(argsFox, appUrlFox)
		bin := LocateFox()
		if bin == "" {
			return fmt.Errorf("firefox executable not found")
		}
		err = exec.Command(bin, argsFox...).Start()
	case Edge:
		args := []string{}
		appUrl := "--app=" + url
		args = append(args, appUrl)
		bin := LocateEdge()
		if bin == "" {
			return fmt.Errorf("edge executable not found")
		}
		err = exec.Command(bin, args...).Start()
	case Yandex:
		args := []string{}
		appUrl := "--app=" + url
		args = append(args, appUrl)
		bin := LocateYandex()
		if bin == "" {
			return fmt.Errorf("yandex browser executable not found")
		}
		err = exec.Command(bin, args...).Start()
	default:
		l := strings.ToLower(url)
		if strings.HasPrefix(l, "https://") {
			err = OpenHttpsLinkInShell(url)
		} else {
			err = OpenHttpLinkInShell(url)
		}
	}
	return err
}
