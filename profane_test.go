package profane_test

import (
	"testing"

	"github.com/alextanhongpin/profane"
)

func TestMatch(t *testing.T) {
	profane := profane.New()

	if profane.Has("hello") {
		t.Errorf("word has been added")
	}

	profane.Add("HELLO")
	profane.Add("hello")
	profane.Add("hElLo")

	tests := []struct {
		given string
		when  string
		then  bool
	}{
		{"a lowercase input", "hello world", true},
		{"an uppercase input", "HELLO world", true},
		{"prefix of word", "hellos", false},
		{"suffix of word", "ahello", false},
		{"empty string", "", false},
		{"test leet", "h3ll0", true},
		{"test capitalization", "H3LL0", true},
	}
	for _, tt := range tests {
		t.Run(tt.given, func(t *testing.T) {
			if got := profane.Match(tt.when); got != tt.then {
				t.Errorf("given %s, want %t, got %t", tt.when, tt.then, got)
			} else {
				t.Logf("%s => %t", tt.when, tt.then)
			}
		})
	}
}

func TestReplaceGarbled(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"", ""},
		{" ", " "},
		{"hello world", "$@!#% world"},
		{"HELLO world", "$@!#% world"},
	}
	profane := profane.New()
	profane.Add("hello")
	for _, tt := range tests {
		got := profane.ReplaceGarbled(tt.in)
		if got != tt.out {
			t.Fatalf("given %s, expected %s, got %s", tt.in, tt.out, got)
		}
		t.Logf("in: %q, out: %q", tt.in, got)
	}
}

func TestReplaceStars(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"", ""},
		{" ", " "},
		{"six", "s*x"},
		{"SIX", "S*X"},
		{"hello world", "h***o world"},
		{"HELLO WORLD", "H***O WORLD"},
		{"HELLO HELLO", "H***O H***O"},
	}

	profane := profane.New()
	profane.Add("Six", "hello")
	for _, tt := range tests {
		got := profane.ReplaceStars(tt.in)
		if got != tt.out {
			t.Fatalf("given %s, expected %s, got %s", tt.in, tt.out, got)
		}
		t.Logf("in: %q, out: %q", tt.in, got)
	}
}

func TestReplaceVowels(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"", ""},
		{" ", " "},
		{"hello world", "h*ll* world"},
		{"HELLO WORLD", "H*LL* WORLD"},
		{"H3LL0 WORLD", "H*LL* WORLD"},
		{"HELLO HELLO", "H*LL* H*LL*"},
		{"H3LL0 HELLO", "H*LL* H*LL*"},
	}
	profane := profane.New()
	profane.Add("HELLO")
	for _, tt := range tests {
		got := profane.ReplaceVowels(tt.in)
		if got != tt.out {
			t.Fatalf("given %s, expected %s, got %s", tt.in, tt.out, got)
		}
		t.Logf("in: %q, out: %q", tt.in, got)
	}
}

func TestReplaceCustom(t *testing.T) {
	tests := []struct {
		in     string
		out    string
		custom string
	}{
		{"", "", ""},
		{" ", " ", ""},
		{"hello world", "[CENSORED] world", "[CENSORED]"},
		{"h3ll0 world", "[CENSORED] world", "[CENSORED]"},
		{"H3LL0 world", "[CENSORED] world", "[CENSORED]"},
		{"hello HELLO", "[CENSORED] [CENSORED]", "[CENSORED]"},
	}
	profane := profane.New()
	profane.Add("HELLO")
	for _, tt := range tests {
		got := profane.ReplaceCustom(tt.in, tt.custom)
		if got != tt.out {
			t.Fatalf("given %s, expected %s, got %s", tt.in, tt.out, got)
		}
		t.Logf("in: %q, out: %q", tt.in, got)
	}
}
