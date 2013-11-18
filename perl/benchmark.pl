#!/usr/bin/perl

use strict;
use Benchmark;
use List::Util qw/shuffle/;
use POSIX qw(pow);

my $MAX = 1000000;

my @array = shuffle (1...$MAX);

sub main {

	timethese (
		1_000_000,
		{
			'1_hash'     => q{
				my @sorted = map {
					$_->{orig}
				} sort {
					$b->{square} <=> $a->{square}
					||
					$b->{square_root} <=> $a->{square_root}
				} map {
					my $square = $_ * $_;
					my $sqrt = POSIX::pow($_, 0.5);
					+{
						orig => $_,
						square => $square,
						square_root => $sqrt,
					}
				} @array;
			},
			'2_array'    => q{
				my @sorted = map {
					$_->[0]
				} sort {
					$b->[1] <=> $a->[1]
					||
					$b->[2] <=> $a->[2]
				} map {
					my $square = $_ * $_;
					my $sqrt = POSIX::pow($_, 0.5);
					+[ $_, $square, $sqrt]
				} @array;
			},
		}
	);
}

main();
