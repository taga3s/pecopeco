package util

import "github.com/fatih/color"

func TextGreen() *color.Color {
	c := color.New(color.FgGreen)
	return c
}

func TextBlue() *color.Color {
	c := color.New(color.FgHiBlue)
	return c
}
