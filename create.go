package main

import (
	"fmt"
	"os"
	"path"
)

func logf(args... interface{}) {
	first := args[0]
	if firstString, ok := first.(string); ok {
		fmt.Printf(firstString, args[1:]...)
	}
}

func workspaceDirPaths(rootDir, name string) []string {
	_, postFinalSlash := path.Split(name)
	parentDir := path.Join(rootDir, postFinalSlash)
	return []string{
		parentDir,
		path.Join(parentDir, "src"),
		path.Join(parentDir, "src", name),
		path.Join(parentDir, "bin"),
		path.Join(parentDir, "pkg"),
	}
}

func createWorkspace(name string) error {
	logf("Creating %s workspace...\n", name)
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}
	for _, dirPath := range workspaceDirPaths(workingDir, name) {
		if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}
		logf("Created: %s\n", dirPath)
	}

	logf("Done.")
	return nil
}
