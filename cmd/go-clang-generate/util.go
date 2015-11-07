package main

import (
	"errors"
	"os"
	"os/exec"
	"syscall"
	"unicode"
)

func ExecToBuffer(cmd ...string) (out []byte, exitStatus int, err error) {
	c := exec.Command(cmd[0], cmd[1:]...)

	out, err = c.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				return out, status.ExitStatus(), err
			}
		}

		return out, 0, err
	}

	return out, 0, nil
}

func stat(filepath string) (os.FileInfo, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return fi, nil
}

var (
	errNotADirectory = errors.New("not a directory")
	errNotAFile      = errors.New("not a file")
)

func DirExists(path string) error {
	fi, err := stat(path)
	if err != nil {
		return err
	}

	if !fi.Mode().IsDir() {
		return errNotADirectory
	}

	return nil
}

func FileExists(filepath string) error {
	fi, err := stat(filepath)
	if err != nil {
		return err
	}

	if fi.Mode().IsDir() {
		return errNotAFile
	}

	return nil
}

func LowerFirstCharacter(s string) string {
	r := []rune(s)

	r[0] = unicode.ToLower(r[0])

	return string(r)
}

func UpperFirstCharacter(s string) string {
	r := []rune(s)

	r[0] = unicode.ToUpper(r[0])

	return string(r)
}

var goKeywordReplacements = map[string]string{
	"range": "r",
}

func ReplaceGoKeywords(s string) string {
	if r, ok := goKeywordReplacements[s]; ok {
		return r
	}

	return ""
}
