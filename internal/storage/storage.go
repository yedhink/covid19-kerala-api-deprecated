package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

type Storage struct {
	BasePath string
	LocalFilePath string
	RemoteFileName string
}


func (s Storage) Download() {
	var err = os.Remove(s.LocalFilePath)
	if err != nil {
		fmt.Printf("os.remove error\t: was not able to delete %s\nEXITING....\n",s.LocalFilePath)
		os.Exit(0)
	} else {
		fmt.Printf("Succesfully removed the local pdf %s\n",s.LocalFilePath)
	}
}

// LocalPDFName retrieves the local pdf file from "data" dir
func (s *Storage) LocalPDFName() []string{
	// Glob ignores file system errors such as I/O errors reading directories.
	// The only possible returned error is ErrBadPattern, when pattern is malformed.
	files, err := filepath.Glob(s.BasePath+"*.pdf")
	if err != nil || len(files) == 0{
		fmt.Println(err)
		os.Exit(1)
	}
	s.LocalFilePath = files[0]
	return files
}