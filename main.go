package main

import (
	"fmt"
	"io/fs"
	"os"
	// "path/filepath"
	"sort"
	"time"

	"github.com/fatih/color"
)

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
	Mode    fs.FileMode
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-ls <directory>")
		os.Exit(1)
	}

	dir := os.Args[1]
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	var fileInfos []FileInfo
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			continue
		}
		fileInfos = append(fileInfos, FileInfo{
			Name:    file.Name(),
			Size:    info.Size(),
			ModTime: info.ModTime(),
			Mode:    info.Mode(),
		})
	}

	// Sort by name (you can add more sorting options)
	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].Name < fileInfos[j].Name
	})

	for _, fileInfo := range fileInfos {
		printFileInfo(fileInfo)
	}
}

func printFileInfo(fileInfo FileInfo) {
	var colorFunc func(string, ...interface{}) string
	switch {
	case fileInfo.Mode.IsDir():
		colorFunc = color.New(color.FgBlue).SprintfFunc()
	case fileInfo.Mode.IsRegular():
		colorFunc = color.New(color.FgWhite).SprintfFunc()
	case fileInfo.Mode&0111 != 0:
		colorFunc = color.New(color.FgGreen).SprintfFunc()
	default:
		colorFunc = color.New(color.FgYellow).SprintfFunc()
	}

	fmt.Printf("%s %10d %s %s\n",
		fileInfo.Mode.String(),
		fileInfo.Size,
		fileInfo.ModTime.Format(time.RFC1123),
		colorFunc(fileInfo.Name),
	)
}
