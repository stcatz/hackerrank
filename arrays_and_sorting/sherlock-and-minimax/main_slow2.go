package main

import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"

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
	fmt.Scanf("%d %d\n", &p, &q)

	max_value, m := 0, p

	for i := p; i <= q; i++ {
		//fmt.Println(a, i, min_minus(a, i))
		if min_minus(a, i) > max_value {
			max_value = min_minus(a, i)
			m = i
		}else if min_minus(a, i) == max_value {
			if m > i {
				m = i
			}
		}
		
	}

	fmt.Println(m)


}