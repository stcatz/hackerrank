#include <cstdlib>
#include <climits>
#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;
 
vector< vector<int> > tree;
vector<bool> visited;
vector<int> val;
int totalVal;
 
int subSum(int root, int& minCut) {
    int sum = 0;
    visited[root] = true;
    for (int i = 0; i < tree[root].size(); i++) {
        int sub = tree[root][i];
        if (visited[sub])
            continue;
        int x = subSum(sub, minCut);
        sum += x;
    }
    sum += val[root];
    minCut = min(minCut, abs(totalVal - 2 * sum));
    return sum;
}
int main() {
    int N;
    cin >> N;
    tree.resize(N + 1);
    visited.resize(N + 1);
    val.resize(N + 1);
    totalVal = 0;
    for (int i = 1; i <= N; i++) {
        cin >> val[i];
        totalVal += val[i];
    }
    for (int i = 0; i < N - 1; i++) {
        int a, b;
        cin >> a >> b;
        tree[a].push_back(b);
        tree[b].push_back(a);
    }
    int minCut = INT_MAX;
    // pick 1 as root
    subSum(1, minCut);
    cout << minCut << endl;
    return 0;
}