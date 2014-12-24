package main
import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"
import "sort"

func main() {
	var n, k int
	fmt.Scanln(&n, &k)

	ar := make([]int, n)

	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	for i, k := range strings.Split(line, " "){
		ar[i], _ = strconv.Atoi(k)
	}

	sort.Ints(ar)
	result, count := 0, 0

	for i := 0; result <= k; i++ {
		result += ar[i]
		count++
	}


	fmt.Println(count-1)
}
