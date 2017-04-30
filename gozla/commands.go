package gozla

import "github.com/spf13/cobra"

type tomlConfig struct {
	Build   build
	Release release
}

var commands = []*cobra.Command{
	cmdNew,
	cmdRelease,
	cmdBuild,
	cmdClean,
}

// subcommand lists
var (
	cmdNew = &cobra.Command{
		Use:           "new",
		Short:         "Create a new go distribution",
		Long:          `gozla`,
		RunE:          runNew, // See new.go
		SilenceErrors: true,
	}

	cmdRelease = &cobra.Command{
		Use:           "release",
		Short:         "Release your go distribution to github",
		Long:          `gozla`,
		RunE:          runRelease, // See release.go
		SilenceErrors: true,
	}

	cmdBuild = &cobra.Command{
		Use:           "build",
		Short:         "Build your go distribution tarball",
		Long:          `gozla`,
		RunE:          runBuild, // See build.go
		SilenceErrors: true,
	}

	cmdClean = &cobra.Command{
		Use:           "clean",
		Short:         "Clean your built go distribution tarballs",
		Long:          `gozla`,
		RunE:          runClean, // See clean.go
		SilenceErrors: true,
	}
)
