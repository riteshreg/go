// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128

// 1 byte => 8bit => 0 to 2^8 - 1 => 0 - 255

package gotype

import "fmt"

func BasicType() {
	var variable uint16 = 255
	// 16 bit => 0 to 2^16 - 1 => 0 - 65535
	// 32 bit => 0 to 2^32 - 1 => 0 - 4,294,967,295
	var comp complex64 = -444

	fmt.Println(variable, comp)
}
