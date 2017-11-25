package runner

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"io/ioutil"
	"encoding/json"
)

var settings = struct {
	ConfigPath   string   `json:"config_path"`
	Root         string   `json:"root"`
	TmpPath      string   `json:"tmp_path"`
	BuildName    string   `json:"build_name"`
	BuildLog     string   `json:"build_log"`
	ValidExt     []string `json:"valid_ext"`
	NoRebuildExt []string `json:"no_rebuild_ext"`
	Ignored      []string `json:"ignored"`
	BuildDelay   int      `json:"build_delay"`
	PreExec      []string `json:"pre_exec"`
}{
	ConfigPath:   "./fresh.json",
	Root:         ".",
	TmpPath:      "./tmp",
	BuildName:    "runner-build",
	BuildLog:     "runner-build-errors.log",
	ValidExt:     []string{"go", "html", "css", "js"},
	NoRebuildExt: []string{"json", "conf", "gitignore", "bat"},
	Ignored:      []string{"tmp"},
	BuildDelay:   600,
	PreExec:      []string{},
}

func loadRunnerConfigSettings() {
	if _, err := os.Stat(settings.ConfigPath); err != nil {
		return
	}

	logger.Printf("Loading settings from %s", settings.ConfigPath)
	buf, err := ioutil.ReadFile("fresh.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(buf, &settings)
	if err != nil {
		return
	}
}

func initSettings() {
	loadRunnerConfigSettings()
}

func buildPath() string {
	p := filepath.Join(settings.TmpPath, settings.BuildName)
	if runtime.GOOS == "windows" && filepath.Ext(p) != ".exe" {
		p += ".exe"
	}
	return p
}

func buildErrorsFilePath() string {
	return filepath.Join(settings.TmpPath, settings.BuildLog)
}

func buildDelay() time.Duration {
	return time.Duration(settings.BuildDelay)
}
