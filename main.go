package main

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
	//"path/filepath"
	//"path"
)

type authUser struct {
	name string
	uid  string
	gid  string
	dir  string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	authUser, err := getAuthUser()
	if err != nil {
		return err
	}
	fmt.Println(authUser.name)
	fmt.Println(authUser.uid)
	fmt.Println(authUser.gid)
	fmt.Println(authUser.dir)
	return nil
}

func getAuthUser() (*authUser, error) {
	// Make sure OS is linux
	if runtime.GOOS != "linux" {
		return nil, errors.New("script requires linux OS")
	}

	// Get current user
	admin, err := user.Current()
	if err != nil {
		return nil, err
	}

	// Make sure user has root access
	if admin.Name != "root" {
		return nil, errors.New("root/sudo access is required")
	}

	// Get current working directory
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Make an assumption about user's home directory
	userName := os.Getenv("SUDO_USER")
	homeDir := "/home/" + userName 

	// Check to see if current working directory contains home directory
	isHomeDir := strings.HasPrefix(pwd, homeDir)
	if isHomeDir != true {
		return nil, errors.New("script requires linux OS")
	}

	// Populate auth user info
	a := &authUser{
		name: userName,
		uid:  os.Getenv("SUDO_UID"),
		gid:  os.Getenv("SUDO_GID"),
		dir:  homeDir,
	}
	return a, nil
}
