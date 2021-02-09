## Remove all white spaces and line breaks

### Function to remove white spaces and line breaks

```go
package main

import (
    "fmt"
    "unicode"
)

// This is the main function to remove all white spaces
// including line breaks (still don't understand how to achieve it)
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
    // output :
    //     ab cd ef gh
    // ijkl nop
    // qrs tuv wx yz
	fmt.Println(removeSpace(str)) // output: abcdefghijklnopqrstuvwxyz
}

// Still don't know what logic
func removeSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}

```

output :

```
ab cd ef gh
ijkl nop
qrs tuv wx yz
abcdefghijklnopqrstuvwxyz
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
