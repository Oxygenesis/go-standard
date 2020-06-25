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

func Discovered() {
	fmt.Printf("test")
}

func main() {
	fmt.Println("welcome to da jungle")
}
