#!/usr/bin/perl -w

use strict;

my $a = ">";
my $b = "<";
my $s = "_";

my $num = 3;

my @init;
my @final;


for (my $i = 0; $i < $num; ++$i) {
	push(@init, $a);
	push(@final, $b);
}

push(@init, $s);
push(@final, $s);

for (my $i = 0; $i < $num; ++$i) {
	push(@init, $b);
	push(@final, $a);
}

sub check {
	my ($src, $dst) = @_;
	my $length = $num * 2 + 1;
	for (my $i = 0; $i < $length; ++$i) {
		if ($src->[$i] ne $dst->[$i]) {
			return 0;
		}	
	}
	return 1;
}

sub try {
	my ($stack, @status) = @_;

	my $length = $num * 2 + 1;
	
	# ending condition
	if (check(\@status, \@final)) {
		print "sucess\n";
		my $i = 1;
		for (@$stack) {
			print $i++, " ", @$_, "\n";
		}
		exit;
	}

	for (my $i = 0; $i < $length; $i ++) {
		if ($status[$i] eq $a && $i + 1 <= $length && $status[$i + 1] eq $s ) {
			($status[$i], $status[$i + 1]) = ($status[$i + 1], $status[$i]);
			push(@$stack, \@status);
			try($stack, @status);
			pop(@$stack);
			($status[$i + 1], $status[$i]) = ($status[$i], $status[$i + 1]);
		}
		if ($status[$i] eq $a && $i + 2 <= $length && $status[$i + 2] eq $s) {
			($status[$i], $status[$i + 2]) = ($status[$i + 2], $status[$i]);
			push(@$stack, \@status);
			try($stack, @status);
			pop(@$stack);
			($status[$i + 2], $status[$i]) = ($status[$i], $status[$i + 2]);
		}
		if ($status[$i] eq $b && $i - 1 >= 0 && $status[$i - 1] eq $s) {
			($status[$i], $status[$i - 1]) = ($status[$i - 1], $status[$i]);
			push(@$stack, \@status);
			try($stack, @status);
			pop(@$stack);
			($status[$i - 1], $status[$i]) = ($status[$i], $status[$i - 1]);
		}
		if ($status[$i] eq $b && $i - 2 >= 0 && $status[$i - 2] eq $s) {
			($status[$i], $status[$i - 2]) = ($status[$i - 2], $status[$i]);
			push(@$stack, \@status);
			try($stack, @status);
			pop(@$stack);
			($status[$i - 2], $status[$i]) = ($status[$i], $status[$i - 2]);
		}
	}
}

sub main {

	print @init, "\n";
	print @final, "\n";

	print "start \n";
	my $stack = +[];
	try($stack, @init);
}

&main();
