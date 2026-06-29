package packagesandmodules

import "golang.org/x/exp/constraints"

// Type to define numbers
// Interface using constraints
type Number interface {
	constraints.Integer | constraints.Float
}

// Adds two numbers and returns the value
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](x, y T) T {
	return x + y
}
