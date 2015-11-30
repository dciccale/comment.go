package types

type Data struct {
	Name     string `json:"name,omitempty"`
	Title    string `json:"title,omitempty"`
	Line     int    `json:"line,omitempty"`
	Filename string `json:"filename,omitempty"`
	Srclink  string `json:"srclink,omitempty"`
	Level    int    `json:"level,omitempty"`
	Text     string `json:"text,omitempty"`
	Type     string `json:"type,omitempty"`
	Head     string `json:"head,omitempty"`
}

type Section struct {
	Data    Data
	Current *[]Data
	Prev    *[]Data
	Mode    string
}

type Docs struct {
	Name     string
	Sections []Data
	// Toc
}

type Toc struct {
	Indent   int
	Name     string
	Type     string
	Brackets string
}
