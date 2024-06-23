package internal

import (
	"context"
	"regexp"
	"strings"
)

type RoutineContext struct {
	AppContext      *context.Context
	FileScanner     *AtomicScanner
	Pattern         *regexp.Regexp
	ShouldMatchCase *bool
}

func NewRoutineContext(
	AppContext *context.Context,
	FileScanner *AtomicScanner,
	Pattern *string,
	ShouldMatchCase *bool) (*RoutineContext, error) {
	var pattern string
	if *ShouldMatchCase {
		pattern = *Pattern
	} else {
		pattern = strings.ToLower(*Pattern)
	}
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &RoutineContext{
		AppContext:      AppContext,
		FileScanner:     FileScanner,
		Pattern:         r,
		ShouldMatchCase: ShouldMatchCase,
	}, nil
}
