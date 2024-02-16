/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	/*
		"fmt"
		"io/ioutil"
		"os"
		"path/filepath"
	*/
	"github.com/minyou08042/kouyi_palm_treasure/cmd"
)

func main() {
	/*
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		tdir := dir

		files, err := ioutil.ReadDir(tdir)
		if err != nil {
			fmt.Println(err)
		}

		for _, file := range files {
			fmt.Println(file.Name(), file.IsDir())
		}
	*/
	cmd.Execute()
}
