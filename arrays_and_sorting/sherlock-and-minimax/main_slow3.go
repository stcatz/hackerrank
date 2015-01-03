package main

import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"
import "sort"

func abs(x int) int {
    switch {
    case x < 0:
        return -x
    case x == 0:
        return 0 // return correctly abs(-0)
    }
    return x
}

func min_minus(a []int, num int) int {
	result := abs(a[0] - num)
	for i := 0; i < len(a); i++ {
		temp := abs(a[i] - num)
		if temp < result {
			result = temp
		}
	}

	return result
}

func main() {
	var n int
	fmt.Scanln(&n)

	a :=  make([]int, n)
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = line[:len(line)-1]
	splited := strings.Split(line, " ")

	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(splited[i])
	}

	var p, q int
    line, _ = in.ReadString('\n')
    splited = strings.Split(line, " ")
    p, _ = strconv.Atoi(splited[0])
    q, _ = strconv.Atoi(splited[1])

	sort.Ints(a)

	max_value, min_value, m := 0, 0, 0

	last_pos, j := 0, 0

	for i := p; i <= q; i++ {
		for j = last_pos;j < n && a[j] < i; j++ {
		}

		last_pos = j

		if j==0{
			min_value = abs(a[j] - i)
		}else if j == n{
			min_value = abs(a[j-1] - i)
		}else{
			if abs(a[j] - i) > abs(a[j-1] - i) {
				min_value = abs(a[j-1] - i)
			}else{
				min_value = abs(a[j] - i)
			}
		}

		if min_value > max_value {
			max_value = min_value
			m = i
		}
		
	}


	fmt.Println(m)


}