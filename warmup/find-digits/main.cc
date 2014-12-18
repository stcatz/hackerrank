#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */
    int num;
    cin >> num;

    for(int i=0; i<num; i++){
    	int count=0;
    	int input;
    	cin >> input;
    	int left = input;

    	while(left>0) {

    		int divider = left % 10;
    		if(divider!=0 && input%divider == 0){
    			count++;
    		}
    		left = int(left/10);

    	};
    	cout << count << endl;

    }
    return 0;
}