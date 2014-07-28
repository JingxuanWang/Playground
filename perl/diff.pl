#!/usr/bin/perl

use strict;
use warnings;
use Data::Dump;

my $a = +[1,2,3,4,5];
my $b = +[1,3,5,7,9];
my $c = +[2,4,6,8,10];
my $d = +[10];
my $e = +[1];
my $f = +[];

sub main {
	my ($old, $new) = @_;

	my $o = 0;
	my $n = 0;

	while($o < scalar(@$old) && $n < scalar(@$new)) {

		my $oo = $old->[$o];
		my $nn = $new->[$n];

		if ($oo == $nn) {
			print "$oo (=) $nn\n";
			$o++;
			$n++;
		} elsif ($oo < $nn) {
			print "$oo (<) $nn\n";
			$o++;
			print "$oo is REMOVED\n"
		} else {
			print "$oo (>) $nn\n";
			$n++;
			print "$nn is ADDED\n"
		}	
	}

	while ($o < scalar(@$old)) {
		my $oo = $old->[$o];
		print "$oo is REMOVED\n";
		$o++;
	}

	while ($n < scalar(@$new)) {
		my $nn = $new->[$n];
		print "$nn is ADDED\n";
		$n++;
	}
}

&main($b, $f);
