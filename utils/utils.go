package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/kuun/slog"
)

type _logType struct{}

var log = slog.GetLogger(&_logType{})

// cmdRunOutput runs the Linux command with the specified
// arguments and returns the output.
func CmdRunOutput(cmd string, args ...string) (string, error) {
	var out []byte
	var err error
	argsStr := strings.Join(args, " ")
	log.Infof("%s %s", cmd, argsStr)

	proc := exec.Command(cmd, args...)
	out, err = proc.Output()

	if err != nil {
		msg := fmt.Sprintf("cmdRunOutput: '%s %s', error: %v", cmd, argsStr, err)
		if exitErr, ok := err.(*exec.ExitError); ok {
			msg += ", stderr: " + string(exitErr.Stderr)
		}
		log.Infof(msg)
		return "", errors.New(msg)
	}
	return string(out), nil
}

// cmdRun runs the command with the specified arguments.
func CmdRun(cmd string, args ...string) error {
	_, err := CmdRunOutput(cmd, args...)
	return err
}

func ScriptRun(script string) error {
	_, err := CmdRunOutput("sh", "-c", script)
	return err
}
