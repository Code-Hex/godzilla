package gozla

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/Code-Hex/godzilla/internal/git"
	"github.com/Code-Hex/godzilla/internal/license"
	"github.com/Code-Hex/godzilla/internal/ui"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-gitconfig"
)

type README struct {
	Name       string
	Email      string
	GithubUser string
}

type LICENSE struct {
	Year       string
	Project    string
	GithubUser string
}

func runNew(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("Missing project name")
	}

	project := args[0]
	if err := os.Mkdir(project, 0755); err != nil {
		return err
	}

	if err := os.Chdir(project); err != nil {
		return err
	}
	choose := ui.Choose("Which license do you want to use?", []string{
		"MIT License",
		"The Unlicense",
		"Apache License 2.0",
		"Mozilla Public License 2.0",
		"GNU General Public License v3.0",
		"GNU Affero General Public License v3.0",
		"GNU Lesser General Public License v3.0",
	}, "MIT License")
	t := template.New("initialization")
	if err := GenerateLICENSE(t, project, choose); err != nil {
		return err
	}
	if err := GeneratePkgGo(project); err != nil {
		return err
	}
	if err := GenerateREADME(t, project); err != nil {
		return err
	}
	if err := GenerateChanges(); err != nil {
		return err
	}
	if err := GenerateGozlaToml(); err != nil {
		return err
	}
	if err := GenerateTravisyml(); err != nil {
		return err
	}
	if err := GenerateGitignore(); err != nil {
		return err
	}
	ui.Printf("Initializing git %s\n", project)
	ui.Printf("[%s] $ git init\n", project)
	if err := git.Init(); err != nil {
		return err
	}

	ui.Printf("[%s] $ git add .\n", project)
	if err := git.Add("."); err != nil {
		return err
	}

	ui.Printf("Finished to create %s project\n", project)

	return nil
}

func GenerateREADME(t *template.Template, project string) error {
	ui.Printf("Writing README.md\n")
	f, err := os.Create("README.md")
	if err != nil {
		return err
	}
	defer f.Close()
	tmpl, err := t.Parse(`{{.Name}} - New cli tool.
====

# SYNOPSIS

    import "github.com/{{.GithubUser}}/{{.Name}}"

# DESCRIPTION

{{.Name}} is ...

# INSTALLATION

    go get github.com/{{.GithubUser}}/{{.Name}}

# AUTHOR

[{{.GithubUser}}](https://github.com/{{.GithubUser}}) <{{.Email}}>`)
	if err != nil {
		return err
	}

	var readme README
	readme.Name = project
	readme.Email, err = gitconfig.Email()
	if err != nil {
		return err
	}
	readme.GithubUser, err = gitconfig.GithubUser()
	if err != nil {
		return err
	}
	return tmpl.Execute(f, readme)
}

func GenerateLICENSE(t *template.Template, project, kind string) error {
	ui.Printf("Writing LICENSE\n")
	f, err := os.Create("LICENSE")
	if err != nil {
		return err
	}
	defer f.Close()

	var choose string
	switch kind {
	case "MIT License":
		choose = license.MIT
	case "The Unlicense":
		choose = license.Unlicense
	case "Apache License 2.0":
		choose = license.Apache
	case "Mozilla Public License 2.0":
		choose = license.MPL2
	case "GNU General Public License v3.0":
		choose = license.GPLv3
	case "GNU Affero General Public License v3.0":
		choose = license.AGPLv3
	case "GNU Lesser General Public License v3.0":
		choose = license.LGPLv3
	}

	var licenseFile LICENSE
	licenseFile.Project = project
	licenseFile.Year = fmt.Sprintf("%d", time.Now().Year())
	licenseFile.GithubUser, err = gitconfig.GithubUser()
	if err != nil {
		return err
	}

	tmpl, err := t.Parse(choose)

	return tmpl.Execute(f, licenseFile)
}

func GenerateGitignore() error {
	ui.Printf("Writing .gitignore\n")
	f, err := os.Create(".gitignore")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(`# Compiled Object files, Static and Dynamic libs (Shared Objects)
*.o
*.a
*.so

# Folders
_build/
_obj/
_test/

# Architecture specific extensions/prefixes
*.[568vq]
[568vq].out

*.cgo1.go
*.cgo2.c
_cgo_defun.c
_cgo_gotypes.go
_cgo_export.*

_testmain.go

*.exe
*.test
*.prof

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# External packages folder
vendor/`)
	return nil
}

func GenerateTravisyml() error {
	ui.Printf("Writing .travis.yml\n")
	f, err := os.Create(".travis.yml")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(`language: go
sudo: false
go:
  - 1.7.3
  - tip`)

	return nil
}
func GenerateGozlaToml() error {
	f, err := os.Create("gozla.toml")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(`[build]
# dir = "."
# os = ["mac", "windows", "linux"]
# arch = ["386", "amd64"]
# archive = ["zip", "tar.gz"]

[migrate]
# badges = ['travis', 'godoc', 'goreport']`)

	return nil
}

func GenerateChanges() error {
	ui.Printf("Writing Changes\n")
	f, err := os.Create("Changes")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(`Changes
=======

{{.Version}} - {{.Date}}

    - original version
`)

	return nil
}

func GeneratePkgGo(project string) error {
	pkg := strings.Split(project, "-")
	pkgname := pkg[len(pkg)-1]
	ui.Printf("Writing %s.go\n", pkgname)
	f, err := os.Create(pkgname + ".go")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(`package ` + pkgname + `

const version = "0.0.1"
`)

	return nil
}
