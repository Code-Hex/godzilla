package gozla

import (
	"os"

	"github.com/Code-Hex/godzilla/internal/ui"
	"github.com/spf13/cobra"
)

const (
	Version = "0.0.1"
	msg     = "gozla v" + Version + "\n"
)

type gozil struct {
	root    *cobra.Command
	version bool
	trace   bool
}

func New() *gozil {
	return new(gozil).init()
}

func (gz *gozil) init() *gozil {
	gz.root = &cobra.Command{
		Use:           "gozla",
		Short:         "`gozla`",
		Long:          "`gozla`",
		RunE:          gz.Execute,
		SilenceErrors: true,
	}
	gz.root.Flags().BoolVarP(&gz.trace, "trace", "t", false, "display detail error messages")
	gz.root.Flags().BoolVarP(&gz.version, "version", "v", false, "display the version of gozil and exit")

	// See commands.go
	for _, subCmd := range commands {
		gz.root.AddCommand(subCmd)
	}

	return gz
}

func (gz *gozil) Execute(cmd *cobra.Command, args []string) error {
	if gz.version {
		os.Stdout.Write([]byte(msg))
		return nil
	}
	return cmd.Usage()
}

func (gz gozil) Run() int {
	if err := gz.root.Execute(); err != nil {
		if gz.trace {
			ui.Errorf("Error:\n%+v\n", err)
		} else {
			ui.Errorf("Error:\n  %v\n", err)
		}
		return 1
	}
	return 0
}
