#include <iostream>
#include <chrono>
using namespace std;

void heavyComputation(){
    int sum = 0;
    for(int i=0; i<=1e9; i++){
        sum += i;
    }
    cout << "Sum: " << sum << endl;
}

int main(){
    auto start = chrono::high_resolution_clock::now();

    heavyComputation();

    auto end = chrono::high_resolution_clock::now();
    auto elapsed = chrono::duration_cast<chrono::milliseconds>(end - start);
    cout << "Elapsed " << elapsed.count() << " milliseconds\n";

    return 0;
}
