#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;

int fibo(int index){
	if(index < 2){return index;}
	int first = 0, result = 0;
	int second = 1;
	for(int i=0; i<(index-1); i++){
		result = first + second;
		first = second;
		second = result;
	}
	return result;
}

string isFibo(int num){
	int result = 0;
	bool flag = false;
	for(int i=0; result < num; i++){
		result = fibo(i);
		if(result == num){
			flag = true;
		}
	}

	if(flag){
		return "IsFibo";
	}else{
		return "IsNotFibo";
	}
}


int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */
    int num;
    cin >> num;

    for(int i=0; i<num; i++){
    	int x;
    	cin >> x;

    	cout << isFibo(x) << endl;
    }
    return 0;
}