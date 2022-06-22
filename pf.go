package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func unixShellFromDos() bool {
	shell := os.Getenv("SHELL")

	found, _ := regexp.MatchString("sh.exe", shell)
	return found
}

func dosToUnix(path string) string {
	path = fmt.Sprintf("/%s%s", strings.ToLower(string(path[0])), strings.ReplaceAll(path[2:], "\\", "/"))
	path = strings.ReplaceAll(path, " ", "\\ ")
	path = strings.ReplaceAll(path, "(", "\\(")
	path = strings.ReplaceAll(path, ")", "\\)")
	return path
}

func printPath(paths []string) {

	sep := PATH_ENV_SEPARATOR

	if unixShellFromDos() && paths[0][1] == ':' {
		sep = ":"

		for i, path := range paths {

			paths[i] = dosToUnix(path)
		}
	}

	fmt.Println("PATH=" + strings.Join(paths, sep))
}

func pathParts() []string {
	startingPath := os.Getenv("PATH")
	pathParts := strings.Split(startingPath, PATH_ENV_SEPARATOR)

	return pathParts
}

func first(pathsToMove []string) {
	paths := pathParts()

	foundPaths := []string{}

	for _, toMove := range pathsToMove {
		for i := 0; i < len(paths); i++ {
			found, _ := regexp.MatchString(toMove, paths[i])
			if found {
				foundPaths = append(foundPaths, paths[i])
				paths = append(paths[:i], paths[i+1:]...)
				i--
				break
			}
		}
	}

	paths = append(foundPaths, paths...)

	printPath(paths)
}

func delete(pathsToDelete []string) {
	paths := pathParts()

	deleteCount := 0

	for _, toMove := range pathsToDelete {
		for i := 0; i < len(paths); i++ {
			found, _ := regexp.MatchString(toMove, paths[i])
			if found {
				paths = append(paths[:i], paths[i+1:]...)
				deleteCount++
				i--
				break
			}
		}
	}

	printPath(paths)
}

func search(pathsToFind []string) {
	paths := pathParts()

	for i := 0; i < len(paths); i++ {
		for _, toFind := range pathsToFind {
			found, _ := regexp.MatchString(toFind, paths[i])
			if found {
				fmt.Println(paths[i])
			}
		}
	}

}

func print() {
	paths := pathParts()
	for _, path := range paths {
		fmt.Println(path)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pf [command]")
		fmt.Println("")
		fmt.Println("  Commands:")
		fmt.Println("    first {pathRegEx }+")
		fmt.Println("    delete {pathRegEx }+")
		fmt.Println("    search {pathRegEx }+")
		fmt.Println("    print")
		os.Exit(0)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch {
	case command == "first":
		first(args)
	case command == "search":
		search(args)
	case command == "print":
		print()
	case command == "delete":
		delete(args)
	}
}
