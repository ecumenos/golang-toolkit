package runtimetools

// TraceDataLine ...
type TraceDataLine struct {
	File     string `json:"file"`
	Line     int    `json:"line"`
	Function string `json:"function"`
}
