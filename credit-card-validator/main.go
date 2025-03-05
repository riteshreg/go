package main

import (
	"fmt"
	"strconv"
)

func luhn(card_number string) bool {

	var revert bool = false
	var sum int = 0

	for i := len(card_number) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(string(card_number[i]))
		if err != nil {
			return false
		}

		if revert {
			num = num * 2
			if num > 9 {
				num = num - 9
			}
		}

		sum += num
		revert = !revert
	}

	return sum%10 == 0

}

func main() {
	fmt.Println(luhn("3742454555400126"))
}
