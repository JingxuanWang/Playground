#!/usr/bin/perl

use strict;
use Data::Dumper;

sub c002 {
	my $pos = <STDIN>;
	my $num = <STDIN>;

	my @strings;
	for (my $i = 0; $i < $num; $i++) {
		$strings[$i] = <STDIN>;
	}

	@strings = map {
		$_->[0]
	} sort {
		$b->[1] <=> $a->[1]
	} map {
		+[$_, length $_]
	} @strings;

	for (my $i = 0; $i < $pos; $i++) {
		print $strings[$i];
	}
}

sub c004 {
	my $line = <STDIN>;
	my @words = split(' ', $line);
	my $dic = +{};
	for my $word (@words) {
		$dic->{$word}++;
	}
	for my $key (@words) {
		if (exists $dic->{$key}) {
			print "$key $dic->{$key}\n";
			delete $dic->{$key};
		}
	} 
}

sub judge_leap_year {
	my $line = <STDIN>;
	for (my $i = 0; $i < $line; $i++) {
		my $year = <STDIN>;
		chomp($year);
		if ($year % 400 == 0 || ($year % 4 == 0 && $year % 100 !=0)) {
			print "$year is a leap year\n";
		} else {
			print "$year is not a leap year\n";
		}
	}
}

&judge_leap_year();
