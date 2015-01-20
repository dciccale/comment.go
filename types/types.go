package types

type Data struct {
	Name     string
	Title    string
	Line     int
	Filename string
	Srclink  string
	Level    int
}

type Section struct {
	Data    Data
	Current []Data
	Prev    []Data
	Mode    string
}

type Docs struct {
	Name     string
	Sections []Data
	// Toc
}
