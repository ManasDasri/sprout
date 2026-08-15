package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	hidden := flag.Bool("hidden", false, "show hidden files and directories")
	depth := flag.Int("depth", -1, "limit directory depth (-1 for unlimited)")
	ignore := flag.String("ignore", "", "comma-separated list of names/patterns to ignore")
	project := flag.Bool("project", false, "ignore common build/dependency artifacts (node_modules, .git, dist, ...)")
	stats := flag.Bool("stats", false, "show project statistics instead of the tree")

	flag.Parse()

	path := "."
	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}

	info, err := os.Stat(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "sprout:", err)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "sprout: %s is not a directory\n", path)
		os.Exit(1)
	}

	opts := Options{
		ShowHidden: *hidden,
		MaxDepth:   *depth,
		Ignore:     buildIgnoreList(*ignore, *project),
	}

	root, err := BuildTree(path, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "sprout:", err)
		os.Exit(1)
	}

	if *project {
		if pi := DetectProject(path); pi != nil {
			fmt.Printf("Detected project:\n  Language:        %s\n  Package manager: %s\n\n", pi.Language, pi.PackageManager)
		}
	}

	if *stats {
		PrintStats(root, path)
		return
	}

	fmt.Println(path)
	PrintTree(root, "")
}
