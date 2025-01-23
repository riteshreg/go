package flowcontrol

import (
	"fmt"
	"runtime"
)

func Os() {
	fmt.Println("Go runs on")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("windows")
	}
}
