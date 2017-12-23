package runner

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"fmt"
)

func preExec() {
	for _, s := range settings.PreExec {
		buildLog(fmt.Sprintf("Exec: %s", s))

		cmd := exec.Command("cmd", "/C", s)
		cmd.Run()
	}
}

func build() (string, bool) {
	buildLog("Building...")

	cmd := exec.Command("go", "build", "-o", buildPath(), settings.WorkingDirectory)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	io.Copy(os.Stdout, stdout)
	errBuf, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()
	if err != nil {
		return string(errBuf), false
	}

	return "", true
}
