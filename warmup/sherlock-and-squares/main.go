package main
import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Scanf("%d%d", &a, &b)

		result := 0
		p := 1
		for p*p <= b {
			if p*p >= a {
				result++
			}
			p++
		}
		fmt.Println(result)
	}
	
}