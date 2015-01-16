package tags

import (
	"fmt"
)

type Tags struct {
	name    string
	tags    map[string]Tag
	symbols map[string]string
}

type Tag struct {
	name   string
	symbol string
	// definition interface{}
}

func (t *Tag) Process(value string, section interface{}) {
	fmt.Println("process")
}

func (t *Tags) Define(name string, symbol string) {
	tag := Tag{name: name, symbol: symbol}
	t.tags = make(map[string]Tag)
	t.tags[symbol] = tag
	t.symbols = make(map[string]string)
	t.symbols[name] = symbol
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
