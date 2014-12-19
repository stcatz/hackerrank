package main
import "fmt"
import "strings"

func main() {
	var s int
	fmt.Scanln(&s)

	ar := []int{}

	for i := 0; i < s; i++ {
		var input int
		fmt.Scanf("%d", &input)
		ar = append(ar, input)
		
	}

	value := ar[len(ar)-1]
	for last, found := len(ar)-2, false; last >=0 && !found ; last-- {
		if ar[last] > value{
			ar[last+1] = ar[last]
		}else{
			ar[last+1] = value
			found = true
		}

		fmt.Println(strings.Trim(fmt.Sprint(ar), "[]"))

		if(last==0 && ar[last] > value ){
			ar[0] = value
			fmt.Println(strings.Trim(fmt.Sprint(ar), "[]"))
		}

	}
}