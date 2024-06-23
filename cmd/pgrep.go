package main

import (
	"bookworm/internal"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

func main() {

	filename := flag.String("file", "./test/loremIpsum.txt", "files (separated by a comma to search")
	pattern := flag.String("pattern", "lorem", "pattern to be searched for")
	case_sense := flag.Bool("case-sensitive", false, "whether search should or not be case sensitive")
	numRoutines := flag.Int("r", 2, "number of goroutines per file")

	flag.Parse()

	files := strings.Split(*filename, ",")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	for _, file := range files {
		err := internal.GrepFileForPattern(&file, pattern, case_sense, numRoutines, &ctx)
		if err != nil {
			fmt.Println(err)
		}
	}

}
