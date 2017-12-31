package main

import (
	"fmt"
	"os"
	"path"
	"io/ioutil"
)

type createWorkspaceOptions struct {
	Silent bool
	DirEnv bool
	MainGo bool
}

func logf(silent bool, args... interface{}) {
	if silent == true {
		return
	}
	first := args[0]
	if firstString, ok := first.(string); ok {
		fmt.Printf(firstString, args[1:]...)
	}
}

func writeFiles(parentDir, projectDir string, options *createWorkspaceOptions) (err error) {
	if options.DirEnv == true {
		filePath := path.Join(parentDir, ".envrc")
		err = ioutil.WriteFile(filePath, []byte(envrc()), os.ModePerm)
		if err != nil {
			return err
		}
		logf(options.Silent, "Created: %s\n", filePath)
	}

	if options.MainGo == true {
		filePath := path.Join(projectDir, "main.go")
		err = ioutil.WriteFile(filePath, []byte(mainGo()), os.ModePerm)
		if err != nil {
			return err
		}
		logf(options.Silent, "Created: %s\n", filePath)
	}

	return err
}

func createWorkspace(projectName string, options *createWorkspaceOptions) error {
	logf(options.Silent, "Creating %s workspace...\n", projectName)
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	_, postFinalSlash := path.Split(projectName)
	parentDir := path.Join(workingDir, postFinalSlash)
	projectDir := path.Join(parentDir, "src", projectName)

	workspaceDirPaths := []string{
		parentDir,
		path.Join(parentDir, "src"),
		projectDir,
		path.Join(parentDir, "bin"),
		path.Join(parentDir, "pkg"),
	}

	for _, dirPath := range workspaceDirPaths {
		if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}
		logf(options.Silent, "Created: %s\n", dirPath)
	}

	err = writeFiles(parentDir, projectDir, options)

	if err != nil {
		return err
	}

	logf(options.Silent, "Done.")
	return nil
}
