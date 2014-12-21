package main
import "fmt"
import "strings"

func partition(x []int, flag *int) ([]int, int) {
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
	*flag = len(small)
	return x, len(small)
}


func quickSort(x []int) []int {
	if(len(x) < 2){
		return nil
	}

	var flag int
	tmp, _ := partition(x, &flag)

	copy(x, tmp)

	quickSort(x[0:flag])
	quickSort(x[flag+1:len(x)])
	fmt.Println(strings.Trim(fmt.Sprint(x), "[]"))
	return x
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

	quickSort(ar)
}