/*
Named return values
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
*/

package components

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // the output will be x=4 and y=6

}
