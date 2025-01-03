package main

// we can even import like this
// import "fmt"
// import "math/rand"
// but go the go extension simply group the import also go extension remove the semicolon

import (
	"fmt"
	"math/rand"
)

func mail() {
	fmt.Println("random number", rand.Intn(10))
}
