package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

func Spinner(suffix, endMsg string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " " + suffix
	s.FinalMSG = endMsg

	return s
}
