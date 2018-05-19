package main

import "fmt"

func eval(a, b int, opt string) (int, error) {
	switch opt {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("%s", opt)
	}
}

func div(a,b int) (q,r int) {
	return a / b, a % b
}


func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers  {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(4,5,"/"))
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(sum(1,2,3,4,5,6))
}
