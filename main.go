package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	system, err := NewSystem()
	if err != nil {
		return err
	}
	dotfiles, err := getDotfiles(system.home)
	if err != nil {
		return err
	}
	err = backupDotfiles(dotfiles)
	return nil
}
