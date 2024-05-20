package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	p := flag.Uint("p", 1, "Paralelization level: total number of goroutines to be ran.")
	target := flag.String("target", "", "Target file to be skimmed through")
	pattern := flag.String("pattern", "", "Pattern to be found in <-target> file.")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if *p <= 0 || *target == "" || *pattern == "" {
		log.Println(p, target, pattern)
		log.Println("Invalid parameters were passed. -p can't be 0 or below. Neither target file nor pattern can be empty.")
		os.Exit(1)
	}

	f, err := os.Open(*target)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(line, *pattern) {
			fmt.Println(*target, ":", line)
		}
	}

}
