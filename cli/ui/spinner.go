package ui

import (
	"time"

	"github.com/briandowns/spinner"
)

func DefaultSpinner(text string) *spinner.Spinner {
	sp := spinner.New(spinner.CharSets[21], 100*time.Millisecond)
	sp.Color("green")
	sp.Suffix = " " + text

	return sp
}
