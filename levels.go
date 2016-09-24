package dieLogger

import "strings"

// Severity of the log message
type Severity uint8

// Definition of all levels available
const (
	Debug Severity = iota
	Info
	Error
)

var (
	severityToString = map[Severity]string{
		Debug: "debug",
		Info:  "info",
		Error: "error",
	}

	stringToSeverity = map[string]Severity{
		"debug": Debug,
		"info":  Info,
		"error": Error,
	}
)

// Severity
func (s Severity) String() string {
	return severityToString[s]
}

// SeverityFromString parses the severity out of it's string representation
func SeverityFromString(s string) Severity {
	if severity, found := stringToSeverity[strings.ToLower(s)]; found {
		return severity
	}

	return Error
}
