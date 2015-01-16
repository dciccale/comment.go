package main

import (
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
	t.Define("text", "*")
	p := parser.Parser{Tags: &t}
	p.Transform(p.Extract(lines, filename))
}
