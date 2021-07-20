// go fmtする前
package main

import (
"math/rand"
"fmt"
)
type T struct {
	Id int
	Name string
}

func main() {
i := rand. Intn(10)
     fmt.Println( i)
}

// go fmt	したあと
package main

import (
	"fmt"
	"math/rand"
)

type T struct {
	Id   int
	Name string
}

func main() {
	i := rand.Intn(10)
	fmt.Println(i)
}
