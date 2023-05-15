package TOS

import (
	"fmt"
	"runtime"
)

const (
	Linux = iota
	Windows
	MacOS
	Unknown
)

func GetOSName() int {
	os := runtime.GOOS
	if os == "windows" {
		fmt.Println("Windows operating system detected")
		return Windows
	} else if os == "linux" {
		fmt.Println("Linux operating system detected")
		return Linux
	} else if os == "darwin" {
		fmt.Println("Mac operating system detected")
		return MacOS
	}
	fmt.Printf("Unknown operating system: %s\n", os)
	return Unknown
}
