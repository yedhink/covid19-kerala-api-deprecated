package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

type Storage struct {
	Path string
}

func (s Storage) LocalPDFName() []string{
	/*
		Glob ignores file system errors such as I/O errors reading directories.
		The only possible returned error is ErrBadPattern, when pattern is malformed.
	*/
	files, err := filepath.Glob(s.Path+"*.pdf")
	if err != nil || len(files) == 0{
		fmt.Println(err)
		os.Exit(1)
	}
	return files
}