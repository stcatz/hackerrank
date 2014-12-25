package main
import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	a :=  make([]uint64, n)
	b :=  make([]uint64, m)
	c :=  make([]uint64, m)
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = line[:len(line)-1]

	splited := strings.Split(line, " ")

	for i:=0; i<n; i++ {
        temp, _ := strconv.Atoi(splited[i])
        a[i] = uint64(temp)
	}

	line, _ = in.ReadString('\n')
	line = line[:len(line)-1]
	splited = strings.Split(line, " ")

	for i:=0; i<m; i++ {
        temp, _ := strconv.Atoi(splited[i])
        b[i] = uint64(temp)
	}

	line, _ = in.ReadString('\n')
	splited = strings.Split(line, " ")

	for i:=0; i<m; i++ {
        temp, _ := strconv.Atoi(splited[i])
        c[i] = uint64(temp)
	}


	for i := 0; i < m; i++ {
		c[i] = c[i] % 1000000007
		factor := b[i]
		for j:=1; ;j++{
			index := factor * uint64(j)
			if index > uint64(n) {
				break
			}
			a[index-1] = (a[index-1] * c[i])% 1000000007
		}
		
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
		
	}
}