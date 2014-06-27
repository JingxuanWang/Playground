#include <iostream>

using namespace std;

int lis(const int * array, int length) {
	int sum = 0;
	int max_sum = 0;
	for (int i = 0; i < length; i++) {
		if (array[i] + sum < array[i]) {
			sum = array[i];
		}
		else {
			sum += array[i];
		}
		if (sum > max_sum) {
			max_sum = sum;
		}
	}
	return max_sum;
}

int main() {
	int array[9] = {1, -2, 3, 4, -6, 11, -2, 13, -20};
	cout<<lis(array, 9)<<endl;
}
