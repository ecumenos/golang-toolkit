package randomtools

import (
	"fmt"
)

func ExampleGetRandBool() {
	result, _ := GetRandBool(1)
	fmt.Print(result)
	// Output:
	// true
}
