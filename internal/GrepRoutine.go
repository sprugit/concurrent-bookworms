package internal

import (
	"fmt"
	"strings"
	"sync"
)

type GrepRoutine struct {
	id      int
	scanner *AtomicScanner
	wg      *sync.WaitGroup
	chn     chan []Line
}

func NewGrepRoutine(id int, scanner *AtomicScanner, wg *sync.WaitGroup, chn chan []Line) *GrepRoutine {
	g := &GrepRoutine{
		id:      id,
		scanner: scanner,
		wg:      wg,
		chn:     chn,
	}
	return g
}

func (g *GrepRoutine) Start(pattern *string) {
	defer g.wg.Done()
	lines := make([]Line, 0, 300)
	for g.scanner.Scan() {
		line := g.scanner.Text()
		fmt.Printf("\nRead line #%d: %s \n", line.Number, line.Content)
		if strings.Contains(line.Content, *pattern) {
			fmt.Printf("Line contains %s \n", *pattern)
			lines = append(lines, *line)
		}
	}
	g.chn <- lines
}
