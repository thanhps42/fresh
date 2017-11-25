package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"io/ioutil"
	"encoding/json"
	"reflect"
)

var settings = map[string]interface{}{
	"config_path":       "./fresh.json",
	"root":              ".",
	"tmp_path":          "./tmp",
	"build_name":        "runner-build",
	"build_log":         "runner-build-errors.log",
	"valid_ext":         []string{"go"},
	"no_rebuild_ext":    []string{"html", "css", "js", "json", "conf", "gitignore", "bat"},
	"ignored":           []string{"views", "public", "static", "assets", "tmp"},
	"build_delay":       600,
	"colors":            true,
	"log_color_main":    "cyan",
	"log_color_build":   "yellow",
	"log_color_runner":  "green",
	"log_color_watcher": "magenta",
	"log_color_app":     "",
}

var colors = map[string]string{
	"reset":          "0",
	"black":          "30",
	"red":            "31",
	"green":          "32",
	"yellow":         "33",
	"blue":           "34",
	"magenta":        "35",
	"cyan":           "36",
	"white":          "37",
	"bold_black":     "30;1",
	"bold_red":       "31;1",
	"bold_green":     "32;1",
	"bold_yellow":    "33;1",
	"bold_blue":      "34;1",
	"bold_magenta":   "35;1",
	"bold_cyan":      "36;1",
	"bold_white":     "37;1",
	"bright_black":   "30;2",
	"bright_red":     "31;2",
	"bright_green":   "32;2",
	"bright_yellow":  "33;2",
	"bright_blue":    "34;2",
	"bright_magenta": "35;2",
	"bright_cyan":    "36;2",
	"bright_white":   "37;2",
}

func logColor(logName string) string {
	settingsKey := fmt.Sprintf("log_color_%s", logName)
	colorName := settings[settingsKey].(string)

	return colors[colorName]
}

func loadRunnerConfigSettings() {
	if _, err := os.Stat(configPath()); err != nil {
		return
	}

	logger.Printf("Loading settings from %s", configPath())
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

func root() string {
	return settings["root"].(string)
}

func tmpPath() string {
	return settings["tmp_path"].(string)
}

func buildName() string {
	return settings["build_name"].(string)
}
func buildPath() string {
	p := filepath.Join(tmpPath(), buildName())
	if runtime.GOOS == "windows" && filepath.Ext(p) != ".exe" {
		p += ".exe"
	}
	return p
}

func buildErrorsFileName() string {
	return settings["build_log"].(string)
}

func buildErrorsFilePath() string {
	return filepath.Join(tmpPath(), buildErrorsFileName())
}

func configPath() string {
	return settings["config_path"].(string)
}

func buildDelay() time.Duration {
	v := settings["build_delay"]
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float64:
		return time.Duration(int(v.(float64)))
	case reflect.Int:
		return time.Duration(v.(int))
	default:
		return time.Duration(1000)
	}
}
