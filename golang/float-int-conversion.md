```golang
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64 = 5.51
	var y float64 = 5.50
	var z float64 = 5.49
	fmt.Println(int(math.Round(x))) // outputs "6"
	fmt.Println(int(math.Round(y))) // outputs "6"
	fmt.Println(int(math.Round(z))) // outputs "5"
}

```

output :

```
6
6
5
```
