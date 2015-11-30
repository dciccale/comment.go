package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/dciccale/comment.go/parser"
	"github.com/dciccale/comment.go/tags"
	"github.com/dciccale/comment.go/types"
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
	t := tags.NewTags()

	t.Define("text", "*", func(value string, section *types.Section) {
		data := types.Data{Text: value}
		*section.Current = append(*section.Current, data)
	}, true)

	t.Define("type", "[", func(value string, section *types.Section) {
		reg := regexp.MustCompile("\\s*]\\s*$")
		value = reg.ReplaceAllString(value, "")
		section.Data.Type = value
	}, false)

	t.Define("head", ">", func(value string, section *types.Section) {
		data := types.Data{Head: value}
		*section.Current = append(*section.Current, data)
	}, false)

	p := parser.NewParser(t)
	p.Transform(p.Extract(lines, filename))

	jsonStr, err := json.Marshal(p.BlockData)
	fmt.Println(string(jsonStr))
}
