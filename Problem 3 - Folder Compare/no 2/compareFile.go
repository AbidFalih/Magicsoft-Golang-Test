package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	var dirSource []string
	dirSource = recursiveDir(args[0], dirSource)
	// fmt.Println(dirSource)

	var dirTarget []string
	dirTarget = recursiveDir(args[1], dirTarget)
	// fmt.Println(dirTarget)

	searchTheDiff(dirSource, dirTarget, "NEW", true)
	searchTheDiff(dirTarget, dirSource, "DELETED", false)
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

func searchTheDiff(dir1 []string, dir2 []string, label string, modif bool) {
	for _, item := range dir1 {
		data1 := strings.Split(item, "//")
		exist := false
		for _, item2 := range dir2 {
			data2 := strings.Split(item2, "//")

			if data1[1] == data2[1] {
				exist = true
				if modif {
					f1, _ := ioutil.ReadFile(item)
					// fmt.Println("f1", f1)

					f2, _ := ioutil.ReadFile(item2)
					// fmt.Println("f2", f2)

					if !bytes.Equal(f1, f2) {
						fmt.Println(data1[1], "MODIFIED")
					}
				}
				break
			}
		}
		if !exist {
			fmt.Println(data1[1], label)
		}
	}
}
