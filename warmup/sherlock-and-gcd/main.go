package main
import "fmt"
//import "strings"

func GCD(a, b int) int {
        for b != 0 {
                a, b = b, a%b
        }
        return a
}

func main() {
	var t int
	fmt.Scanln(&t)

	//fmt.Println(strings.Repeat("na", 2))

	for i := 0; i < t; i++ {
		var n int
		fmt.Scanln(&n)
		input := []int{}
		for j := 0; j < n; j++ {
			var x int
			fmt.Scanf("%d", &x)
			input = append(input, x)
		}

		flag := false

		for j := 0; j < len(input); j++ {
			if input[j]==1 {
				flag = true
				continue
			}
			for k := j+1;  k < len(input); k++ {
				if GCD(input[j], input[k]) == 1 {
					flag = true
				}
			}
			
		}

		if flag {
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}

	}

}