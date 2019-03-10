package main

import (
	"os/exec"
	"path/filepath"
)

const (
	codeCmd = "code"
)

func diff(masterDir string, prDir string, path string) error {
	masterFile := filepath.Join(masterDir, path)
	prFile := filepath.Join(prDir, path)

	return exec.Command(codeCmd, "-d", masterFile, prFile).Run()
}
