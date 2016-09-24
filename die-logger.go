package dieLogger

import (
	"io"
	"os"
	"time"
)

// Logger implementation
type Logger struct {
	// Formatters
	formatters map[Severity]*Formatter

	// Outputs
	output        map[Severity]io.Writer
	defaultOutput io.Writer

	// Log level ignore mechanism
	ignore map[Severity]bool

	// Time measurement mechanism
	start, end *time.Time
}

type loggerFormatters struct {
	debug *Formatter
	info  *Formatter
	err   *Formatter
}

type loggerOutputs struct {
	defaultOutput io.Writer
	debugOutput   io.Writer
	infoOutput    io.Writer
	errorOutput   io.Writer
}

// New instantiates a new logger
func New(allowDebug bool) *Logger {
	defaultOutput := os.Stdout

	l := &Logger{
		formatters: map[Severity]*Formatter{
			Debug: NewFormatter("FgHiBlack"),
			Info:  NewFormatter("FgHiYellow"),
			Error: NewFormatter("FgHiRed"),
		},

		output: map[Severity]io.Writer{
			Debug: nil,
			Info:  nil,
			Error: nil,
		},
		defaultOutput: defaultOutput,

		ignore: map[Severity]bool{
			Debug: allowDebug,
			Info:  false,
			Error: false,
		},
	}

	l.formatters[Debug].SetPrefix("bold", "FgHiBlack", "Debug: ")
	l.formatters[Info].SetPrefix("bold", "FgHiYellow", "Info:  ")
	l.formatters[Error].SetPrefix("bold", "FgHiRed", "Error: ")

	return l
}

// Formatter getter for the formatter of specified severity level
func (l *Logger) Formatter(severity Severity) *Formatter {
	return l.formatters[severity]
}

// Debug logs a debug severity message
func (l *Logger) Debug(message string) {
	l.log(message, Debug)
}

// DebugFields logs a debug severity message with either embedded or attached fields
func (l *Logger) DebugFields(message string, extra ...interface{}) {
	l.log(message, Debug, extra...)
}

// Info logs an info severity message
func (l *Logger) Info(message string) {
	l.log(message, Info)
}

// InfoFields logs a info severity message with either embedded or attached fields
func (l *Logger) InfoFields(message string, extra ...interface{}) {
	l.log(message, Info, extra...)
}

// Error logs an info severity message
func (l *Logger) Error(message string) {
	l.log(message, Error)
}

// ErrorFields logs a error severity message with either embedded or attached fields
func (l *Logger) ErrorFields(message string, extra ...interface{}) {
	l.log(message, Error, extra...)
}

// TimerStart allows starting a timer to allow measuring of execution time of arbitrary code sequances
func (l *Logger) TimerStart() {
	aux := time.Now()
	l.start = &aux
}

// TimerStop allows stopping the timer that allow measuring of execution time of arbitrary code sequances.
// This method returns the elapsed time at execution time
func (l *Logger) TimerStop() int64 {
	var duration int64

	if l.start != nil {
		aux := time.Now()
		l.end = &aux

		duration = l.end.UnixNano() - l.start.UnixNano()
		l.start = nil
	}

	return duration
}

// Logs a message throw the useual logging mechanism
func (l *Logger) log(message string, severity Severity, extras ...interface{}) {
	var msg []byte

	if !l.ignore[severity] {
		if len(extras) > 0 {
			msg = []byte(l.formatters[severity].Format(message, extras...))
		} else {
			msg = []byte(l.formatters[severity].Format(message))
		}

		if l.output[severity] != nil {
			l.output[severity].Write([]byte(msg))
		} else {
			l.defaultOutput.Write([]byte(msg))
		}
	}
}
