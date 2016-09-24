package dieLogger

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	// Map of all supported styles (mapping to underlying color package)
	mapColorStyles = map[string]color.Attribute{
		"reset":        color.Reset,
		"bold":         color.Bold,
		"faint":        color.Faint,
		"italic":       color.Italic,
		"underline":    color.Underline,
		"blinkslow":    color.BlinkSlow,
		"blinktapid":   color.BlinkRapid,
		"reversevideo": color.ReverseVideo,
		"concealed":    color.Concealed,
		"crossedout":   color.CrossedOut,
		"fgblack":      color.FgBlack,
		"fgred":        color.FgRed,
		"fggreen":      color.FgGreen,
		"fgyellow":     color.FgYellow,
		"fgblue":       color.FgBlue,
		"fgmagenta":    color.FgMagenta,
		"fgcyan":       color.FgCyan,
		"fgwhite":      color.FgWhite,
		"fghiblack":    color.FgHiBlack,
		"fghired":      color.FgHiRed,
		"fghigreen":    color.FgHiGreen,
		"fghiyellow":   color.FgHiYellow,
		"fghiblue":     color.FgHiBlue,
		"fghimagenta":  color.FgHiMagenta,
		"fghicyan":     color.FgHiCyan,
		"fghiwhite":    color.FgHiWhite,
		"bgblack":      color.BgBlack,
		"bgred":        color.BgRed,
		"bggreen":      color.BgGreen,
		"bgyellow":     color.BgYellow,
		"bgblue":       color.BgBlue,
		"bgmagenta":    color.BgMagenta,
		"bgcyan":       color.BgCyan,
		"bgwhite":      color.BgWhite,
		"bghiblack":    color.BgHiBlack,
		"bghired":      color.BgHiRed,
		"bghigreen":    color.BgHiGreen,
		"bghiyellow":   color.BgHiYellow,
		"bghiblue":     color.BgHiBlue,
		"bghimagenta":  color.BgHiMagenta,
		"bghicyan":     color.BgHiCyan,
		"bghiwhite":    color.BgHiWhite,
	}
)

// Formatter is the main formatting functionality provider
type Formatter struct {
	c      *color.Color
	prefix string
	suffix string
	format string
}

// NewFormatter instantiates a new formatter
func NewFormatter(style ...string) *Formatter {
	var colorAttrs = make([]color.Attribute, len(style))

	for i, c := range style {
		t := strings.ToLower(c)
		if cc, found := mapColorStyles[t]; found {
			colorAttrs[i] = cc
		}
	}

	return &Formatter{
		format: "%s %s %s\n",
		c:      color.New(colorAttrs...),
	}
}

// SetPrefix allow definition of a new prefix for all strings created throw this formatter
// This method allows introduction of styling directives alongside constant segments
// and will use all constant segments as prefix and all styling as stylers for
// those segments in order of aparition
func (f *Formatter) SetPrefix(prefix ...string) {
	f.prefix = f.formatInline(prefix...)
}

// SetSuffix allows definition of a new suffix for all strings create throw this formatter
// This method allows introduction of styling directives alongside constant segments
// and will use all constant segments as sufix and all styling as stylers for
// those segments in order of aparition
func (f *Formatter) SetSuffix(suffix ...string) {
	f.suffix = f.formatInline(suffix...)
}

// Formats a strings sequence by specified formatters
func (f *Formatter) formatInline(segments ...string) string {
	var currentStyle []color.Attribute
	var constant []string

	for _, s := range segments {
		t := strings.ToLower(s)

		if ss, found := mapColorStyles[t]; found {
			// Add the styling to the current styling array
			currentStyle = append(currentStyle, ss)
		} else {
			// Found a constant segment
			if len(currentStyle) > 0 {
				// Format the string as specified
				c := color.New(currentStyle...)
				constant = append(constant, c.SprintfFunc()(s))

				// Reset styling
				currentStyle = make([]color.Attribute, 0)
			} else {
				constant = append(constant, s)
			}
		}
	}

	return strings.Join(constant, "")
}

// Format the input against registered format
func (f *Formatter) Format(what string, a ...interface{}) string {
	content := f.c.SprintfFunc()(what, a...)
	return fmt.Sprintf(f.format, f.prefix, content, f.suffix)
}
