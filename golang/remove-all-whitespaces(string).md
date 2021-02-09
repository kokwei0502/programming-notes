## Remove all white spaces and line breaks

### Function to remove white spaces and line breaks

```golang
package main

import (
    "fmt"
    "unicode"
)

// This is the main function to remove all white spaces
// including line breaks (still don't understand how to achieve it)
func removeSpace(s string) string {
    rr := make([]rune, 0, len(s))
    for _, r := range s {
        if !unicode.IsSpace(r) {
            rr = append(rr, r)
        }
    }
    return string(rr)
}

func main() {
    s := "I d skd a efju N"
    fmt.Println(s) // output: I d skd a efju N
    s = removeSpace(s)
    fmt.Println(s) // output: IdskdaefjuN
}
```

### Strings.TrimSpace

```golang
package main

import (
	"fmt"
	"strings"
)

func main() {
	txt := "  ABCD Efg hij klmn o pqr st"
	// TrimSpace returns a slice of the string s, with all leading and trailing white space removed
	fmt.Println(strings.TrimSpace(txt)) // output: ABCD Efg hij klmn o pqr st
}

```
