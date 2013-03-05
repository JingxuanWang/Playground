#!/usr/bin/perl

use strict;
use Benchmark;


sub main {

	timethese (
		10_000_000,
		{
			'local' => q{local $a = $_; $a *= 2;},
			'my' 	=> q{   my $a = $_; $a *= 2;},
		}
	);
}

main();
