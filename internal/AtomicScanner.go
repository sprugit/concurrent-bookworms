package internal

import (
	"bufio"
	"sync"
)

type AtomicScanner struct {
	counter int
	scanner *bufio.Scanner
	lock    sync.Mutex
}

func NewAtomicScanner(scanner *bufio.Scanner) *AtomicScanner {
	return &AtomicScanner{
		counter: 0,
		scanner: scanner,
		lock:    sync.Mutex{},
	}
}

func (s *AtomicScanner) Text() (*Line, bool) {

	var line Line
	s.lock.Lock()
	canRead := s.scanner.Scan()
	if canRead {
		s.counter += 1
		line = *NewLine(s.counter, s.scanner.Text())
	}
	s.lock.Unlock()

	return &line, canRead
}
