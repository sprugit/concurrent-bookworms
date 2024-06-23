package internal

import "strconv"

type Line struct {
	Number  int
	Content string
}

func NewLine(line_num int, contents string) *Line {
	return &Line{line_num, contents}
}

func (l *Line) toString() string {
	return strconv.FormatInt(int64(l.Number), 10) + ":" + l.Content
}
