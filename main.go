package main

import (
	"fmt"
	"github.com/dciccale/comment.go/parser"
	"github.com/dciccale/comment.go/tags"
	"github.com/dciccale/comment.go/types"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := "kimbo.js"
	data, err := ioutil.ReadFile(filename)
	check(err)

	var lines = strings.Split(string(data), "\n")
	t := tags.Tags{}

	t.Define("text", "*", func(value string, section types.Section) {
		fmt.Println(value)
		// data := make(map[string]types.Data, 0)
		// data["Text"] = value
		// section.Current = append(section.Current, data)
	}, true)
	// tag.Process("process")

	p := parser.Parser{Tags: &t}
	p.Transform(p.Extract(lines, filename))
}
