package main

import (
	"os"
	"errors"
	"io"
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

func copyFile(src string, dstPath, dstName string) error {
	if !dirExists(dstPath) {
		return errors.New("Target directory "+dstPath+" does not exist.")
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
        return
    }
    var closeErr error
    defer func() {
        closeErr = out.Close()
    }()

    // copy file content
    if _, err = io.Copy(out, in); err != nil {
        return
    }

    // commit file contents
    if err := out.Sync(); err != nil {
    	return err
    }

    return closeErr
}