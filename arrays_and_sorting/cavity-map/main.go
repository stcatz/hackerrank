package main
import "fmt"
//import "strings"
//import "bufio"
//import "os"
import "strconv"

func main() {
	var n int
	fmt.Scanln(&n)

	ar := make([][]int, n)

	for i := 0; i < n; i++ {
		var line string
		fmt.Scanln(&line)

		ar[i] = make([]int, n)
		for j, k := range(line) {
			ar[i][j], _ = strconv.Atoi(string(k))
		}

	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			if ar[i+1][j] < ar[i][j] && ar[i-1][j] < ar[i][j] && ar[i][j+1] < ar[i][j] && ar[i][j-1] < ar[i][j] {
				ar[i][j] = 10
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if ar[i][j] != 10 {
				fmt.Printf("%d",ar[i][j])
			}else{
				fmt.Printf("%s", "X")

			}
		}
		fmt.Println("")
	}
}