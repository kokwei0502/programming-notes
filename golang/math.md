## math

### 3 methods to convert float to designated decimal places

```golang
package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {
	x := 12.3456
	fmt.Println(math.Floor(x*100) / 100) // 12.34 (round down)
	fmt.Println(math.Round(x*100) / 100) // 12.35 (round to nearest)
	fmt.Println(math.Ceil(x*100) / 100)  // 12.35 (round up)
}
```

output :

```
12.34
12.35
12.35
```
