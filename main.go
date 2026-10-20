package main

import (
	"fmt"
	"path/filepath"
)

func parseArgs(args []string) (string, string, error) {
	panic("todo")
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
