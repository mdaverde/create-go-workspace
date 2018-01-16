package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var lastFilesWritten [][]string
	var lastLogged []string

	write = func(path, contents string) error {
		lastFilesWritten = append(lastFilesWritten, []string{path, contents})
		return nil
	}

	logf = func(silent bool, args ...interface{}) {
		if !silent {
			first := args[0]
			if firstString, ok := first.(string); ok {
				lastLogged = append(lastLogged, fmt.Sprintf(firstString, args[1:]...))
			}
		}
	}

	afterEachTest := func() {
		lastFilesWritten = nil
		lastLogged = nil
	}

	t.Run("Write all files", func(t *testing.T) {
		err := writeFiles("parentDir", "projectDir", "projectName", &createWorkspaceOptions{
			Silent: true,
			DirEnv: true,
			MainGo: true,
			ReadMe: true,
		})

		if err != nil {
			t.Error(err)
		}

		if len(lastLogged) > 0 {
			t.Errorf("log messages were generated when silent was set to true: %v", lastLogged)
		}

		if lastFilesWritten[0][0] != "parentDir/.envrc" {
			t.Error(".envrc was not written")
		}

		if lastFilesWritten[1][0] != "projectDir/main.go" {
			t.Error("main.go was not written")
		}

		if lastFilesWritten[2][0] != "projectDir/README.md" {
			t.Error("README.md was not written")
		}

		afterEachTest()
	})

	t.Run("Write partial", func(t *testing.T) {
		err := writeFiles("parentDir", "projectDir", "projectName", &createWorkspaceOptions{
			Silent: true,
			DirEnv: true,
			MainGo: false,
			ReadMe: true,
		})

		if err != nil {
			t.Error(err)
		}

		if len(lastLogged) > 0 {
			t.Errorf("log messages were generated when silent was set to true: %v", lastLogged)
		}

		if lastFilesWritten[0][0] != "parentDir/.envrc" {
			t.Error(".envrc was not written")
		}

		if lastFilesWritten[1][0] != "projectDir/README.md" {
			t.Error("README.md was not written")
		}

		afterEachTest()
	})

	t.Run("Silent false", func(t *testing.T) {
		err := writeFiles("parentDir", "projectDir", "projectName", &createWorkspaceOptions{
			Silent: false,
			DirEnv: true,
			MainGo: true,
			ReadMe: true,
		})

		if err != nil {
			t.Error(err)
		}

		if len(lastLogged) < 1 {
			t.Error("log messages were not generated when silent was set to false")
		}

		afterEachTest()
	})
}
