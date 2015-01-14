package parser

import (
	"fmt"
	"github.com/dciccale/comment.go/tags"
	"path"
	"regexp"
	"strings"
)

type CommentBlock struct {
	comment  string
	line     int
	filename string
}

type Parser struct {
	root    map[string]interface{}
	tags    tags.Tags
	tag     interface{}
	section struct {
		data map[string]interface{}
		// current []string{}
		// prev []string{}
		mode string
	}
}

func (p *Parser) Extract(lines []string, filename string) map[string][]CommentBlock {

	var REGEX_START_COMMENT = regexp.MustCompile("^\\s*\\/\\*\\\\s*$")
	var REGEX_END_COMMENT = regexp.MustCompile("^\\s*\\\\\\*/\\s*$")

	commentmap := make(map[string][]CommentBlock)
	line := ""
	linenum := 0
	comment := ""
	l := len(lines)

	for i := 0; i < l; i++ {
		line = lines[i]
		if REGEX_START_COMMENT.MatchString(line) {
			commentlines := []string{}

			for i < l && !REGEX_END_COMMENT.MatchString(line) {
				commentlines = append(commentlines, line)
				i++
				line = lines[i]
				linenum = i + 2
			}

			// Remove starting comment /*\
			commentlines = append(commentlines[:0], commentlines[1:]...)

			comment = strings.Join(commentlines, "\n")
			commentmap[filename] = append(commentmap[filename], CommentBlock{comment: comment, line: linenum, filename: filename})
		}
	}
	return commentmap
}

func (p *Parser) Transform(commentmap map[string][]CommentBlock) {
	for file := range commentmap {
		blocks := commentmap[file]

		// Process all comment blocks for the current file
		for i := 0; i < len(blocks); i++ {
			p.ProcessBlock(blocks[i])
		}

		// Generate the toc
		// p.generateTOC(p.root);
	}
}

func (p *Parser) generateTOC() {}

func (p *Parser) ProcessBlock(block CommentBlock) {
	var REGEX_ROW_DATA = regexp.MustCompile("^\\s*(\\S)(?:[^\n])\\s*(.*)$")
	blocklines := strings.Split(block.comment, "\n")

	var firstline = false
	line := ""
	symbol := ""
	value := ""
	title := []string{}
	data := []string{}

	for i := 0; i < 1; i++ {
		line = blocklines[i]
		data = REGEX_ROW_DATA.FindAllStringSubmatch(line, -1)[0]

		// fmt.Println(line)
		// result_slice := re1.FindAllStringSubmatch(line, -1)
		// fmt.Printf("%v", result_slice)
		// fmt.Println(err)

		// fmt.Println(data[0])
		if i == 0 {
			firstline = true
			// fmt.Println(firstline)
			// p.pointer = p.root
		}

		if len(data) > 0 {
			symbol = data[1]
			value = data[2]

			if symbol == p.tags.Get("text") && firstline {
				firstline = false

				// fmt.Println(value)
				title = strings.Split(value, ".")
				// fmt.Println(strings.Split(value, "\\."))
				// fmt.Println(strings.Join(value, ""))
				// for j := 0; j < len(title); j++ {
				// }

				p.section.data = make(map[string]interface{}, 0)
				p.section.data["name"] = value
				p.section.data["title"] = strings.Replace(value, ".", "-", -1)
				p.section.data["line"] = block.line
				p.section.data["filename"] = path.Base(block.filename)
				p.section.data["srclink"] = path.Base(strings.Replace(path.Base(block.filename), path.Ext(block.filename), "", -1))
				p.section.data["level"] = len(title) + 1
			} else {
				p.tag = p.tags.Get(symbol)

				// Change the mode when not matching the current one
				// if p.section.mode != tag.name {
				// 	p.section.current = p.section.prev
				// }

				// Process the value
				// tag.process(value, p.section);

				fmt.Println(p.tag)
				// p.section.mode = p.tag.name
			}
		}
	}
}
