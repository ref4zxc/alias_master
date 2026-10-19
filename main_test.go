package main

import "testing"

func TestFormatLine(t *testing.T) {
	got := formatLine("cl", "clear")
	want := "alias cl='clear'\n"

	if got != want {
		t.Errorf("formatLine() = %q, want %q", got, want)
	}
}
