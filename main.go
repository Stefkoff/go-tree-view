package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ByFileType []fs.DirEntry

func (a ByFileType) Len() int {
	return len(a)
}

func (a ByFileType) Less(i, j int) bool {
	if a[i].IsDir() && !!a[j].IsDir() {
		return true
	} else if !a[i].IsDir() && a[j].IsDir() {
		return false
	}

	return false
}

func (a ByFileType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func PrintTail(deep int) {
	for d := 0; d < deep; d++ {
		if d == 0 {
			fmt.Print("|    ")
		} else if d > 0 && deep%d == 0 {
			fmt.Print("|   ")
		} else if d > deep-4 {
			fmt.Print("|   ")
		}
	}
}

func ListDirContent(dirName string, deep int, maxDepth int, showHiddenFile bool) {
	if maxDepth > 0 && deep > maxDepth {
		return
	}
	PrintTail(deep)
	dir, err := os.ReadDir(dirName)
	check(err)
	dirNameToShow := dirName

	if deep > 0 {
		dirNameToShow = filepath.Base(dirNameToShow)
	}

	fmt.Printf("\033[34m%s\n\033[0m", dirNameToShow)

	sort.Sort(ByFileType(dir))
	for _, entry := range dir {
		if entry.IsDir() {
			ListDirContent(path.Join(dirName, entry.Name()), deep+1, maxDepth, showHiddenFile)
		} else {
			PrintTail(deep)
			fmt.Printf("|---%s\n", entry.Name())
		}
	}
}

func PrintHelp() {
	fmt.Println("Usage: go-tree-view [options] path/name")
	fmt.Println("Options:")
	fmt.Println("\t-d: Maximum deep; Default 0 which means no deep")
	fmt.Println("\t-h: Show hidden files. By Default they are not shown")
	fmt.Println("\t--help: Show this help")
}

func CheckForValidPath(pathName string) (string, bool) {
	curPath, err := os.Getwd()

	var pathSeparator string = "/"

	if err != nil {
		return "", false
	}
	if strings.Compare(pathName, ".") == 0 {
		pathName, err = filepath.Abs(curPath)

		if err != nil {
			return "", false
		}
	} else if strings.Contains(pathName, ".."+string(os.PathSeparator)) {
		pathName, err = filepath.Abs(curPath + pathSeparator + pathName)

		if err != nil {
			return "", false
		}
	}
	if !CheckDirExists(pathName) {
		return "", false
	}
	return pathName, true
}

func CheckDirExists(pathName string) bool {
	_, err := os.ReadDir(pathName)

	return err == nil
}

func main() {

	args := os.Args[1:]
	argc := len(args)
	var (
		maxDeep         int = 0
		pathName        string
		showHiddenFiles bool
	)

	if argc < 1 {
		PrintHelp()
	} else if argc == 1 {
		absolutPath, exists := CheckForValidPath(args[0])

		if exists == false {
			fmt.Println("Invalid path provided!")
			PrintHelp()
			return
		}

		pathName = absolutPath

	} else {
		absolutePath, exists := CheckForValidPath(args[argc-1])

		if exists == false {
			fmt.Println("Invalid path provided")
			PrintHelp()
			return
		}
		pathName = absolutePath
		for i, arg := range args {
			if arg == "-d" && len(args)-1 > i+1 {
				mDeep, err := strconv.Atoi(args[i+1])

				if err != nil {
					continue
				}

				maxDeep = mDeep
				continue
			} else if arg == "-h" {
				showHiddenFiles = true
				continue
			} else if arg == "--help" {
				PrintHelp()
				return
			}
		}
	}

	ListDirContent(pathName, 0, maxDeep, showHiddenFiles)
	//ListDirContent("C:\\Users\\stefkoff\\CLionProjects", 0)
}
