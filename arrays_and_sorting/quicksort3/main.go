package main
import "fmt"
import "strings"

func partition(x []int) int {
	key := x[len(x)-1]

	availabel := 0

	for i := 0; i < len(x)-1; i++ {
		if x[i] < key {
			availabel = i+1
			continue
		}else{
			for k := i; k < len(x)-1; k++ {
				if x[k] > key {
					continue
				}else{
					x[i], x[k] = x[k], x[i]
					availabel += 1
					break
				}
				
			}
		}
		
	}

	x[availabel], x[len(x) -1] = x[len(x) -1], x[availabel]

	//fmt.Println(strings.Trim(fmt.Sprint(x), "[]"))

	return availabel
}

func quickSort(x []int, l []int) []int {
	if(len(x) < 2){
		return nil
	}

	flag := partition(x)
	fmt.Println(strings.Trim(fmt.Sprint(l), "[]"))


	quickSort(x[0:flag], l)
	quickSort(x[flag+1:len(x)], l)

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

	quickSort(ar, ar)
}