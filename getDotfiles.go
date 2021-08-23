package main

import (
	"os"
	"path/filepath"
)

func getDotfiles(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	var dotfiles []string
	for _, file := range files {
		filename := file.Name()
		if filename[0:1] == "." {
			dotfiles = append(dotfiles, filepath.Join(dir, file.Name()))
		}
	}
	return dotfiles, nil
}
