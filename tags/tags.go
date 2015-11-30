package tags

import
// "fmt"

"github.com/dciccale/comment.go/types"

type process func(string, *types.Section)

type Tags struct {
	tags    map[string]Tag
	symbols map[string]string
}

type Tag struct {
	Name     string
	Symbol   string
	Process  process
	IsSingle bool
}

func NewTags() *Tags {
	return &Tags{tags: make(map[string]Tag), symbols: make(map[string]string)}
}

func (t *Tags) Define(name string, symbol string, fn process, isSingle bool) Tag {
	tag := Tag{Name: name, Symbol: symbol, Process: fn, IsSingle: isSingle}
	t.tags[symbol] = tag
	t.symbols[name] = symbol
	return tag
}

func (t *Tags) Get(q string) Tag {
	if _, ok := t.tags[q]; ok {
		return t.tags[q]
	} else {
		return Tag{}
	}
}

func (t *Tags) GetSymbol(q string) string {
	if _, ok := t.symbols[q]; ok {
		return t.symbols[q]
	} else {
		return ""
	}
}

func (t *Tags) Exist(q string) bool {
	_, ok := t.tags[q]
	return ok
}
