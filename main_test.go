package main

import "testing"

func TestFormatLine(t *testing.T) {
	tests := []struct {
		name    string
		alias   string
		command string
		want    string
	}{
		{"Простая команда", "cl", "clear", "alias cl='clear'\n"},
		{"Команда с флагами", "ll", "ls -la", "alias ll='ls -la'\n"},
		// TODO: разобрать случай с одинарными кавычка, например, grep 'x'
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
