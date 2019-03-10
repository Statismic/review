package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	gitCmd = "git"
)

// repoClone clones the repository given in the url and ref.
// repoClone will return the path to directory where the
// repository is located.
func repoClone(url string, ref string) (string, error) {
	// Create a tmp dir
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		os.RemoveAll(dir)
		return "", err
	}

	// use curDir to go back to where it was
	curDir, err := os.Getwd()
	if err != nil {
		os.RemoveAll(dir)
		return "", err
	}
	defer os.Chdir(curDir)

	if err = os.Chdir(dir); err != nil {
		os.RemoveAll(dir)
		return "", err
	}

	cmd := exec.Command(gitCmd, "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		os.RemoveAll(dir)
		return "", err
	}

	cmd = exec.Command(gitCmd, "pull", url, ref)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		os.RemoveAll(dir)
		return "", err
	}

	return dir, nil
}

func repoCloneMaster(url string) (string, error) {
	return repoClone(url, "refs/heads/master")
}

func repoClonePR(url string, id uint64) (string, error) {
	return repoClone(url, fmt.Sprintf("pull/%d/head", id))
}
