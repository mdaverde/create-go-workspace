package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var writeFilePermission os.FileMode = 0644

type createWorkspaceOptions struct {
	Silent bool
	DirEnv bool
	MainGo bool
	ReadMe bool
}

var write = func(path, contents string) error {
	return ioutil.WriteFile(path, []byte(contents), writeFilePermission)
}

var logf = func(silent bool, args ...interface{}) {
	if silent == true {
		return
	}
	first := args[0]
	if firstString, ok := first.(string); ok {
		fmt.Printf(firstString, args[1:]...)
	}
}

func writeFiles(parentDir, projectDir, projectName string, options *createWorkspaceOptions) (err error) {
	if options.DirEnv == true {
		filePath := path.Join(parentDir, ".envrc")
		err = write(filePath, envrc())
		if err != nil {
			return err
		}
		logf(options.Silent, "Created: %s\n", filePath)
	}

	if options.MainGo == true {
		filePath := path.Join(projectDir, "main.go")
		err = write(filePath, mainGo())
		if err != nil {
			return err
		}
		logf(options.Silent, "Created: %s\n", filePath)
	}

	if options.ReadMe == true {
		filePath := path.Join(projectDir, "README.md")
		err = write(filePath, readMe(projectName))
		if err != nil {
			return err
		}
		logf(options.Silent, "Created: %s\n", filePath)
	}

	return err
}

func createWorkspace(projectRepo string, options *createWorkspaceOptions) error {
	logf(options.Silent, "Creating %s workspace...\n", projectRepo)
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	_, projectName := path.Split(projectRepo)
	parentDir := path.Join(workingDir, projectName)
	projectDir := path.Join(parentDir, "src", projectRepo)

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

	err = writeFiles(parentDir, projectDir, projectName, options)

	if err != nil {
		return err
	}

	logf(options.Silent, "Done.")
	return nil
}
