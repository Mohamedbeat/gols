package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"text/tabwriter"

	"github.com/fatih/color"
)

var (
	cblue        = color.New(color.FgBlue)
	cred         = color.New(color.FgRed)
	cgreen       = color.New(color.FgGreen)
	cyellow      = color.New(color.FgYellow)
	showHidden   bool
	showFullPath bool
)

func main() {
	flag.BoolVar(&showHidden, "a", false, "show hidden files.")
	flag.BoolVar(&showFullPath, "p", false, "show full path")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(cwd)

	e, err := os.ReadDir(cwd)
	if err != nil {
		panic(err)

	}
	total := len(e)
	if total > 100 {
		scanner := bufio.NewScanner(os.Stdin)
		cblue.Printf("show %d results? y/n ", total)
		scanner.Scan()
		fmt.Println()
		input := scanner.Text()
		if input == "y" {
			printEntries(e, &cwd)

			return
		} else {
			return
		}
	}
	printEntries(e, &cwd)

}

func printEntries(entries []os.DirEntry, cwd *string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	for _, e := range entries {
		name := e.Name()
		hidden := len(name) > 0 && name[0] == '.'
		if hidden && !showHidden {
			continue
		}
		if showFullPath {
			name = path.Join(*cwd, name)
		}
		isDir := e.IsDir()

		info, err := e.Info()
		if err != nil {
			fmt.Fprintf(w, "%s\t%v\n", name, isDir)
			continue
		}

		size := info.Size()
		modTime := info.ModTime().Format("02-01-2006 15:04:05")
		fileMode := info.Mode()

		// Format size with thousands separator
		sizeStr := formatSize(size)
		// Color formatting
		var coloredName string
		switch {
		case isDir && hidden:
			coloredName = cred.Sprint(name)
		case isDir:
			coloredName = cblue.Sprint(name)
		case hidden:
			coloredName = cyellow.Sprint(name)
		default:
			coloredName = cgreen.Sprint(name)
		}

		// Write aligned columns:
		// Permissions | Size | Name | Modification Time
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			fileMode,
			sizeStr,
			coloredName,
			modTime,
		)
	}
}

func formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d bytes", size)
	}

	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	units := []string{"KB", "MB", "GB", "TB", "PB", "EB"}
	if exp >= len(units) {
		exp = len(units) - 1
	}

	return fmt.Sprintf("%.1f %s", float64(size)/float64(div), units[exp])
}
