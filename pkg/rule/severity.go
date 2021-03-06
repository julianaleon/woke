package rule

import (
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

// Severity is a log severity
type Severity int

const (
	// SevError translates to Error
	// This will be the default severity
	SevError Severity = iota
	// SevWarn translates to Warn
	SevWarn
	// SevInfo translates to Info
	SevInfo
)

// NewSeverity turns a string into a Severity
func NewSeverity(s string) Severity {
	switch s {
	case SevInfo.String():
		return SevInfo
	case SevWarn.String():
		return SevWarn
	case SevError.String():
		return SevError
	}
	return SevError
}

func (s Severity) String() string {
	return [...]string{"error", "warn", "info"}[s]
}

func (s *Severity) Colorize() string {
	switch *s {
	case SevInfo:
		return color.GreenString(s.String())
	case SevWarn:
		return color.YellowString(s.String())
	case SevError:
		return color.RedString(s.String())
	}
	return color.RedString(s.String())
}

// compile-time check that Severity satisfies the yaml Unmarshaler
var _ yaml.Unmarshaler = (*Severity)(nil)

// UnmarshalYAML to unmarshal severity string
func (s *Severity) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	*s = NewSeverity(str)

	return nil
}
