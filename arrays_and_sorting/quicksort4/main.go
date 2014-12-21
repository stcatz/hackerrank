package main
import "fmt"

func partition(x []int) (int, int) {
	key := x[len(x)-1]

	availabel, swap := 0, 0

	for i := 0; i < len(x)-1; i++ {
		if x[i] < key {
			availabel = i+1
			swap += 1
			continue
		}else{
			for k := i; k < len(x)-1; k++ {
				if x[k] > key {
					continue
				}else{
					x[i], x[k] = x[k], x[i]
					swap += 1
					availabel += 1
					break	
				}
				
			}
		}
		
	}

	x[availabel], x[len(x) -1] = x[len(x) -1], x[availabel]
	swap += 1

	//fmt.Println(strings.Trim(fmt.Sprint(x), "[]"))

	return availabel, swap
}

func quickSort(x []int, l []int) int {
	if(len(x) < 2){
		return 0
	}

	flag, swap := partition(x)

	swap_1 := quickSort(x[0:flag], l)
	swap_2 := quickSort(x[flag+1:len(x)], l)

	return swap + swap_1 + swap_2
}

func calcInsert(x []int) int {
	result := 0

	for i := 1; i < len(x); i++ {
		j := i-1
		key := x[i]

		for j>=0 && x[j]>key {
			x[j+1] = x[j]
			result += 1
			j -= 1
		}
		x[j+1] = key
		
	}
	return result

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
	ai := make([]int, len(ar))

	copy(ai, ar)

	fmt.Println(calcInsert(ai) - quickSort(ar, ar))
}