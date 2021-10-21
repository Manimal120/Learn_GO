package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func fileLine(path string) int {
	count := 0
	file, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		ctx, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		txt := strings.TrimSpace(string(ctx))
		if txt == "" || strings.HasPrefix(txt, "//") {
			continue
		}
		count++
	}
	return count
}

func main() {
	dir := "."
	total := 0
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			total += fileLine(path)
		}
		return nil
	})
	fmt.Println(total)
}
