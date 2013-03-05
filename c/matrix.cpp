#include <iostream>

using namespace std;

// print original matrix
void printMatrix(int **a, int x, int y) {
	for (int i = 0; i < y; ++i) {
		for (int j = 0; j < x; ++j) {
			cout<<a[i][j]<<'\t';
		}
		cout<<endl;
	}
}

// solution 1
// recursive
void printMatrix2(int **a, int bx, int by, int x, int y) {
	int i = 0;
	int j = 0;
	for (j = bx; j < bx + x; ++j) {
		cout<<a[by][j]<<" ";
	}
	for (i = by + 1; i < by + y; ++i) {
		cout<<a[i][bx + x - 1]<<" ";
	}
	if (y > 1) {
		for (j = bx + x - 2; j >= bx; --j) {
			cout<<a[by + y - 1][j]<<" ";
		}
	}
	if (x > 1) {
		for (i = by + y - 2; i > by; --i) {
			cout<<a[i][bx]<<" ";
		}
	}
	if (x > 2 && y > 2) {
		printMatrix2(a, bx + 1, by + 1, x - 2, y - 2);
	}
}

// solution 1
// loop
void printMatrix3(int **a, int bx, int by, int x, int y) {
	int i = 0;
	int j = 0;
	
	while (x > 0 && y > 0) {
		
		for (j = bx; j < bx + x; ++j) {
			cout<<a[by][j]<<" ";
		}
		for (i = by + 1; i < by + y; ++i) {
			cout<<a[i][bx + x - 1]<<" ";
		}
		if (y > 1) {
			for (j = bx + x - 2; j >= bx; --j) {
				cout<<a[by + y - 1][j]<<" ";
			}
		}
		if (x > 1) {
			for (i = by + y - 2; i > by; --i) {
				cout<<a[i][bx]<<" ";
			}
		}

		// equal to recursive call
		++bx;
		++by;
		x -= 2;
		y -= 2;
	}
}

// solution 2
// recursive
void printMatrix4(int **a, int left, int up, int right, int down) {
	int i = 0;
	int j = 0;
	
	for (i = left; i <= right; ++i) {
		cout<<a[up][i]<<" ";
	}
	
	++up;
	
	for (j = up; j <= down; ++j) {
		cout<<a[j][right]<<" ";
	}
	
	--right;

	if (up <= down) {
		for (i = right; i >= left; --i) {
			cout<<a[down][i]<<" ";
		}
	
		--down;
	}

	if (left <= right) {
		for (j = down; j >= up; --j) {
			cout<<a[j][left]<<" ";
		}

		++left;
	}

	if (up <= down && left <= right) {
		printMatrix4(a, left, up, right, down);
	}
}

// solution 2
// loop
void printMatrix5(int **a, int left, int up, int right, int down) {
	int i = 0;
	int j = 0;

	while(up <= down && left <= right) {
		for (i = left; i <= right; ++i) {
			cout<<a[up][i]<<" ";
		}
		
		++up;
		
		for (j = up; j <= down; ++j) {
			cout<<a[j][right]<<" ";
		}
		
		--right;

		if (up <= down) {
			for (i = right; i >= left; --i) {
				cout<<a[down][i]<<" ";
			}
		
			--down;
		}

		if (left <= right) {
			for (j = down; j >= up; --j) {
				cout<<a[j][left]<<" ";
			}

			++left;
		}
	}
}


int main () {

	const int CASE_NUM = 5;

	int test_case[CASE_NUM][2] = {
		{3, 5},
		{5, 3},
		{3, 3},
		{4, 4},
		{5, 5}
	};

	// for each case
	for (int n = 0; n < CASE_NUM; ++n) {
		int x = test_case[n][0]; // col number of the matrix
		int y = test_case[n][1]; // row number of the matrix
		int c = 1;
		
		// alloc memory for a[y][x]
		int **a = new int* [y];
		for (int i = 0; i < y; ++i) {
			a[i] = new int [x];
			for ( int j = 0; j < x; ++j) {
				a[i][j] = c;
				c++;
			}
		}

		// print original matrix
		cout<<"==========CASE "<<n<<"========"<<endl;
		printMatrix(a, x, y);
		
		// print the matrix clockwisely
		cout<<"==========clock wise=========="<<endl;
		
		cout<<"Solution 1: recursive "<<endl;
		printMatrix2(a, 0, 0, x, y);
		cout<<endl;
		
		cout<<"Solution 1: loop "<<endl;
		printMatrix3(a, 0, 0, x, y);
		cout<<endl;

		cout<<"Solution 2: recursive "<<endl;
		printMatrix4(a, 0, 0, x - 1, y - 1);
		cout<<endl;
		
		cout<<"Solution 2: loop "<<endl;
		printMatrix5(a, 0, 0, x - 1, y - 1);
		cout<<endl;
	
		// free memory for a[y][x]
		for (int i = 0; i < y; ++i) {
			delete a[i];
		}
		delete a;
	}
	return 0;
}
