package main

import (
	"github.com/dciccale/comment.go/parser"
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
	p := parser.Parser{}
	p.Transform(p.Extract(lines, filename))
}
