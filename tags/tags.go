package tags

import (
// "fmt"
)

type process func(string)

type Tags struct {
	name    string
	tags    map[string]Tag
	symbols map[string]string
}

type Tag struct {
	name    string
	symbol  string
	Process process
	// definition interface{}
}

func (t *Tags) Define(name string, symbol string, fn process) Tag {
	tag := Tag{name: name, symbol: symbol, Process: fn}
	t.tags = make(map[string]Tag)
	t.tags[symbol] = tag
	t.symbols = make(map[string]string)
	t.symbols[name] = symbol
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
