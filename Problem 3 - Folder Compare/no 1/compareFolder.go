package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	var dirSource []string
	dirSource = recursiveDir(args[0], dirSource)

	var dirTarget []string
	dirTarget = recursiveDir(args[1], dirTarget)

	// fmt.Println(dirSource)
	// fmt.Println(dirTarget)

	searchTheDiff(dirSource, dirTarget, "NEW")
	searchTheDiff(dirTarget, dirSource, "DELETED")

}

func recursiveDir(path string, dirReturn []string) []string {
	items, _ := ioutil.ReadDir(path)
	for _, item := range items {
		if item.IsDir() {
			dirReturn = recursiveDir(path+"/"+item.Name(), dirReturn)
		} else {
			// fmt.Println(path + "/" + item.Name())
			dirReturn = append(dirReturn, path+"/"+item.Name())
		}
	}
	return dirReturn
}

func searchTheDiff(dir1 []string, dir2 []string, label string) {
	for _, item := range dir1 {
		data1 := strings.Split(item, "//")
		exist := false
		for _, item2 := range dir2 {
			data2 := strings.Split(item2, "//")

			if data1[1] == data2[1] {
				exist = true
				break
			}
		}
		if !exist {
			fmt.Println(data1[1], label)
		}
	}
}

//source: https://stackoverflow.com/questions/31729262/go-iterate-through-directores-files-in-current-directory
