package internal

import (
	"bufio"
	"cmp"
	"context"
	"fmt"
	"log"
	"os"
	"slices"
	"sync"
)

func GrepFileForPattern(
	filename *string,
	pattern *string,
	matchCase *bool,
	numRoutines *int,
	ctx *context.Context,
) error {

	file, err := os.Open(*filename)
	if err != nil {
		return fmt.Errorf("Task Creation for file %t seeking pattern %t (case sensitive: %t ) using %t routines failed.\nReason:%t",
			*filename, *pattern, *matchCase, *numRoutines, err)
	}
	defer file.Close()

	rout_ctx, err := NewRoutineContext(
		ctx,
		NewAtomicScanner(bufio.NewScanner(file)),
		pattern,
		matchCase,
	)
	if err != nil {
		return fmt.Errorf("Unexpected error related to pattern: %s", err.Error())
	}

	log.Printf("Created Task for file %t seeking pattern %t (case sensitive: %t ) using %t routines.",
		*filename, *pattern, *matchCase, *numRoutines)

	results := make([]Line, 0, 500)
	wg := new(sync.WaitGroup)
	wg.Add(*numRoutines)

	channels := make([]chan []Line, *numRoutines)
	routines := make([]GrepRoutine, *numRoutines)

	for i := 0; i < *numRoutines; i++ {
		var iteration int = i
		channels[iteration] = make(chan []Line)
		routines[iteration] = *NewGrepRoutine(iteration, rout_ctx, wg, channels[iteration])
		go routines[iteration].Start()
	}

	for _, element := range channels {
		lines := <-element
		for _, line := range lines {
			results = append(results, line)
		}
	}

	wg.Wait()

	slices.SortFunc(results, func(a, b Line) int {
		return cmp.Compare(a.Number, b.Number)
	})

	fmt.Printf("\nFound the pattern %s in %d lines.\nListing the matches:\n", *pattern, len(results))
	for _, line := range results {
		fmt.Printf("\n%d:%s\n", line.Number, line.Content)
	}

	return nil
}
