package main

import (
	"bytes"
	"fmt"
	"os"
)

func dirTree(f *bytes.Buffer, path string, printFiles bool) error {
	printNextFence := []int{0}
	return dirTreeRec(f, path, printFiles, 0, printNextFence)
}

func dirTreeRec(f *bytes.Buffer, path string, printFiles bool, level int, printNextFence []int) error {
	files, _ := os.ReadDir(path)

	if len(files) == 0 {
		return nil
	}

	for i := 0; i < len(files); i++ {
		info, _ := files[i].Info()

		line := ""
		for _, f := range printNextFence {
			if f == 1 {
				line += "|\t"
			} else {
				line += "\t"
			}
		}

		if info.IsDir() || (!info.IsDir() && printFiles) {

			printNextFence = append(printNextFence, 1)
			if i == len(files)-1 {
				printNextFence[len(printNextFence)-1] = 0
			}

			decor := "├───"
			if i == len(files)-1 {
				decor = "└───"
			}

			fmt.Printf("%s%s%s", line, decor, info.Name())
			if info.Size() == 0 {
				fmt.Printf(" (empty)\n")
			} else {
				fmt.Printf(" (%db)\n", info.Size())
			}
			dirTreeRec(f, path+"/"+info.Name(), printFiles, level+1, printNextFence)
			printNextFence = printNextFence[:len(printNextFence)-1]
		}
	}

	return nil
}

func main() {
	// out := os.Stdout
	out := new(bytes.Buffer)
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
