//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Installing dependencies
func Prepare() error {
	return sh.Run("go", "mod", "download")
}

// Linting any errors
func Lint() error {
	return sh.Run("go", "vet", "./...")
}

// Building single binary
func Build() error {
	return sh.Run("go", "build", "main.go")
}

// Run code to debug something
func Run() error {
	return sh.Run("go", "run", "main.go")
}

// Test package
func Test() error {
	return sh.Run("go", "test", "./...")
}

// (Re)generate go-files from templates
func Generate() error {
	env := map[string]string{
		"PYTHONPATH": "peg_generator/",
	}
	err := sh.RunWith(
		env,
		"python3",
		"peg_generator/pegen",
		"-v",
		"go",
		"grammar/python.gram",
		"grammar/Tokens",
		"--out",
		"parser/parse.go",
	)
	if err != nil {
		return err
	}
	return sh.Run("go", "fmt", "parser/parse.go")
}
