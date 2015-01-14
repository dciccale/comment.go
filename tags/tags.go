package tags

import (
	"fmt"
)

type Tags struct {
	name    string
	tags    map[string]interface{}
	symbols map[string]interface{}
}

type Tag struct {
	name       string
	symbol     string
	definition interface{}
}

func (t *Tag) Process(value string, section interface{}) {
	fmt.Println("process")
}

func (t *Tags) Define(name string, definition interface{}) Tag {
	tag := Tag{name: name, definition: definition}
	t.tags[tag.symbol] = tag
	t.symbols[name] = tag.symbol
	return tag
}

func (t *Tags) Get(q string) interface{} {
	if val, ok := t.tags[q]; ok {
		return val
	} else if val, ok := t.symbols[q]; ok {
		return val
	} else {
		return nil
	}
}
