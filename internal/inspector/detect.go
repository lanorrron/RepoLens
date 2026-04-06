package inspector

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func DetectLanguage(root string) string {
	counts := map[string]int{}

	ignoreDirs := []string{"node_modules", "dist", ".git"}

	info, err := os.Stat(root)
	if err != nil || !info.IsDir() {
		fmt.Printf("Invalid path: %v\n", root)
		return "unknown"
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			for _, d := range ignoreDirs {
				if info.Name() == d {
					return filepath.SkipDir
				}
			}
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		counts[ext]++

		return nil
	})

	fmt.Println("Extension counts:", counts)

	if counts[".ts"]+counts[".tsx"] > 0 {
		return "typescript"
	}
	if counts[".js"]+counts[".jsx"] > 0 {
		return "javascript"
	}
	if counts[".go"] > 0 {
		return "go"
	}

	return "unknown"
}
