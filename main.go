package goStandard

import "fmt"

func UncoveredIf(a bool) bool {
	if a {
		return false
	}
	return true
}

func FullyCovered() bool {
	return true
}

func Uncovered() bool {
	return true
}

func moreUncovered() {
	fmt.Print("asd")
}
