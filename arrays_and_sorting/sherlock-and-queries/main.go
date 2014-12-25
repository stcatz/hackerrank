package main
import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	a :=  make([]int64, n)
	mark := make([]int64, n)
	b :=  make([]int64, m)
	c :=  make([]int64, m)
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = line[:len(line)-1]

	splited := strings.Split(line, " ")

	for i:=0; i<n; i++ {
        temp, _ := strconv.Atoi(splited[i])
        a[i] = int64(temp)
        mark[i] = -1

	}

	line, _ = in.ReadString('\n')
	line = line[:len(line)-1]
	splited = strings.Split(line, " ")

	for i:=0; i<m; i++ {
        temp, _ := strconv.Atoi(splited[i])
        b[i] = int64(temp)
	}

	line, _ = in.ReadString('\n')

	splited = strings.Split(line, " ")

	for i:=0; i<m; i++ {
        temp, _ := strconv.Atoi(splited[i])
        c[i] = int64(temp)
	}

	for i := 0; i < m; i++ {
		if(mark[b[i]-1] == -1){
			mark[b[i]-1] = c[i]
		}else{
			mark[b[i]-1] = (mark[b[i]-1] * c[i]) % 1000000007
		}
	}



	for i := 1; i <= n; i++ {
		for j := 1; (j * i) <=n ; j++ {
			if mark[i-1] != -1{
				a[i*j-1] = (a[i*j-1] * mark[i-1]) % 1000000007
			}
			
		}
		
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
		
	}
}