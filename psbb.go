package main

import (
	"fmt"
	"sort"
)

func Psbb(n int, fams []int) {
	if len(fams) != n {
		fmt.Println("Input must be equal with count of family")
		return
	}

	sort.Ints(fams)

	left := 0
	right := n - 1
	bus := 0

	for left <= right {

		//tersissa 1 fams
		if left == right {
			bus++
			break
		}

		//gabungkan fams terbesar n terkecil
		if fams[left]+fams[right] <= 4 {
			left++
		}

		right--

		bus++

	}

	fmt.Println("Minimum bus required is:", bus)
}

func main() {

	family := 5
	member := []int{1, 2, 4, 3, 3}
	Psbb(family, member)

	family1 := 8
	member1 := []int{2, 3, 4, 4, 2, 1, 3, 1}
	Psbb(family1, member1)

	family2 := 2
	member2 := []int{1, 5, 10}
	Psbb(family2, member2)

}
