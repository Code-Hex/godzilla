package gozla

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/Code-Hex/godzilla/internal/command"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

type build struct {
	OS       []string `toml:"os"`
	Arch     []string `toml:"arch"`
	GCFlags  string   `toml:"gcflags"`
	LDFlags  string   `toml:"ldflags"`
	Parallel int      `toml:"parallel"`
	Target   string   `toml:"target_dir"`
	Name     string   `toml:"binary_name"`

	// to run archive
	Archive []string `toml:"archive"`
}

func runBuild(cmd *cobra.Command, args []string) error {
	buildDir := "_build"
	var config tomlConfig
	if _, err := toml.DecodeFile("gozla.toml", &config); err != nil {
		return err
	}

	buildOpts := config.Build

	if buildOpts.Parallel <= 0 {
		cpus := runtime.NumCPU()
		if cpus < 2 {
			buildOpts.Parallel = 1
		} else {
			buildOpts.Parallel = cpus - 1
		}

		// reason is here https://github.com/mitchellh/gox/blob/master/main.go#L60
		if runtime.GOOS == "solaris" {
			buildOpts.Parallel = 3
		}
	}

	if _, err := exec.LookPath("go"); err != nil {
		return errors.New(`This program necessary "go".
Please run after install "go"`)
	}

	os.MkdirAll(buildDir, os.ModeDir)

	var g errgroup.Group
	for _, os := range buildOpts.OS {
		for _, arch := range buildOpts.Arch {
			g.Go(func() error {
				return nil
			})
		}
	}

	if buildOpts.LDFlags != "" {
		options = append(options, mkflag("ldflags", buildOpts.LDFlags))
	}

	if buildOpts.GCFlags != "" {
		options = append(options, mkflag("gcflags", buildOpts.GCFlags))
	}

	if buildOpts.Output != "" {
		joined := filepath.Join(buildDir, "{{.OS}}_{{.Arch}}", buildOpts.Output)
		options = append(options, mkflag("output", joined))
	} else {
		joined := filepath.Join(buildDir, "{{.OS}}_{{.Arch}}", "{{.Dir}}")
		options = append(options, mkflag("output", joined))
	}

	return gox(options...)
}

func gox(args ...string) error {
	e := command.Executor{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	return e.Exec("gox", args...)
}

func mkflag(flag, val string) string {
	return fmt.Sprintf("-%s=%s", flag, val)
}
