// Slices are like references to arrays
// A slice does not store any data, it just describes a section of an underlying array.

// Changing the elements of a slice modifies the corresponding elements of its underlying array.

// Other slices that share the same underlying array will see those changes.

package moretype

import "fmt"

func SlicePointers() {
	var names = [5]string{"Ritesh", "Nayan", "Nabin", "Raut", "Bikrant"}
	a1 := names[0:2] //[ritesh, nayan]
	a2 := names[1:4] //[nayan, nabin, raut]
	a1[0] = "Hagu"   // this will change names[0] to hagu
	fmt.Println(a1, a2, names)
}
