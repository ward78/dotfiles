package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"errors"
	"os/user"
//	"path/filepath"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	err := verifySystem()
	if err != nil {
		return err
	}
	return nil
}

func verifySystem() error {
	if runtime.GOOS != "linux" {
		return errors.New("script requires linux OS")
	}
	u, err := user.Current()
	if err != nil {
		return err
	}
	if u.Gid != "0" {
		return errors.New("root/sudo access is required")
	}
	return nil
}

func backupRun() error {
	oldLocation := "/var/www/html/test.txt"
	newLocation := "/var/www/html/src/test.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func backupReload() error {
	return nil
}
