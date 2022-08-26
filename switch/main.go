package main

import "fmt"

type Level int8
const (
	LogLevelInfo Level = iota+1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
)
func main(){
	text := "INFO"
	var l Level
	switch string(text) {
	case "debug", "DEBUG":
		l = DebugLevel
	case "info", "INFO", "": // make the zero value useful
		l = InfoLevel
	case "warn", "WARN":
		l = WarnLevel
	case "error", "ERROR":
		l = ErrorLevel
	case "dpanic", "DPANIC":
		l = DPanicLevel
	case "panic", "PANIC":
		l = PanicLevel
	case "fatal", "FATAL":
		l = FatalLevel
	default:
		l = FatalLevel
	}

	fmt.Printf("l:%#v\n", l)
}
