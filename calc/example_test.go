package calc

import "fmt"

// Example functions also count toward testing examples; they are useful for documentation.
func ExampleAdd() {
	fmt.Println(Add(1, 2))
	// Output: 3
}

func ExampleDiv_error() {
	if _, err := Div(1, 0); err != nil {
		fmt.Println("error")
	}
	// Output: error
}
