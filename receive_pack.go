package main

import (
	"fmt"
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/file"
)

type CmdReceivePack struct {
	cmd

	Args struct {
		GitDir string `positional-arg-name:"git-dir" required:"true"`
	} `positional-args:"yes"`
}

func (CmdReceivePack) Usage() string {
	//TODO: git-receive-pack returns error code 129 if arguments are invalid.
	return fmt.Sprintf("usage: %s <git-dir>", os.Args[0])
}

func (c *CmdReceivePack) Execute(args []string) error {
	gitDir, err := filepath.Abs(c.Args.GitDir)
	if err != nil {
		return err
	}

	_, errOpen := git.PlainOpen(gitDir)
	if errOpen != nil {
		if errOpen == git.ErrRepositoryNotExists {
			_, errInit := git.PlainInit(gitDir, true)
			if errInit != nil {
				writeToLogFile("init error: " + errInit.Error())
			}
		} else {
			writeToLogFile("open error: " + errOpen.Error())
		}
	}

	if err := file.ServeReceivePack(gitDir); err != nil {
		fmt.Fprintln(os.Stderr, "ERR:", err)
		os.Exit(128)
	}

	return nil
}
