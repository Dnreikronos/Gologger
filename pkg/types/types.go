package types

type LogLine struct {
	Filename string
	Line     string
}

type ParsedLog struct {
	Filename string
	Level    string
	Message  string
	Timestamp string
}

