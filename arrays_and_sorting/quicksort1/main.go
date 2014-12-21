package main
import "fmt"
import "strings"

func partition(x []int) []int, int {
	key := x[0]

	small, big := []int{}, []int{}

	for i := 0; i < len(x); i++ {
		if(x[i] < key){
			small = append(small, x[i])
		}else{
			big = append(big, x[i])
		}
	}

	x = append(small, big...)
	return x, len(small)
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

	fmt.Println(strings.Trim(fmt.Sprint(partition(ar)), "[]"))
}