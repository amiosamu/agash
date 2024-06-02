package main

import (
	"fmt"
	"os"
	"path/filepath"
)


func printTree(path string, ident string){ 
	dir, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer dir.Close()

	// Read directory entries. -1, read all entries at once
	entries, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Loop over the directory entries
	for i, entry := range entries {

		if entry.IsDir() && entry.Name()[0] == '.' {
			continue
		}

		// Create the proper indent based on the current level
		prefix := ident
		if i == len(entries) -1 {
			prefix += "└── "
		} else {
			prefix += "├── "
		}

		fmt.Println(prefix + entry.Name())
		
		// If the entry is a directory, recursively, print its contents
		if entry.IsDir() {
			newIdent := ident
			if i == len(entries) - 1 {
				newIdent += "    "
			} else  {
				newIdent += "│   "
			}

			printTree(filepath.Join(path, entry.Name()), newIdent)
		}
	}

}
func main(){

	if len(os.Args) < 2 {
		fmt.Println("Usage: agash [directory]")
		return
	}


	root := os.Args[1]

	fmt.Println(root)

	printTree(root, "")

}