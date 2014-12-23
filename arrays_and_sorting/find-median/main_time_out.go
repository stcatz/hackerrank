package main
import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"

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

func find_mid(x []int, pos int) int {
	if(len(x) < 2){
		return x[0]
	}

	pivot := partition(x)

	n := pivot - pos

	if n == 0 {
		return x[pivot]
	}else if n > 0 {
		return find_mid(x[0:pivot], pos)
	}else{
		return find_mid(x[pivot:len(x)], -n)
	}
}


func main() {
	var s int
	fmt.Scanln(&s)

	ar := make([]int, s)

	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	for i, k := range strings.Split(line, " "){
		ar[i], _ = strconv.Atoi(k)
	}

	result := find_mid(ar, len(ar)/2)
	fmt.Println(result)
	
	//fmt.Println(strings.Trim(fmt.Sprint(ar), "[]"))
}