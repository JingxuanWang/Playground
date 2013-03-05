#include <iostream>
#include <stdio.h>

using namespace std;

typedef float float32;
typedef int int32;

/// This is a approximate yet fast inverse square-root.
// WJXNOTE: return approximate sqrt(x)/x  for x
inline float32 b2InvSqrt(float32 x)
{
	union
	{   
		float32 x;
		int32 i;
	} convert;

	cout<<x<<'\t';
	
	convert.x = x;
	printf("%x\t",convert.i);
	float32 xhalf = 0.5f * x;

	cout<<float(0x5f3759df)<<'\t';
	cout<<convert.i<<'\t';
	cout<<(convert.i>>1)<<'\t';

	convert.i = 0x5f3759df - (convert.i >> 1); 
	
	cout<<convert.i<<'\t';
	
	x = convert.x;
	
	cout<<convert.x<<endl;

	x = x * (1.5f - xhalf * x * x); 
	return x;
};


int atoi(const char * str) {
	if (str == NULL) {
		return 0;
	}

	int minus = 1;
	int value = 0;
	if (*str == '-') {
		minus = -1;
		++str;
	}

	while(*str != '\0') {
		if (*str - '0' <= 9 && *str - '0' >= 0) {
			value *= 10;
			value += *str - '0';
		}
		++str;
	}
	return minus * value;
}

char * itoa(int value, char * str, int base) {
	char ta[33];
	char *p = ta;
	int minus = 0;
	if (value < 0) {
		minus = 1;
		value *= -1;
	}
	char tb[] = "0123456789abcdef";
	do {
		*p++ = tb[value % base];
		value /= base;
	} while(value > 0);
	if (minus) {
		*p = '-';
	} else {
		--p;
	}

	int i = 0;
	while(p >= ta) {
		str[i++] = *p--; 
	}
	str[i] = '\0';
	return str;
}

int main() {
	//for (int i = 0 ; i < 10; ++i) {
		//cout<<i<<'\t'<<b2InvSqrt(i)<<'\t'<<b2InvSqrt(i) * i<<endl;
		//b2InvSqrt(i);
	//}
	/*
	char atoi_case[][20] = {
		"-12345", 
		"+12345",
		"012345",
		"12345",
		"102345",
		"+-102345",
		"-+102345",
		"--+102345",
		"-1",
	};
	for (int i = 0; i < 9; ++i) {
		cout<<atoi(atoi_case[i])<<"\t";
		cout<<endl;
	}

	int itoa_case[20] = {
		0, 
		12345,
		12345,
		1024,
		-127,
		102345,
		102345,
		-102345,
		-102345,
		-1,
	};

	cout<<"=================="<<endl;

	char num[100];
	for (int i = 0; i < 9; ++i) {
		cout<<itoa(itoa_case[i], num, 16)<<"\t";
		cout<<endl;
	}

	//char buffer[33] = "";
	//cout<<itoa(345, buffer, 10)<<endl;
	//cout<<itoa(345, buffer, 8)<<endl;
	//cout<<itoa(345, buffer, 16)<<endl;
	//cout<<itoa(345, buffer, 2)<<endl;
	*/
}



