package runtimetools

import (
	"runtime"
)

// GetCurrentFnName ...
func GetCurrentFnName() string {
	var (
		pc       = make([]uintptr, 10)
		n        = runtime.Callers(2, pc)
		frame, _ = runtime.CallersFrames(pc[:n]).Next()
	)
	return frame.Function
}

// GetTraceLogs ...
func GetTraceLogs() []*TraceDataLine {
	var (
		pc     = make([]uintptr, 15)
		n      = runtime.Callers(2, pc)
		frames = runtime.CallersFrames(pc[:n])
		tl     []*TraceDataLine
	)

	for {
		frame, more := frames.Next()

		tl = append(tl, &TraceDataLine{
			File:     frame.File,
			Line:     frame.Line,
			Function: frame.Function,
		})

		if !more {
			break
		}
	}
	return tl
}
