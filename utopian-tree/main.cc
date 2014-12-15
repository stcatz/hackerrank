#include <iostream>
using namespace std;

int height(int n) {
    int result = 1;    
    for(int i=0; i<n; i++){
        if((i+1)%2==1){
            result *= 2;
        }else{
            result += 1;
        }
    }
    return result;
}
int main() {
    int T;
    cin >> T;
    while (T--) {
        int n;
        cin >> n;
        cout << height(n) << endl;
    }
}
