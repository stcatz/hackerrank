package main

import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"
import "sort"

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
/*
    //local code
	var p, q int
	fmt.Scanf("%d %d\n", &p, &q)
*/
	sort.Ints(a)

	max_value, m := 0, 0

	i, j := 0, n-1
	for ; a[i] < p; i++ {
	}

	for ; a[j] > q; j-- {
	}

	for k := i; k < j; k++ {
		if int((a[k+1] - a[k])/2) > max_value {
			max_value = int((a[k+1] - a[k])/2)
			m = int((a[k+1] + a[k])/2)
		}
	}

	left_min, left_m := 0, 0

	for k := p; k < a[i]; k++ {
		every_min := 0

		if i == 0 {
			every_min = a[i] - k
		}else{
			if k-a[i-1] > a[i] - k{
				every_min = a[i] - k
			}else{
				every_min = k-a[i-1]
			}
		}

		if every_min > left_min {
			left_min = every_min
			left_m = k
		}
	}

	if left_min >= max_value {
		max_value = left_min
		m = left_m
	}

	right_m, right_min := 0,0 

	for k := a[j]; k <= q; k++ {
		every_min := 0

		if j == n-1 {
			every_min = k - a[j]
		}else{
			if a[j+1] - k > k - a[j]{
				every_min =  k - a[j]
			}else{
				every_min = a[j+1] - k
			}
		}

		if every_min > right_min {
			right_min = every_min
			right_m = k
		}
	}

	if right_min > max_value {
		max_value = right_min
		m =right_m
	}

	fmt.Println(m)


}