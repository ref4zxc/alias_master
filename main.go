package main

import (
	"errors"
	"fmt"
	"os"
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

// TODO: разобрать случай с одинарными кавычками, например, grep 'x'
func formatLine(alias string, command string) string {
	return fmt.Sprintf("alias %s='%s'\n", alias, command)
}

func writeToConfig(path string, str string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("open %s: %w", path, err)
	}

	defer file.Close()

	if _, err := file.WriteString(str); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	return nil
}

func main() {
	args := os.Args[1:]
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	alias, command, err := parseArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	path := configPath(homedir, ".config/fish/config1.fish")
	formattedLine := formatLine(alias, command)

	if err := writeToConfig(path, formattedLine); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
