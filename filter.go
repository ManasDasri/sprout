package main

import (
	"path/filepath"
	"strings"
)

// defaultProjectIgnores are common development artifacts that --project
// filters out automatically, without the user having to type them every time.
var defaultProjectIgnores = []string{
	"node_modules",
	".git",
	"dist",
	"build",
	".cache",
	"__pycache__",
	".venv",
	"venv",
	".idea",
	".vscode",
	"target",
	"bin",
	"obj",
	"coverage",
	".next",
	".nuxt",
	".pytest_cache",
	".mypy_cache",
	".DS_Store",
	"*.egg-info",
}

// buildIgnoreList merges the user-supplied --ignore patterns with either
// the --project defaults, or a minimal baseline (just .git) otherwise.
func buildIgnoreList(userIgnore string, project bool) []string {
	var patterns []string

	if project {
		patterns = append(patterns, defaultProjectIgnores...)
	} else {
		patterns = append(patterns, ".git")
	}

	if userIgnore != "" {
		for _, p := range strings.Split(userIgnore, ",") {
			p = strings.TrimSpace(p)
			if p != "" {
				patterns = append(patterns, p)
			}
		}
	}

	return patterns
}

// isIgnored checks a bare file/dir name against glob-style patterns.
func isIgnored(name string, patterns []string) bool {
	for _, pattern := range patterns {
		if matched, _ := filepath.Match(pattern, name); matched {
			return true
		}
	}
	return false
}
