package main

import (
	"fmt"

	gomod4a_v1 "github.com/relax-space/gomod4a"
	gomod4a_v2 "github.com/relax-space/gomod4a/v2"
)

func main() {
	fmt.Printf("gomod4a_v1.AddA(1,2)=%d\n", gomod4a_v1.AddA(1, 2))
	fmt.Printf("gomod4a_v2.AddA(1,2)=%d\n", gomod4a_v2.AddA(1, 2))
}
