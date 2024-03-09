package main

import (
	"fmt"
	"go-examples/framework"
)
import "rsc.io/quote"

func main() {
	fmt.Printf(quote.Go())

	greet := framework.Greet("Kichu")

	fmt.Printf("\n %s", greet)
}
