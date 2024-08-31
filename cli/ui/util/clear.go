package util

import "fmt"

// clear terminal
func Clear() {
	fmt.Println("\033[H\033[2J")
}
