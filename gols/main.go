package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func getDirContents(path string) {
	fmt.Printf("%s:\n", path)
	dirContents, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln("Can't read dir contents:", err)
	}
	for _, f := range dirContents {
		fmt.Println(f.Name())
	}
}

func main() {
	if len(os.Args) == 1 {
		curDirPath, err := os.Getwd()
		if err != nil {
			log.Fatalln("Can't get current dir path:", err)
		}

		getDirContents(curDirPath)
	} else {
		for _, arg := range os.Args[1:] {
			getDirContents(arg)
			fmt.Println()
		}
	}

}
