#include <iostream>
#include <stdio.h>

using namespace std;

int binarySearch2(int * a, int t, int length) {
	if (NULL == a || length <= 0) {
		return -1;
	}

	int low = 0;
	int high = length - 1;
	
	while(low <= high) {
		int mid = (low + high) / 2;
		
		if (a[mid] == t) {
			return mid;
		} else if (a[mid] > t) {
			high = mid - 1;
		} else {
			low = mid + 1;
		}
	}
	return -1;
}

int binarySearch(int * a, int t, int low, int high) {
	if (NULL == a || low < 0 || high < 0 || low > high) {
		return -1;
	}

	int mid = (low + high) / 2;
	
	if (a[mid] == t) {
		return mid;
	} else if (a[mid] > t) {
		return binarySearch(a, t, low, mid - 1);
	} else {
		return binarySearch(a, t, mid + 1, high);
	}

	return -1;
}

int search(int *a, int t, int low, int high) {
	if (low > high) {
		return -1;
	}
	int mid = low + (high - low) / 2;
	if (a[low] <= a[mid]) {
		if (a[low] <= a[high]) {
			return binarySearch(a, t, low, high);
		} else {
			if (t >= a[low] && t <= a[mid]) {
				return binarySearch(a, t, low, mid);
			} else {
				return search(a, t, mid + 1, high);
			}
		}
	} else {
		if (t > a[mid] && t <= a[high]) {
			return binarySearch(a, t, mid + 1, high);
		} else {
			return search(a, t, low, mid);
		}
	}
}

int main() {
	int a[] = {1, 2, 3, 4, 5, 6, 7};
	for (int i = 0; i < 7; ++i) {
		cout<<a[i]<<'\t'<<binarySearch(a, a[i], 0, 6)<<endl;
	}
}



