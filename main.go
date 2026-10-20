package main

import (
	"errors"
	"fmt"
	"path/filepath"
)

func parseArgs(args []string) (string, string, error) {
	if len(args) != 2 {
		return "", "", errors.New("usage: alias_master <alias> <command>")
	}

	return args[0], args[1], nil
}

func configPath(homeDir string, relPath string) string {
	return filepath.Join(homeDir, relPath)
}

func formatLine(alias string, command string) string {
	return fmt.Sprintf("alias %s='%s'\n", alias, command)
}

func writeToConfig(path string, str string) error {
	panic("todo")
}

func main() {

}
