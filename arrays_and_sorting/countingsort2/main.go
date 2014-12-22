package main
import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"

func main() {
	var s int
	fmt.Scanln(&s)

	result :=  make(map[int]int, s)

	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	for _, x := range(strings.Split(line, " ")) {
		i, _ := strconv.Atoi(x)
		result[i] += 1
		
	}

	for i := 0; i < s; i++ {
		for j :=0; j<result[i]; j++ {
			fmt.Printf("%d ", i)
		}
	}

}
