package main
import "fmt"
import "sort"

func main() {
	var n, k int
	fmt.Scanln(&n)
	fmt.Scanln(&k)

	candies := []int{}

	for i := 0; i < n; i++ {
		var x int
		fmt.Scanln(&x)
		candies = append(candies, x)
	}

	sort.Ints(candies)

	min := candies[k-1] - candies[0]

	for i := 0; i < n-k; i++ {
		if candies[i+k-1] - candies[i] < min {
			min = candies[i+k-1] - candies[i]
		}
	}

	fmt.Println(min)
}