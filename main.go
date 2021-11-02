package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//./txtreplace -ext txt -old Bin28329 -new New999
func main() {
	var searchPath = flag.String("path", "C:\\", "The path to search for the file extension")
	var extension = flag.String("ext", "rdp", "The extension to search for")

	var old = flag.String("old", "", "The text to replace in the files")
	var new = flag.String("new", "", "The new text to replace old")
	flag.Parse()

	fmt.Println(*searchPath)
	fmt.Println(*extension)
	fmt.Println(*old)
	fmt.Println(*new)
	fmt.Println("Running")
	files, err := WalkMatch(*searchPath, "*."+*extension)
	if err != nil {
		fmt.Println(err)
	}

	CheckFiles(files, *old, *new)
}

func CheckFiles(files []string, oldStr string, newStr string) {
	for _, elm := range files {
		input, err := ioutil.ReadFile(elm)
		if err != nil {
			fmt.Println("Failed to read from " + elm + " Aborting and skipping to next")
			continue
		}
		var inStr = string(input)
		if strings.Contains(inStr, oldStr) {
			fmt.Println("File found: " + elm)
			out := strings.ReplaceAll(inStr, oldStr, newStr)
			ioutil.WriteFile(elm, []byte(out), 0)
		}
	}
}

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		//fmt.Println(path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
