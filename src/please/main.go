package main

import (
	"gopkg.in/yaml.v2"
	"fmt"
	"os"
)

type PleaseData struct {
	Playlist []string `yaml:"playlist"`
}

func main() {
	pls := PleaseData{}
	args := os.Args[1:]
	if len(args) < 1 {
		panic("please pass a file path to load")
	}
	path := args[0]
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("could not access file at path `%s`\n", path))
	}
	if err := yaml.Unmarshal([]byte(dat), &pls); err != nil {
		panic(fmt.Sprintf("encountered a problem when loading the playlist file:\n%v\n", err))
	}
	fmt.Printf("%+v\n", pls)
}
