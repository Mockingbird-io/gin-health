package main

import "fmt"

type TwoInit struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInit)
	two1.a = 2
	two1.b = 3
	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInit{
		a: 20,
		b: 30,
	}
	fmt.Printf("The sum is: %d\n", two2.AddToParam(1))
}

func (tn *TwoInit) AddThem() int {
	return tn.a + tn.b
}
func (tn *TwoInit) AddToParam(param int) int {
	return tn.a + tn.b + param
}
