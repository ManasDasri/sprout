package main

import (
	"os"
	"path/filepath"
)

// ProjectInfo is what --project detects about the root directory from
// well-known marker files. This is intentionally simple pattern matching,
// not the fuller --smart heuristics planned for a later release.
type ProjectInfo struct {
	Language       string
	PackageManager string
}

var projectMarkers = []struct {
	File           string
	Language       string
	PackageManager string
}{
	{"go.mod", "Go", "go modules"},
	{"package.json", "JavaScript/TypeScript", "npm"},
	{"pyproject.toml", "Python", "poetry/pip"},
	{"requirements.txt", "Python", "pip"},
	{"Cargo.toml", "Rust", "cargo"},
	{"pom.xml", "Java", "maven"},
	{"build.gradle", "Java/Kotlin", "gradle"},
	{"Gemfile", "Ruby", "bundler"},
	{"composer.json", "PHP", "composer"},
}

// DetectProject looks for the first recognized manifest file in root.
// Returns nil if nothing matched.
func DetectProject(root string) *ProjectInfo {
	for _, marker := range projectMarkers {
		if _, err := os.Stat(filepath.Join(root, marker.File)); err == nil {
			return &ProjectInfo{Language: marker.Language, PackageManager: marker.PackageManager}
		}
	}
	return nil
}
