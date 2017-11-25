package runner

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

func initFolders() {
	runnerLog("InitFolders")
	path := settings.TmpPath
	runnerLog("mkdir %s", path)
	err := os.Mkdir(path, 0755)
	if err != nil {
		runnerLog(err.Error())
	}
}

func isTmpDir(path string) bool {
	absolutePath, _ := filepath.Abs(path)
	absoluteTmpPath, _ := filepath.Abs(settings.TmpPath)

	return absolutePath == absoluteTmpPath
}

func isIgnoredFolder(path string) bool {
	paths := strings.Split(path, "/")
	if len(paths) <= 0 {
		return false
	}

	for _, value := range settings.Ignored {
		if value == paths[0] {
			return true
		}
	}
	return false
}

func isWatchedFile(path string) bool {
	absolutePath, _ := filepath.Abs(path)
	absoluteTmpPath, _ := filepath.Abs(settings.TmpPath)

	if strings.HasPrefix(absolutePath, absoluteTmpPath) {
		return false
	}

	ext := filepath.Ext(path)

	for _, e := range settings.ValidExt {
		if fmt.Sprintf(".%s", e) == ext {
			return true
		}
	}

	return false
}

func shouldRebuild(eventName string) bool {
	for _, e := range settings.NoRebuildExt {
		ext := "." + e
		fileName := strings.Replace(strings.Split(eventName, ":")[0], `"`, "", -1)
		if strings.HasSuffix(fileName, ext) {
			return false
		}
	}

	return true
}

func createBuildErrorsLog(message string) bool {
	file, err := os.Create(buildErrorsFilePath())
	if err != nil {
		return false
	}

	_, err = file.WriteString(message)
	if err != nil {
		return false
	}

	return true
}

func removeBuildErrorsLog() error {
	err := os.Remove(buildErrorsFilePath())

	return err
}
