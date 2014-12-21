package main
import "fmt"
import "strings"

func partition(x []int) error {
	i,j := 0, len(x)-1
	key := x[0]

	for i < j {
		for x[j] > key{
			j--
		}
		x[i], x[j] = x[j], x[i]

		for x[i] < key{
			i++
		}
		x[i], x[j] = x[j], x[i]
	}

	return nil
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

	partition(ar)

	fmt.Println(strings.Trim(fmt.Sprint(ar), "[]"))
}