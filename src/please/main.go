package main

import (
	"gopkg.in/yaml.v2"
	"fmt"
	"os"
)

type PleaseData struct {
	Title string `yaml:"title"`
	Author string `yaml:"author"`
	Sort string `yaml:"sort"`
	Playlist []string `yaml:"playlist"`
}

func main() {
	var pls = PleaseData{}
	var args = os.Args[1:]
	if len(args) < 1 {
		panic("please pass a file path to load")
	}
	var path = args[0]
	var dat, fileErr = os.ReadFile(path)
	if fileErr != nil {
		panic(fmt.Sprintf("could not access file at path `%s`\n", path))
	}
	var yamlErr = yaml.Unmarshal([]byte(dat), &pls)
	if yamlErr != nil {
		panic(fmt.Sprintf("encountered a problem when loading the playlist file:\n%v\n", yamlErr))
	}
	fmt.Printf("%+v\n", pls)
}
