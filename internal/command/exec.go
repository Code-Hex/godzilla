package command

import (
	"io"
	"os/exec"
)

type Runner interface {
	Exec(string, ...string) error
	Run(string, ...string) (string, error)
}

type Executor struct {
	Stdout io.Writer
	Stderr io.Writer
}

func (e Executor) Exec(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = e.Stdout
	cmd.Stderr = e.Stderr
	return cmd.Run()
}

func (e Executor) Run(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}
