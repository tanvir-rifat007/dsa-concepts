// leetcode 3516
package main

import "fmt"

func findClosest(x, y, z int) int {
	var resX, resY int

	if z < x {
		resX = x - z

	} else if z > x {
		resX = z - x
	}

	if z < y {
		resY = y - z

	} else if z > y {
		resY = z - y
	}

	if resX < resY {
		return 1
	} else if resY < resX {
		return 2
	} else {
		return 0
	}

}

func main() {
	fmt.Printf("%d\n", findClosest(1, 5, 3))

}
