// Stacking defers
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

// To learn more about defer statements read this blog post.
// https://go.dev/blog/defer-panic-and-recover

package flowcontrol

import "fmt"

func DefMulti() {
	fmt.Println("Starting")

	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("end")
}
