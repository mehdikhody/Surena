package env

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func GetXrayConfigPath() string {
	configPath := os.Getenv("XRAY_CONFIG_FILE")
	if configPath == "" {
		configPath = "config.json"
	}

	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return configPath
	}

	return absPath
}

func GetXrayExecutablePath() string {
	configPath := os.Getenv("XRAY_EXECUTABLE")
	if configPath == "" {
		configPath = "xray"
	}

	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return configPath
	}

	return absPath
}

func GetXrayVersion() string {
	cmd := exec.Command(GetXrayExecutablePath(), "--version")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	versionRegex, err := regexp.Compile("^Xray\\s(\\d.\\d.\\d)")
	if err != nil {
		panic(err)
	}

	matches := versionRegex.FindStringSubmatch(string(out))
	if len(matches) < 2 {
		panic("Could not find version in output")
	}

	return matches[1]
}
