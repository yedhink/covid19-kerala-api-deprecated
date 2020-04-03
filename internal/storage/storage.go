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
	LocalFileExist bool
}


func (s Storage) Delete() {
	if s.LocalFileExist {
		var err = os.Remove(s.LocalFilePath)
		if err != nil {
			fmt.Printf("os.remove error\t: was not able to delete %s\nEXITING....\n",s.LocalFilePath)
		} else {
			fmt.Printf("Succesfully removed the local pdf %s\n",s.LocalFilePath)
		}
	}
}

// LocalPDFName retrieves the local pdf file from "data" dir
func (s *Storage) LocalPDFName() string{
	// Glob ignores file system errors such as I/O errors reading directories.
	// The only possible returned error is ErrBadPattern, when pattern is malformed.
	files, err := filepath.Glob(s.BasePath+"*.pdf")
	if err != nil {
		fmt.Printf("glob error : no local pdf file exists in %s\n",s.BasePath)
	}
	if len(files) == 0{
		fmt.Println("since no pdf file exists - directly download the latest file")
		s.LocalFilePath = "data/01-04-2020.pdf"
	} else {
		s.LocalFilePath = files[0]
		s.LocalFileExist = true
	}
	return s.LocalFilePath
}