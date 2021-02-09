package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := `ab cd ef gh
	ijkl  nop 
        			qrs tuv      wx yz`
	fmt.Println(str)
	fmt.Println(removeSpace(str))
}

func removeSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}
