package locale

import (
	"bytes"
	"os/exec"
	"strings"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

var detectors = []detector{
	detectViaEnvLanguage,
	detectViaEnvLc,
	detectViaSyscall,
	detectViaPowershell,
	detectViaRegistry,
}

func detectViaPowershell() (langs []string, err error) {
	path, err := exec.LookPath("powershell.exe")
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(path, "Get-Culture | select -exp Name")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	return []string{strings.Trim(out.String(), "\r\n")}, nil
}

func detectViaSyscall() (langs []string, err error) {
	return windows.GetSystemPreferredUILanguages(windows.MUI_LANGUAGE_NAME)
}

// detectViaRegistry will detect language via Windows Registry
//
// ref: https://renenyffenegger.ch/notes/Windows/registry/tree/HKEY_CURRENT_USER/Control-Panel/International/index
func detectViaRegistry() (langs []string, err error) {
	defer func() {
		if err != nil {
			err = &Error{"detect via registry", err}
		}
	}()

	key, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\International`, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer key.Close()

	lang, _, err := key.GetStringValue("LocaleName")
	if err != nil {
		return nil, err
	}

	return []string{lang}, nil
}
