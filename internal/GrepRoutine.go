package internal

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

type GrepRoutine struct {
	id      int
	scanner *AtomicScanner
	wg      *sync.WaitGroup
	chn     chan []Line
	ctx     *context.Context
}

func NewGrepRoutine(id int, scanner *AtomicScanner, wg *sync.WaitGroup, chn chan []Line, ctx *context.Context) *GrepRoutine {
	g := &GrepRoutine{
		id:      id,
		scanner: scanner,
		wg:      wg,
		chn:     chn,
		ctx:     ctx,
	}
	return g
}

func (g *GrepRoutine) Start(pattern *string) {
	fmt.Printf("Hello, I am thread")
	defer g.wg.Done()
	lines := make([]Line, 0, 300)
	var last_val = true
	select {
	case stop := <-(*g.ctx).Done():
		fmt.Printf("Routine %d ceasing execution. Reason:", g.id)
		fmt.Print(stop)
		fmt.Printf("\n")
	default:
		for last_val {
			line, canRead := g.scanner.Text()
			last_val = canRead
			if last_val {
				fmt.Printf("\nRead line #%d: %s \n", line.Number, line.Content)
				if strings.Contains(line.Content, *pattern) {
					fmt.Printf("Line contains %s \n", *pattern)
					lines = append(lines, *line)
				}
			}
		}
		g.chn <- lines
	}
}
