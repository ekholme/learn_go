package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

//the purpose of this code is to recurse through all files in a directory

func main() {
	//option 1
	// files, err := ioutil.ReadDir(".")

	// if err != nil {
	// 	log.Fatalf("Couldn't read directory: %v", err)
	// }

	// for _, file := range files {
	// 	// fmt.Printf("%s, %v\n", file.Name(), file.IsDir())

	// 	if file.IsDir() {

	// 		p := fmt.Sprintf("./%s", file.Name())

	// 		// fmt.Println(p)

	// 		file2, err := ioutil.ReadDir(p)

	// 		if err != nil {
	// 			log.Fatalf("Couldn't read nested directory")
	// 		}

	// 		for _, f := range file2 {
	// 			fmt.Printf("%s, %v\n", f.Name(), f.IsDir())
	// 		}
	// 	}
	// }

	//option 2
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Error: %v", err)
			return err
		}

		fmt.Printf("dir: %v, name: %s\n", info.IsDir(), path)

		return nil
	})

	if err != nil {
		log.Fatalf("Couldn't read filepath: %v", err)
	}
	//this is handles recursion in a more straightforward manner
}
