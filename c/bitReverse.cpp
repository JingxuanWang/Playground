#include<stdio.h>
#include<iostream>

using namespace std;

int main() {

	for (int i = 0; i < 256; ++i) {
		int x = i;
		printf("%d: ", x);
		printf("%x\t", x);
		x = (x & 0x55) << 1 | (x & 0xAA) >> 1;
		printf("%x\t", x);
        x = (x & 0x33) << 2 | (x & 0xCC) >> 2;
		printf("%x\t", x);
		x = (x & 0x0F) << 4 | (x & 0xF0) >> 4;
		printf("%x\n", x);
	}
}
