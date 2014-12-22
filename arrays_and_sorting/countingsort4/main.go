package main
import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"

func main() {
	var s int
	fmt.Scanln(&s)

	result :=  make(map[int][]string, 100)

	in := bufio.NewReader(os.Stdin)
	for i := 0; i < s; i++ {
		line, _ := in.ReadString('\n')
		index, _ := strconv.Atoi(strings.Split(line, " ")[0])
		//strings.Join(result[index], strings.Split(line, " ")[1])
		var str string
		str = strings.Replace(strings.Split(line, " ")[1], "\n", "", -1)

		if i<s/2 {
			str = string('-')
		}

		result[index] = append(result[index], str)
	}

	var str string
	for i := 0; i < len(result); i++ {
		str += strings.Join(result[i], " ")
		str += " "
	}
	fmt.Printf("%s", str)

}
