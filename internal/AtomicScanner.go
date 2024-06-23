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

func (s *AtomicScanner) Scan() bool {
	return s.scanner.Scan()
}

func (s *AtomicScanner) Text() *Line {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.counter += 1
	return NewLine(s.counter, s.scanner.Text())
}
