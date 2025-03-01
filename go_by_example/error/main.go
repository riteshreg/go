package error

import (
	"errors"
	"fmt"
	"log"
)

func registerUser(fullName string, age int) (string, error) {
	if fullName == "" {
		return "", errors.New("full_name is null")
	}

	if age <= 0 {
		return "", errors.New("age can't not smaller than 1")
	}

	return fullName, nil
}

func f(num int) (int, error) {
	if num == 20 {
		return -1, errors.New("can't work with number")
	}

	return num + 5, nil

}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}

	return nil
}

func Error() {
	if fullName, err := registerUser("Ritesh Khadka", 21); err != nil {
		log.Fatalf("error registering user: %s", err)
	} else {
		fmt.Println("Successfully inserted", fullName)
	}

	if num, err := f(40); err != nil {
		log.Fatalf("error in f() %s", err)
	} else {
		fmt.Println("Successfully getten the value of num", num)

	}

	for i := 0; i <= 5; i++ {
		if err := makeTea(i); err != nil {
			log.Fatalf("error registering user: %s", err)
		} else {
			fmt.Println("Successfully maked tea")
		}
	}

}
