package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func readDirectory(path string, prefix string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var visibleEntries []os.DirEntry
	for _, entry := range entries {
		if entry.Name() != ".git" {
			visibleEntries = append(visibleEntries, entry)
		}
	}

	for i, entry := range visibleEntries {
		connector := "├── "
		if i == len(visibleEntries)-1 {
			connector = "└── "
		}
		fmt.Println(prefix + connector + entry.Name())
		if entry.IsDir() {
			childPath := filepath.Join(path, entry.Name())
			childPrefix := prefix + "│   "
			if i == len(visibleEntries)-1 {
				childPrefix = prefix + "    "
			}
			readDirectory(childPath, childPrefix)
		}
	}
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	fmt.Println(path)
	readDirectory(path, "")
}
