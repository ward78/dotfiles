package main

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"strconv"
)

type system struct {
	user string
	uid  int
	gid  int
	path
}

type path struct {
	home     string
	dotfiles string
	backup   string
}

func NewSystem() (*system, error) {
	if runtime.GOOS != "linux" {
		return nil, errors.New("script requires linux OS")
	}
	admin, err := user.Current()
	if err != nil {
		return nil, err
	}
	if admin.Name != "root" {
		return nil, errors.New("root/sudo access is required")
	}
	dotfiles, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	user := os.Getenv("SUDO_USER")
	home := filepath.Join("/home", user)
	if !strings.HasPrefix(dotfiles, home) {
		return nil, errors.New("working directory is not a subdirectory of user's home directory")
	}
	uid, err := strconv.Atoi(os.Getenv("SUDO_UID"))
	if err != nil {
		return nil, err
	}
	gid, err := strconv.Atoi(os.Getenv("SUDO_GID"))
	if err != nil {
		return nil, err
	}
	s := &system{
		user: user,
		uid:  uid,
		gid:  gid,
		path: path{
			home:     home,
			dotfiles: dotfiles,
			backup:   filepath.Join(dotfiles, "backup"),
		},
	}
	return s, nil
}
