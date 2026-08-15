package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// Stats is the aggregate result of walking a tree for --stats.
type Stats struct {
	Files       int
	Directories int
	TotalSize   int64
	Languages   map[string]int
}

// languageNames maps file extensions to a human-readable language label.
// Extensions not listed here fall back to their uppercased form (e.g. .toml -> TOML).
var languageNames = map[string]string{
	".go":   "Go",
	".py":   "Python",
	".js":   "JavaScript",
	".jsx":  "JavaScript",
	".ts":   "TypeScript",
	".tsx":  "TypeScript",
	".css":  "CSS",
	".scss": "SCSS",
	".md":   "Markdown",
	".json": "JSON",
	".yaml": "YAML",
	".yml":  "YAML",
	".html": "HTML",
	".java": "Java",
	".c":    "C",
	".h":    "C Header",
	".cpp":  "C++",
	".rs":   "Rust",
	".rb":   "Ruby",
	".php":  "PHP",
	".sh":   "Shell",
	".sql":  "SQL",
}

func collectStats(node *Node, s *Stats) {
	for _, child := range node.Children {
		if child.IsDir {
			s.Directories++
			collectStats(child, s)
			continue
		}

		s.Files++
		s.TotalSize += child.Size

		ext := strings.ToLower(filepath.Ext(child.Name))
		if ext == "" {
			continue
		}

		lang, ok := languageNames[ext]
		if !ok {
			lang = strings.ToUpper(strings.TrimPrefix(ext, "."))
		}
		s.Languages[lang]++
	}
}

func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// PrintStats renders the my-project/ summary shown in the README's --stats example.
func PrintStats(root *Node, path string) {
	s := &Stats{Languages: make(map[string]int)}
	collectStats(root, s)

	fmt.Printf("%s/\n", strings.TrimSuffix(path, "/"))
	fmt.Printf("Files:          %d\n", s.Files)
	fmt.Printf("Directories:    %d\n", s.Directories)
	fmt.Printf("Total size:     %s\n", formatSize(s.TotalSize))

	if len(s.Languages) == 0 {
		return
	}

	type langCount struct {
		Name  string
		Count int
	}
	var langs []langCount
	for name, count := range s.Languages {
		langs = append(langs, langCount{name, count})
	}
	sort.Slice(langs, func(i, j int) bool {
		if langs[i].Count != langs[j].Count {
			return langs[i].Count > langs[j].Count
		}
		return langs[i].Name < langs[j].Name
	})

	fmt.Println("Languages:")
	for _, l := range langs {
		fmt.Printf("  %-14s %d files\n", l.Name, l.Count)
	}
}
