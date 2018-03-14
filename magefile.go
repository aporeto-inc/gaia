// +build mage

// nolint
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	domingo "github.com/aporeto-inc/domingo/golang"
)

// Init initialize the project.
func Init() {
	mg.Deps(
		domingo.InstallDependencies,
		Elegen,
	)
}

// Test runs the unit tests.
func Test() {
	mg.Deps(
		domingo.Test,
	)
}

// Elegen installs the latest elegen command.
func Elegen() error {

	return sh.Run("go", "get", "-u", "github.com/aporeto-inc/elemental/cmd/elegen")
}

// All builds all models.
func All() {

	mg.Deps(
		func() error { return codegenModel("v1") },
	)
}

func codegenModel(modelPath string) error {

	// if err := sh.Run("rego", "beautify", "--src", path.Join(modelPath, "/specs")); err != nil {
	// 	return err
	// }

	codegenPath := path.Join(modelPath, "codegen")
	specPath := path.Join(modelPath, "specs")
	golangPath := path.Join(modelPath, "golang")

	if err := os.RemoveAll(codegenPath); err != nil {
		return err
	}

	if err := sh.Run("elegen", "folder", "-d", specPath, "-o", codegenPath); err != nil {
		return err
	}

	defer os.RemoveAll(codegenPath)

	if err := eachFile(
		golangPath,
		func(f os.FileInfo) error {
			n := f.Name()
			if path.Ext(n) == ".go" {
				if err := os.Remove(path.Join(golangPath, n)); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
		return err
	}

	if err := eachFile(
		path.Join(codegenPath, "elemental"),
		func(f os.FileInfo) error {
			n := f.Name()
			if path.Ext(n) == ".go" {
				if err := copyFile(path.Join(codegenPath, "elemental", n), path.Join(golangPath, n)); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
		return err
	}

	fmt.Println("generated: ", modelPath)

	return nil
}

func eachFile(dir string, job func(f os.FileInfo) error) error {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if err := job(f); err != nil {
			return err
		}
	}

	return nil
}

func copyFile(src, dst string) error {

	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}

	return nil
}
