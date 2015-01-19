package main

import (
	"fmt"
	"github.com/dciccale/comment.go/parser"
	"github.com/dciccale/comment.go/tags"
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

	var tag = t.Define("text", "*", func(msg string) { fmt.Println(msg) })
	tag.Process("process")

	p := parser.Parser{Tags: &t}
	p.Transform(p.Extract(lines, filename))
}
