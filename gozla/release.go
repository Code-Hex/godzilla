package gozla

import (
	"fmt"
	"os"

	"github.com/Code-Hex/go-version-update"
	"github.com/Code-Hex/godzilla/internal/git"
	"github.com/Code-Hex/godzilla/internal/ui"
	"github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
)

type release struct {
	AfterScript string `toml:"after_script"`
}

func runRelease(cmd *cobra.Command, args []string) error {
	if err := releaseVersionCheck(); err != nil {
		return err
	}
	return nil
}

func showRelease() error {
	fmt.Println("show!")
	return nil
}

func noTestRelease() error {
	fmt.Println("no-test")
	return nil
}

func releaseVersionCheck() error {
	now, err := previousTagVersion()
	if err == nil {
		ui.Printf("The most recently tagged version is %s\n", now)
	}
	next := enterNextVersion()
	files, err := update.GrepVersion(".")
	if err != nil {
		return err
	}

	// Rewrite the version variable on the go file
	for _, f := range files {
		fi, err := os.OpenFile(f.Path, os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		if err := update.NextVersion(fi, next, f.Path); err != nil {
			fi.Close()
			return err
		}
		fi.Close()
	}

	return nil
}

// Enter input the next version with stdin
func enterNextVersion() (input string) {
	for {
		ui.Printf("Next version is: ")
		input = ui.Readline()
		if input != "" {
			// Check the semantic version format
			_, err := version.NewVersion(input)
			if err != nil {
				ui.Errorf("Error: %s\n", err.Error())
				continue
			}
			break
		}
	}

	return
}

func previousTagVersion() (string, error) {
	return git.Describe("--tags", "--abbrev=0")
}

func getRevisionOfHEAD() (string, error) {
	return git.Log("--oneline", `--format="%h"`, "-n", "1")
}

func getCommitLog(tagA, tagB string) (string, error) {
	return git.Log("--oneline", "--abbrev-commit", tagA+".."+tagB)
}
