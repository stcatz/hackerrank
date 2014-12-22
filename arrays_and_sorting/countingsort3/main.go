package main
import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"

func main() {
	var s int
	fmt.Scanln(&s)

	result :=  make(map[int]int, 100)

	for i := 0; i < s; i++ {
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')
		i, _ := strconv.Atoi(strings.Split(line, " ")[0])
		result[i] += 1
	}


	for i := 0; i < 100; i++ {
		final := 0
		for j :=0; j<=i; j++ {
			final += result[j]
		}
		fmt.Printf("%d ", final)
	}

}
