#!/usr/bin/perl

use Inline C;

sub main {
	print "stair2\tstair3\n";
	for (my $i = 1; $i <= 30; ++$i) {
		print "$i\t".stair2($i)."\t".stair3($i)."\n";
	}
}

&main();

__END__
__C__


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
	//for ( int i = 4; i <= n; ++i) {
	int i = 4;
	for ( i = 4; i <= n; ++i) {
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
	//for ( int i = 3; i <= n; ++i) {
	int i = 3;
	for ( i = 3; i <= n; ++i) {
		f = f1 + f2;
		f2 = f1;
		f1 = f;
	}
	return f;
}

