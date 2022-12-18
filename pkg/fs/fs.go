package fs

import (
	"errors"
	"golang.org/x/sys/unix"
	"os"
)

func CheckFolder(folder string) error {
	// check that upload folder exists
	if _, err := os.Stat(folder); err != nil {
		return err
	}

	// check is folder writable and readable
	if unix.Access(folder, unix.W_OK|unix.R_OK) != nil {
		return errors.New("can't read or write upload folder")
	}

	return nil
}

func CreateFolder(folder string) error {
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func FileExists(filepath string) bool {
	info, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
