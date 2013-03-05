#!/usr/bin/perl

use strict;
use Time::HiRes;
use Data::Dumper;

$Data::Dumper::Sortkeys = 1;
my $MAX = 1000;

sub init {
	my $rResult = +{
		100 => 0,
		200 => 0,
		300 => 0,
		400 => 0,
		500 => 0,
		600 => 0,
		700 => 0,
		800 => 0,
		900 => 0,
		1000 => 0,
	};
	return $rResult;
}

sub setPos {
	my ($value, $rResult) = @_;

	for my $key (sort {$a <=> $b} keys %{$rResult}) {
		if ($value < $key) {
			++$rResult->{$key};
			return;
		}
	}
}

sub test1 {
	my ($round) = @_;
	
	my $rResult = init();
	
	#srand(time());

	for (my $i = 0; $i < $round; ++$i) {
		my $value = int(rand($MAX));
		print "$value\n";
		setPos($value, $rResult);
	}

	print Dumper $rResult;
}

sub test2 {
	my ($round) = @_;
	
	my $rResult = init();
	
	#srand(time());

	for (my $i = 0; $i < $round; ++$i) {
		my $time = int(Time::HiRes::time * 10000000);
		srand($time);
		#srand(time ^ $$);
		my $value = int(rand($MAX));
		print "$value\t$time\n";
		setPos($value, $rResult);
	}

	print Dumper $rResult;

}

sub main {

	test1(10);
	print "--------------result of test 1----------------\n";

	test2(10);
	print "--------------result of test 2----------------\n";
}


&main();
