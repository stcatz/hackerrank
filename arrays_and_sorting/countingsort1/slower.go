package main
import "fmt"

func main() {
	var s int
	fmt.Scanln(&s)

	ar := []int{}

	for i := 0; i < s; i++ {
		var input int
		fmt.Scanf("%d", &input)
		ar = append(ar, input)
	}

	result :=  make(map[int]int, 100)

	for i := 0; i < len(ar); i++ {
		result[ar[i]] += 1
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", result[i])
	}

}
