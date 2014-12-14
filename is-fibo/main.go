package main
import "fmt"

func isFibo(val uint64) bool{
	if val<4 {
		return true
	}else{
		var first, second, result uint64
		first,second,result = 0,1,0
		for (result < val) {
			result = first + second
			first = second
			second = result
		}
		if(result == val){
			return true
		}else{
			return false
		}
	}

}

func main() {
	var num int
	fmt.Scanln(&num)

	for i:=0; i<num; i++ {
		var x uint64
		fmt.Scanln(&x)
		result := isFibo(x)
		if result {
			fmt.Println("IsFibo")
		}else{
			fmt.Println("IsNotFibo")
		}
	}
}