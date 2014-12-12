package main
import "fmt"
import "strconv"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
	var num int
	fmt.Scanln(&num)

	for i:=0; i<num; i++ {
		var x string
		var count = 0
		fmt.Scanln(&x)
		total, _ := strconv.ParseInt(x, 0, 64)
		for _, ch := range x {
			ch_num, _ := strconv.ParseInt(string(ch), 0, 64)
			if( ch_num != 0 && total % ch_num == 0){
				count++
			}	
		}

		fmt.Println(count)

	}
    
}