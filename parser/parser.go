package parser

import (
	"path"
	"regexp"
	"strings"

	"github.com/dciccale/comment.go/tags"
	"github.com/dciccale/comment.go/types"
)

type CommentBlock struct {
	comment  string
	line     int
	filename string
}

type Parser struct {
	Tags      *tags.Tags
	root      map[string]interface{}
	pointer   *map[string]interface{}
	lvl       []string
	Section   types.Section
	tocData   map[string]types.Data
	BlockData map[string][]types.Data
	toc       map[string]types.Toc
	utoc      map[string]int
}

// func (p *Parser) Parse(filemap map[string]string) types.Docs {
// 	p.Transform(p.Extract(filemap))
// 	docs := types.Docs{Name: p.docsName, Sections: p.sections /*, Toc: p.toc*/}
// 	return docs
// }

var REGEX_ROW_DATA = regexp.MustCompile("^\\s*(\\S)(?:[^\n])\\s*(.*)$")

func NewParser(t *tags.Tags) *Parser {
	return &Parser{
		Tags:      t,
		BlockData: make(map[string][]types.Data),
		tocData:   make(map[string]types.Data),
		root:      make(map[string]interface{}),
		lvl:       make([]string, 0),
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
		// p.generateTOC(&p.root)
	}
}

func (p *Parser) ProcessBlock(block CommentBlock) {
	blocklines := strings.Split(block.comment, "\n")

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

		if i == 0 {
			firstline = true
			// p.pointer = &p.root
		}

		if len(match) > 0 {
			data = match[0]

			if len(data) >= 0 {
				symbol = data[1]
				value = data[2]

				if symbol == p.Tags.GetSymbol("text") && firstline {
					firstline = false

					title = strings.Split(value, ".")
					// for j := 0; j < len(title); j++ {
					// 	if val, ok := p.pointer.(map[string]interface{})[title[j]].(map[string]interface{}); ok {
					// 		p.pointer = p.pointer[title[j]]
					// 	} else {
					// 		p.pointer[title[j]] = make(map[string]interface{})
					// 	}
					// }

					p.Section.Data = types.Data{
						Name:     value,
						Title:    strings.Replace(value, ".", "-", -1),
						Line:     block.line,
						Filename: path.Base(block.filename),
						Srclink:  path.Base(strings.Replace(path.Base(block.filename), path.Ext(block.filename), "", -1)),
						Level:    len(title) + 1,
					}

					dataBlock := []types.Data{p.Section.Data}
					p.Section.Current = &dataBlock
					p.Section.Prev = &dataBlock

				} else if p.Tags.Exist(symbol) {
					tag = p.Tags.Get(symbol)

					// Change the mode when not matching the current one
					if p.Section.Mode != tag.Name {
						p.Section.Current = p.Section.Prev
					}

					// Process the value
					tag.Process(value, &p.Section)

					p.Section.Mode = tag.Name
				}
			}
		}
	}

	// Map the Section data by name to generate the toc later
	p.tocData[p.Section.Data.Name] = p.Section.Data

	// Map each Section by name to be able to order it later according the toc
	p.BlockData[p.Section.Data.Name] = *p.Section.Prev
}

// func (p *Parser) generateTOC(pointer *map[string]interface{}) {
// 	levels := make([]string, 0)
// 	var level string
// 	var name string
// 	// var brackets bytes.Buffer
// 	// var indent int
// 	var sectionData types.Data
// 	// var isMethod bool
//
// 	for k, _ := range *pointer {
// 		levels = append(levels, k)
// 	}
// 	fmt.Println(levels)
//
// 	// Sort alphabetically and format data
// 	// levels = levels.sort();
// 	l := len(levels)
// 	for i := 0; i < l; i++ {
// 		level = levels[i]
// 		//
// 		p.lvl = append(p.lvl, level)
// 		name = strings.Join(p.lvl, ".")
// 		sectionData = p.tocData[name]
// 		fmt.Println(sectionData.Type, name)
// 		//
// 		//   name = this.lvl.join('.');
// 		//   sectionData = this.tocData[name];
// 		//   isMethod = sectionData.type && sectionData.type.indexOf('method') + 1;
// 		//   indent = this.lvl.length - 1;
// 		//
// 		//   if isMethod {
// 		//     if sectionData.params && len(sectionData.params) > 0 {
// 		//       if len(sectionData.params) == 1 {
// 		//         brackets.WriteString("(")
// brackets.WriteString(strings.Join(sectionData.params[0], ", "))
// brackets.WriteString(")")
// 		//       } else {
// brackets.WriteString("(\u2026)")
// 		//         brackets = '(\u2026)';
// 		//       }
// 		//     } else {
// 		//       brackets.WriteString("()")
// 		//     }
// 		//   } else {
// 		//     brackets.WriteString("")
// 		//   }
// 		//
// 		//   sectionData.brackets = brackets.String()
// 		//
// 		//   // Prevent duplicates
// 		//   if (!p.utoc[name]) {
// 		//     p.sections.push(this.blockData[name]);
// if isMethod {
// 	b = "()"
// } else {
// 	b = ""
// }
// toc := types.Toc{
// 		      Indent: indent,
// 		      Name: name,
// 		      Type: sectionData.Type,
// 		      Brackets: b
// }
// p.toc = append(p.toc, toc)
// 		//     p.utoc[name] = 1;
// 		//   }
// 		// p.generateTOC(pointer[level])
// 		p.lvl = p.lvl[:len(p.lvl)-1]
// 	}
// }
