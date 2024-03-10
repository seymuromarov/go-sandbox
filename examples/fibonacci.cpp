#include <iostream>
using namespace std;

int main() {
    int number = 10, n1 = 0, n2 = 1, nextTerm;

    cout << "Fibonacci Series in C++:" << endl;

    for (int i = 1; i <= number; ++i) {
        cout << n1 << endl;
        nextTerm = n1 + n2;
        n1 = n2;
        n2 = nextTerm;
    }

    return 0;
}