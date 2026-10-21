package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFormatLine(t *testing.T) {
	tests := []struct {
		name    string
		alias   string
		command string
		want    string
	}{
		{"Простая команда", "cl", "clear", "alias cl='clear'\n"},
		{"Команда с флагами", "ll", "ls -la", "alias ll='ls -la'\n"},
		// TODO: разобрать случай с одинарными кавычками, например, grep 'x'
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatLine(tt.alias, tt.command)
			if got != tt.want {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}

func TestConfigPath(t *testing.T) {
	tests := []struct {
		name    string
		homeDir string
		relPath string
		want    string
	}{
		{"Без слэша с обеих сторон", "/home/user", ".config/fish/config1.fish", "/home/user/.config/fish/config1.fish"},
		{"Слэш в начале relPath", "/home/user", "/.config/fish/config1.fish", "/home/user/.config/fish/config1.fish"},
		{"Слэш в конце homeDir", "/home/user/", ".config/fish/config1.fish", "/home/user/.config/fish/config1.fish"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := configPath(tt.homeDir, tt.relPath)
			if got != tt.want {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantAlias   string
		wantCommand string
		wantErr     bool
	}{
		{"Передано ноль аргументов", []string{}, "", "", true},
		{"Передан один аргумент", []string{"ll"}, "", "", true},
		{"Передано два аргумента", []string{"ll", "ls -la"}, "ll", "ls -la", false},
		{"Передано больше двух аргументов", []string{"ll", "ls", "-la"}, "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAlias, gotCommand, err := parseArgs(tt.args)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if gotAlias != tt.wantAlias {
				t.Errorf("Alias: want %q, got %q", tt.wantAlias, gotAlias)
			}

			if gotCommand != tt.wantCommand {
				t.Errorf("Command: want %q, got %q", tt.wantCommand, gotCommand)
			}
		})
	}
}

func TestWriteToConfigAppend(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "config1.fish")
	first := "alias ll='ls -la'\n"
	second := "alias cl='clear'\n"
	want := first + second

	if err := writeToConfig(path, first); err != nil {
		t.Fatalf("first write: %v", err)
	}

	if err := writeToConfig(path, second); err != nil {
		t.Fatalf("second write: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read back: %v", err)
	}

	if got := string(data); got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestWriteToConfigMissingDir(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "missing_dir", "config1.fish")
	alias := "alias ll='ls -la'\n"

	if err := writeToConfig(path, alias); err == nil {
		t.Fatal("expected error, got nil")
	}
}
