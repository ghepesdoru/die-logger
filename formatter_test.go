package dieLogger

import "testing"

func TestFormatter(t *testing.T) {
	var expected = "\x1b[1;33m[\x1b[0m\x1b[0;3;32mInfo\x1b[0m\x1b[1;33m]\x1b[0m\x1b[0m:\x1b[0m \x1b[4;90mRandom error here: \"random\"\x1b[0m \x1b[1;31m!\x1b[0m\n"

	f := NewFormatter("Underline", "FgHiBlack")
	f.SetPrefix("bold", "FgYellow", "[", "reset", "italic", "fggreen", "Info", "bold", "FgYellow", "]", "reset", ":")
	f.SetSuffix("bold", "FgRed", "!")

	actual := f.Format("Random error here: \"%s\"", "random")
	if expected != actual {
		t.Errorf("Invalid output from formatter:\n Expected:\t%s\n Actual:\t%s", expected, actual)
	}
}
