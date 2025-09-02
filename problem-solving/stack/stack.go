// this implementation of stack using the array under the hood

package main

import "fmt"

type Stack struct {
	arr []int
}

func (s *Stack) push(elm ...int) []int {

	s.arr = append(s.arr, elm...)
	return s.arr

}

func (s *Stack) pop() []int {
	s.arr = s.arr[:len(s.arr)-1]

	return s.arr

}

func (s *Stack) read() int {
	return s.arr[len(s.arr)-1]

}

func main() {
	var stack Stack

	arr := stack.push(1, 2, 3, 4)

	fmt.Printf("After pushing : %v\n", arr)

	fmt.Printf("After poping the array becomes : %v\n", stack.pop())

	fmt.Printf("After poping the array becomes : %v\n", stack.pop())
	fmt.Printf("%d\n", stack.read())

}
