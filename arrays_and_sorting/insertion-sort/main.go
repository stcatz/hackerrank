/***
Problem Statement

Insertion Sort is a simple sorting technique which was covered in previous challenges. Sometimes, arrays may be too large for us to wait around for insertion sort to finish. Is there some other way we can calculate the number of times Insertion Sort shifts each elements when sorting an array?

If ki is the number of elements over which ith element of the array has to shift then total number of shift will be k1 + k2 + ... + kN.

Input: 
The first line contains the number of test cases T. T test cases follow. The first line for each case contains N, the number of elements to be sorted. The next line contains N integers a[1],a[2]...,a[N].

Output: 
Output T lines, containing the required answer for each test case.

Constraints: 
1 <= T <= 5 
1 <= N <= 100000 
1 <= a[i] <= 1000000

Sample Input:

2  
5  
1 1 1 2 2  
5  
2 1 3 1 2
Sample Output:

0  
4   
Explanation 
First test case is already sorted, therefore there's no need to shift any element. In second case, it will proceed in the following way.

Array: 2 1 3 1 2 -> 1 2 3 1 2 -> 1 1 2 3 2 -> 1 1 2 2 3
Moves:   -        1       -    2         -  1            = 4
***/

package main
import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"


func merge(left, right []int) int {

  size, i, j := len(left)+len(right), 0, 0
  slice := make([]int, size, size)
  result := 0

  for k := 0; k < size; k++ {
    if i > len(left)-1 && j <= len(right)-1 {
      slice[k] = right[j]
      j++
    } else if j > len(right)-1 && i <= len(left)-1 {
      if i != len(left)-1{
      	//result += len(right)
      }
      slice[k] = left[i]
      i++
    } else if left[i] <= right[j] {
      slice[k] = left[i]
      i++
    } else {
      result += len(left) - i
      //fmt.Println(left, right)
      slice[k] = right[j]
      j++
    }
  }
  copy(left, slice[0:len(left)])
  copy(right, slice[len(left):len(right)+len(left)])
  return result
}

func merge_sort( ar []int) int {
	if len(ar) < 2 {
		return 0
	}

	result := 0

	middle := len(ar)/2

	result += merge_sort(ar[0:middle])
	result += merge_sort(ar[middle:len(ar)])
	//fmt.Println(ar[0:middle], ar[middle:len(ar)])
	result += merge(ar[0:middle], ar[middle:len(ar)])

	return result
}


func main() {

	var t int

	fmt.Scanln(&t)

	in := bufio.NewReader(os.Stdin)

	for ts := 0; ts < t; ts++ {
		var s int
		//fmt.Scanln(&s)
		line1, _ := in.ReadString('\n')
		line1 = line1[:len(line1)-1]
		s, _ = strconv.Atoi(line1)

		ar := make([]int, s)
		
		line, _ := in.ReadString('\n')

		//Necessary on the hackerrank due to different go version I think
		line = line[0:len(line)-1]

		for i, k := range strings.Split(line, " "){
			ar[i], _ = strconv.Atoi(k)
		}

		//fmt.Println(ar)
		//ans := 0

		result := merge_sort(ar)
		//fmt.Println(ar)
		fmt.Println(result)

	}

}