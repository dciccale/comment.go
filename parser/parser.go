package parser

import (
	// "fmt"
	"github.com/dciccale/comment.go/tags"
	"github.com/dciccale/comment.go/types"
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
	Tags      *tags.Tags
	root      map[string]interface{}
	section   types.Section
	tocData   map[string]types.Data
	blockData map[string][]types.Data
}

// func (p *Parser) Parse(filemap map[string]string) types.Docs {
// 	p.Transform(p.Extract(filemap))
// 	docs := types.Docs{Name: p.docsName, Sections: p.sections /*, Toc: p.toc*/}
// 	return docs
// }

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
			commentmap[filename] = append(commentmap[filename], CommentBlock{
				comment:  comment,
				line:     linenum,
				filename: filename,
			})
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
	blocklines := strings.Split(block.comment, "\n")

	var REGEX_ROW_DATA = regexp.MustCompile("^\\s*(\\S)(?:[^\n])\\s*(.*)$")
	var firstline = false
	var tag tags.Tag
	var line string
	var symbol string
	var value string
	title := []string{}
	match := [][]string{}
	data := []string{}

	for i := 0; i < len(blocklines); i++ {
		line = blocklines[i]
		match = REGEX_ROW_DATA.FindAllStringSubmatch(line, -1)

		if len(match) > 0 {
			data = match[0]
			if i == 0 {
				firstline = true
				// p.pointer = p.root
			}

			if len(data) >= 0 {
				symbol = data[1]
				value = data[2]

				if symbol == p.Tags.GetSymbol("text") && firstline {
					firstline = false

					title = strings.Split(value, ".")
					// for j := 0; j < len(title); j++ {
					// 	p.pointer = p.pointer[title[j]] = p.pointer[title[j]] || {};
					// }

					p.section.Data.Name = value
					p.section.Data.Title = strings.Replace(value, ".", "-", -1)
					p.section.Data.Line = block.line
					p.section.Data.Filename = path.Base(block.filename)
					p.section.Data.Srclink = path.Base(strings.Replace(path.Base(block.filename), path.Ext(block.filename), "", -1))
					p.section.Data.Level = len(title) + 1

					dataBlock := make([]types.Data, 0)
					dataBlock = append(dataBlock, p.section.Data)
					p.section.Current = dataBlock
					p.section.Prev = dataBlock

				} else if p.Tags.Exist(symbol) {
					tag = p.Tags.Get(symbol)

					// Change the mode when not matching the current one
					if p.section.Mode != tag.Name {
						p.section.Current = p.section.Prev
					}

					// Process the value
					// tag.Process(value, p.section)

					p.section.Mode = tag.Name
				}
			}
		}
	}

	// Map the section data by name to generate the toc later
	tocData := make(map[string]types.Data, 0)
	tocData[p.section.Data.Name] = p.section.Data
	p.tocData = tocData

	// Map each section by name to be able to order it later according the toc
	blockData := make(map[string][]types.Data, 0)
	blockData[p.section.Data.Name] = p.section.Prev
	p.blockData = blockData
}
