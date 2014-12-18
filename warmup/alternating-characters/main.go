package main
import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		var x string
		fmt.Scanln(&x)
		prev, count := 0, 0
		for j := 1; j < len(x); j++ {
			if x[prev] == x[j] {
				count++
			}
			prev++
		}
		fmt.Println(count)

	}
}