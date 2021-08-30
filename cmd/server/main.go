package main

import (
	"github.com/ward78/dotfiles/pkg/storage/json"
	"github.com/ward78/dotfiles/pkg/storage/memory"
	"github.com/ward78/dotfiles/pkg/system"
)

// StorageType defines available storage types
type Type int

const (
	// JSON will store data in JSON files saved on disk
	JSON Type = iota
	// Memory will store data in memory
	Memory
)

func main() {
	// set up storage
	storageType := JSON // this could be a flag; hardcoded here for simplicity
	var adder system.Service
	var lister system.Service
	var backuper system.Service
	switch storageType {
	case Memory:
		s := new(memory.Storage)
		adder = system.NewService(s)
		lister = system.NewService(s)
		backuper = system.NewService(s)
	case JSON:
		// error handling omitted for simplicity
		s, _ := json.NewStorage()
		adder = system.NewService(s)
		lister = system.NewService(s)
		backuper = system.NewService(s)
	}
}
