package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// check if path exists and is directory
func dirExists(s string) bool {
	stat, err := os.Stat(s)
	if os.IsNotExist(err) {
		return false
	}
	return stat.IsDir()
}

// copy the contents of the source directory into the 
// destination directory
func copyDirContents(src, dst string) error {
	if !dirExists(src) || !dirExists(dst) {
		return errors.New("Source or target directory does not exists.")
	}

	// list contents of source directory
	contents, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, fInfo := range contents {
		// for subdirectories, pass the copy call down
		if fInfo.IsDir() {
			dstSubdir := filepath.Join(dst, fInfo.Name())
			srcSubdir := filepath.Join(src, fInfo.Name())
			if err := os.Mkdir(dstSubdir, 0777); err != nil {
				return err
			}
			if err := copyDirContents(srcSubdir, dstSubdir); err != nil {
				return err
			}
		} else {
			// copy file to destination directory
			srcFile := filepath.Join(src, fInfo.Name())
			copyFile(srcFile, dst, fInfo.Name())
		}
	}

	return nil
}

// copy file at src path to file at dstPath with dstName
func copyFile(src string, dstPath, dstName string) error {
	if !dirExists(dstPath) {
		return errors.New("Target directory " + dstPath + " does not exist.")
	}

	// open source file
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	// create out file
	out, err := os.Create(filepath.Join(dstPath, dstName))
	if err != nil {
		return err
	}
	var closeErr error
	defer func() {
		closeErr = out.Close()
	}()

	// copy file content
	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	// commit file contents
	if err := out.Sync(); err != nil {
		return err
	}

	return closeErr
}
