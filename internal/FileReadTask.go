package internal

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type FileReadTask struct {
	filename     string
	num_routines int
	pattern      string
}

func NewFileReadTask(filename string, pattern string, num_routines int) *FileReadTask {
	return &FileReadTask{
		filename:     filename,
		num_routines: num_routines,
		pattern:      pattern,
	}
}

func (ft *FileReadTask) Start() error {

	fmt.Printf("Initializing task for file %s with %d routines seeking pattern %s\n",
		ft.filename, ft.num_routines, ft.pattern)

	results := make([]string, 0, 500)
	wg := new(sync.WaitGroup)
	wg.Add(ft.num_routines)
	channels := make([]chan []Line, ft.num_routines)
	routines := make([]GrepRoutine, ft.num_routines)
	file, err := os.Open(ft.filename)
	if err != nil {
		return fmt.Errorf("", err)
	}
	defer file.Close()
	bufsc := bufio.NewScanner(file)
	bufsc.Split(bufio.ScanLines)
	scanner := NewAtomicScanner(bufsc)

	for i := 0; i < ft.num_routines; i++ {
		var iteration int = i
		channels[iteration] = make(chan []Line)
		routines[iteration] = *NewGrepRoutine(iteration, scanner, wg, channels[iteration])
		go routines[iteration].Start(&ft.pattern)
	}

	for _, element := range channels {
		lines := <-element
		for _, line := range lines {
			results = append(results, line.toString())
		}
	}

	wg.Wait()

	fmt.Printf("\nFound the pattern %s %d times.\nPattern is contained in the following lines %s.\n", ft.pattern, len(results), ft.pattern)
	for _, line := range results {
		fmt.Printf("\n%s\n", line)
	}

	return nil
}
