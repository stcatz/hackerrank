package main
import "fmt"

func main() {
	var s int
	fmt.Scanln(&s)

	result :=  make(map[int]int, 100)

	for i := 0; i < s; i++ {
		var input int
		fmt.Scanf("%d", &input)
		result[input] += 1
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", result[i])
	}

}
