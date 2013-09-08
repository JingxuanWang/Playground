#!/usr/bin/perl

use strict;
use Benchmark;


sub main {

	my $max_index = 10000;
	my @arr = 1..$max_index;

	timethese (
		1_000_000,
		{
			'1_indep' => q{
					my $rand = int(rand());
					for (my $i = 0; $i < $max_index; $i++) {
						$arr[$i] = $rand;
					}
				},
			'2_dep' => q{
					my $rand = int(rand());
					$arr[0] = $rank;
					for (my $i = 1; $i < $max_index; $i++) {
						$arr[$i] = $arr[$i - 1];
					}
				},
		}
	);
}

main();
