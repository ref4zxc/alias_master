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
