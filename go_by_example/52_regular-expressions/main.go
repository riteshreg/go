package main

import (
	"fmt"
	"regexp"
)

func main() {

	// we are matching a string with regex
	// this regex will return true if the string start with r and should in between a-z and end with sh at last
	match, _ := regexp.MatchString("r([a-z]+)sh", "ritesh")
	fmt.Println(match)

	r, _ := regexp.Compile("k([a-z]+)ka")
	fmt.Println(r.MatchString("khadka"))

	fmt.Println(r.FindString("ritesh khadka"))
	fmt.Println(r.FindStringIndex("ritesh khadka"))
	fmt.Println(r.FindStringSubmatch("ritesh khadka"))

}
