package gozla

import (
	"os"

	"github.com/Code-Hex/godzilla/internal/ui"
	"github.com/spf13/cobra"
)

func runClean(cmd *cobra.Command, args []string) error {
	if ui.YN("Remove all files or directories under the _build?", true) {
		return os.RemoveAll("_build")
	}
	return nil
}
