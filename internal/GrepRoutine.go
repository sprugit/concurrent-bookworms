package internal

import (
	"fmt"
	"strings"
	"sync"
)

type GrepRoutine struct {
	id  int
	ctx *RoutineContext
	wg  *sync.WaitGroup
	chn chan []Line
}

func NewGrepRoutine(id int, context *RoutineContext, workgroup *sync.WaitGroup, channel chan []Line) *GrepRoutine {
	g := &GrepRoutine{
		id:  id,
		ctx: context,
		wg:  workgroup,
		chn: channel,
	}
	return g
}

func (g *GrepRoutine) Start() {

	var (
		last_val = true
		contents string
	)

	defer g.wg.Done()
	lines := make([]Line, 0, 300)
	for last_val {
		select {
		case stop := <-(*g.ctx.AppContext).Done():
			fmt.Printf("Routine %d ceasing execution. Reason:", g.id)
			fmt.Print(stop)
			fmt.Printf("\n")
			last_val = false
		default:
			line, canRead := g.ctx.FileScanner.Text()
			last_val = canRead
			if last_val {
				if !*g.ctx.ShouldMatchCase {
					contents = strings.ToLower(line.Content)
				} else {
					contents = line.Content
				}
				if g.ctx.Pattern.Find([]byte(contents)) != nil {
					lines = append(lines, *line)
				}
			}
		}
	}
	g.chn <- lines
}
