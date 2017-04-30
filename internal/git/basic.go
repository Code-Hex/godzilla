package git

import (
	"errors"
	"os"
	"strings"

	"github.com/Code-Hex/godzilla/internal/command"
)

var exec = &command.Executor{
	Stdout: os.Stdout,
	Stderr: os.Stderr,
}

func Init() error {
	return exec.Exec("git", "init")
}

func Add(args ...string) error {
	return exec.Exec("git", append([]string{"add"}, args...)...)
}

func Commit(args ...string) error {
	return exec.Exec("git", append([]string{"commit"}, args...)...)
}

func Log(args ...string) (string, error) {
	result, err := exec.Run("git", append([]string{"log"}, args...)...)
	if err != nil {
		return "", errors.New(strings.TrimSpace(result))
	}
	return "", nil
}

func Remote(args ...string) (string, error) {
	result, err := exec.Run("git", append([]string{"remote"}, args...)...)
	if err != nil {
		return "", errors.New(strings.TrimSpace(result))
	}
	return "", nil
}

func Describe(args ...string) (string, error) {
	result, err := exec.Run("git", append([]string{"describe"}, args...)...)
	if err != nil {
		return "", errors.New(strings.TrimSpace(result))
	}
	return "", nil
}
