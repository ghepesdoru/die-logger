package dieLogger

import (
	"errors"
	"testing"
)

func TestLogger(t *testing.T) {
	l := New(false)
	l2 := New(true)

	l.Debug("Test")
	l.DebugFields("Test %s", "random")
	l.DebugFields("Test %s", "random", 1, 2, 3)
	l2.Debug("Test")

	l.Info("Random informative message")
	l.ErrorFields("Random error message", errors.New("Random error here"))

	l.Formatter(Debug).SetSuffix("bold", "FgHiBlue", "!")
	l.TimerStart()
	l.Debug("Test with suffix")
	end := l.TimerStop()
	l.DebugFields("The time to write one line was: %#vns", end)
}
