package main

import "testing"

func TestIsIgnored(t *testing.T) {
	patterns := []string{"node_modules", ".git", "*.log"}

	cases := map[string]bool{
		"node_modules": true,
		".git":         true,
		"debug.log":    true,
		"main.go":      false,
		"src":          false,
	}

	for name, want := range cases {
		if got := isIgnored(name, patterns); got != want {
			t.Errorf("isIgnored(%q) = %v, want %v", name, got, want)
		}
	}
}

func TestBuildIgnoreListDefault(t *testing.T) {
	list := buildIgnoreList("foo,bar", false)
	if len(list) != 3 { // .git + foo + bar
		t.Errorf("expected 3 patterns, got %d: %v", len(list), list)
	}
}

func TestBuildIgnoreListProject(t *testing.T) {
	list := buildIgnoreList("", true)

	found := false
	for _, p := range list {
		if p == "node_modules" {
			found = true
		}
	}
	if !found {
		t.Error("expected project ignore list to include node_modules")
	}
}

func TestBuildIgnoreListTrimsWhitespace(t *testing.T) {
	list := buildIgnoreList(" foo , bar ", false)

	want := map[string]bool{".git": true, "foo": true, "bar": true}
	for _, p := range list {
		if !want[p] {
			t.Errorf("unexpected pattern %q in ignore list", p)
		}
		delete(want, p)
	}
	if len(want) != 0 {
		t.Errorf("missing expected patterns: %v", want)
	}
}
