### Example 1 :

> Find all text in "( )"

```golang
package main

import (
    "regexp"
    "fmt"
)

func main() {
    str := "foo(bar)foo(baz)golang"
    rex := regexp.MustCompile(`\(([^)]+)\)`)
    out := rex.FindAllStringSubmatch(str, -1)

    for _, i := range out {
        fmt.Println(i[1])
    }
}
```

output :

```
bar
baz
```

```golang
package main

import (
	"fmt"
	"regexp"
)

var rgx = regexp.MustCompile(`\((.*?)\)`)

func main() {
	s := `ad(tag)SomeText`
	rs := rgx.FindStringSubmatch(s)
	fmt.Println(rs[1])
}
```

output :

```
tag
```

```golang
package main

import (
	"fmt"
	"os"
	"regexp"
)

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	emails := []string{
		"brown@fox",
		"brown@fox.",
		"brown@fox.com",
		"br@own@fox.com",
	}

	pattern := `^\w+@\w+\.\w+$`
	for _, email := range emails {
		matched, err := regexp.Match(pattern, []byte(email))
		check(err)
		if matched {
			fmt.Printf("√ '%s' is a valid email\n", email)
		} else {
			fmt.Printf("X '%s' is not a valid email\n", email)
		}
	}
}

```

output :

```
X 'brown@fox' is not a valid email
X 'brown@fox.' is not a valid email
√ 'brown@fox.com' is a valid email
X 'br@own@fox.com' is not a valid email
```
