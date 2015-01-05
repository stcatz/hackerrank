package main
import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"

func main() {
	var n, k int
	fmt.Scanln(&n, &k)

	ar := make([]int, n)

	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	for i, k := range strings.Split(line, " "){
		ar[i], _ = strconv.Atoi(k)
	}

	

	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", result[i])
	}

}
