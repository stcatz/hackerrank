package main
import "fmt"
//import "strings"

func partition(x []int, pos int) int {

	if(len(x) < 2){
		return x[0]
	}

	i,j := 0, len(x)-1
	key := x[len(x)-1]

	for i < j {
		for x[i] < key{
			i++
		}
		x[i], x[j] = x[j], x[i]

		for x[j] > key{
			j--
		}
		x[i], x[j] = x[j], x[i]
	}


	pivot := i
	n := pivot - pos
	if n == 0 {
		return x[pivot]
	}else if n > 0 {
		return partition(x[0:pivot], pos)
	}else{
		return partition(x[pivot:len(x)], -n)
	}
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

	result := partition(ar, len(ar)/2)

	fmt.Println(result)
}