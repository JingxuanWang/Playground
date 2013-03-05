#include <iostream>

using namespace std;

// stair 3 
// recursive method
// time complexity O(2^n)
// space complexity O(1)
int stair3_recursive(int n) {
	if (n <= 0) {
		return 0;
	}
	if (n == 1) {
		return 1;
	}
	if (n == 2) {
		return 2;
	}
	if (n == 3) {
		return 4;
	}
	return stair3_recursive(n - 1) + stair3_recursive(n - 2) + stair3_recursive(n - 3);
}

// stair 3 
// dynamic programming method
// time complexity O(n)
// space complexity O(1)
int stair3(int n) {
	if (n <= 0) {
		return 0;
	}
	if (n == 1) {
		return 1;
	}
	if (n == 2) {
		return 2;
	}
	if (n == 3) {
		return 4;
	}

	int f1 = 4; // for f(n-1)
	int f2 = 2; // for f(n-2)
	int f3 = 1; // for f(n-3)
	int f = f1 + f2;
	for ( int i = 4; i <= n; ++i) {
		f = f1 + f2 + f3;
		f3 = f2;
		f2 = f1;
		f1 = f;
	}
	return f;
}

// stair 2 
// recursive method
// time complexity O(2^n)
// space complexity O(1)
int stair2_recursive(int n) {
	if (n <= 0) {
		return 0;
	}
	if (n == 1) {
		return 1;
	}
	if (n == 2) {
		return 2;
	}
	return stair2_recursive(n - 1) + stair2_recursive(n - 2);
}

// stair 2 
// dynamic programming method
// time complexity O(n)
// space complexity O(1)
int stair2(int n) {
	if (n <= 0) {
		return 0;
	}
	if (n == 1) {
		return 1;
	}
	if (n == 2) {
		return 2;
	}

	int f1 = 2; // for f(n-1)
	int f2 = 1; // for f(n-2)
	int f = f1 + f2;
	for ( int i = 3; i <= n; ++i) {
		f = f1 + f2;
		f2 = f1;
		f1 = f;
	}
	return f;
}

int main() {

	cout<<"n\tstair2\tstair3"<<endl;
	for (int i = 1; i <= 30; ++i) {
		cout<<i<<"\t"<<stair2(i)<<"\t"<<stair3(i)<<endl;
		
		// compare dynamic programming to recursive method
		// while n > 32, we may notice that 
		// recursive method is far slower than dynamic programming method
		//cout<<"Recur:\t"<<stair2_recursive(i)<<"\t"<<stair3_recursive(i)<<endl;
	}
}
