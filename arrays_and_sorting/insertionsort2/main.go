package main
import "fmt"
import "strings"

func main() {
	var s int
	fmt.Scanln(&s)

	ar := []int{}

	for i := 0; i < s; i++ {
		var input int
		fmt.Scanf("%d", &input)
		ar = append(ar, input)
		
	}

	for i := 1; i < len(ar); i++ {
		for j := 0; j < i; j++ {
			if ar[j] < ar[i]{
				continue
			}else{
				value := ar[i]
				for k := i-1; k >= j; k-- {
					ar[k+1] = ar[k]
				}
				ar[j] = value
			}
			
		}

		fmt.Println(strings.Trim(fmt.Sprint(ar), "[]"))
		
	}
}