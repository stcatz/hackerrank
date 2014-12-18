#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
#include <iterator>  
using namespace std;

int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */
    int num;
    cin >> num;
    vector<unsigned long long> *fibo = new vector<unsigned long long>(0);
    fibo->push_back(0);
    fibo->push_back(1);

	vector<unsigned long long>::iterator first = fibo->begin();
	vector<unsigned long long>::iterator second = first + 1;

    for(int i=0; i<num; i++){
    	unsigned long long x;
    	cin >> x;

    	if(x<4){cout << "IsFibo" << endl; continue;}
    	bool flag = false;

    	while(*second < x){
    		unsigned long long next_var = *first + *second;

    		if(next_var == x){flag=true;}
    		fibo->push_back(next_var);

            second = prev(fibo->end());
            first = prev(second);
    	}

    	if(!flag){
    		for(vector<unsigned long long>::iterator it=fibo->begin(); it!=fibo->end(); it++){
    			//cout << *it << endl;
    			if(*it == x){flag=true;}
    		}
    	}

    	if(flag){
    		cout << "IsFibo" << endl;
    	}else{
    		cout << "IsNotFibo" << endl;
    	}


    }
    return 0;
}