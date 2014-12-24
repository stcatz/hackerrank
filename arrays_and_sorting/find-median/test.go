package main
import "fmt"
//import "strings"


func partition(x []int) int {
	i,j := 0, len(x)-1
	key := x[j]

	for i < j {
		for x[i] <= key && i<j {
			i++
		}

		x[i], x[j] = x[j], x[i]

		for x[j] > key && i<j {
			j--
		}

		x[i], x[j] = x[j], x[i]

	}

	return j
}


func main() {
	var s int
	fmt.Scanln(&s)

	ar := []int{}

	for i := 0; i < s; i++ {
		var input int
		fmt.Scanf("%d", &input)
		ar = append(ar, input)
	}

	pivot := partition(ar)
	fmt.Println(ar, pivot)
}