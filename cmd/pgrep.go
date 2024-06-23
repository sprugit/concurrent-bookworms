package main

import (
	"bookworm/internal"
)

func main() {

	internal.NewFileReadTask("../test/loremIpsum.txt", "Lorem", 2).Start()

}
