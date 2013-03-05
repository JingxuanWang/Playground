#!/usr/bin/perl

use strict;
use POSIX qw /pow/;
use Inline qw /C/;

use wjxUtil;
use Data::Dumper;
use IO::Compress::Gzip qw/gzip $GzipError/;

my $List = [
	{
		shop_id 	=> 1,
		item_id 	=> 101,
		item_type 	=> 1,
		item_price 	=> 100,
	},
	{
		shop_id 	=> 2,
		item_id 	=> 102,
		item_type 	=> 2,
		item_price 	=> 200,
	},
	{
		shop_id 	=> 1,
		item_id 	=> 103,
		item_type 	=> 3,
		item_price 	=> 300,
	},
];

my $rr = +{
	1 => 0.25,
	3 => 0.2,
	7 => 0.15,
	14 => 0.10,
	30 => 0.05,
	60 => 0.01,
	90 => 0.01,
};

sub calc_future_dau {
	my ($daily_reg, $days, $rr_diff) = @_;
	
	my $dau = 0;
	my $return_rate = 1;

	for (my $i = 0; $i < $days; ++$i) {
		if (exists $rr->{$i}) {
			$return_rate = $rr->{$i} + $rr_diff;
		}
		my $plus = ($daily_reg * $return_rate);
		#print "Day $i: $plus users\n";
		$dau += int($daily_reg * $return_rate);
	}
	return $dau;
}

sub main {

	my $price;
	my $from = 26;
	my $to = 270;
	my $factor = 1.7;
	$price = int($from * POSIX::pow($to, $factor) / 47.5);

	#print "$from : $to : $factor : $price \n";
	my $num = 135280;
	my $max = 100000000;
	my $rate = 1;
	for (my $i = 0; $i < 1000; $i++) {
		$rate *= 1 - $num/$max;
		++$num;
	}
	$rate = 1 - $rate;
	print "$rate\n";

	my $user;
	$user->testFunc();
}

sub add {
	my ($a, $b) = @_;
	return $a + $b;
}

sub protect {
	my ($chara) = @_;
	test($chara, "dead");
}
sub test {
	my ($chara, $status) = @_;
	$chara->{status} = $status;
	if ($status eq "dead") {
		--$chara->{chara_remain};
	}
}

#&main();

sub abc {
	my $prime_count = 0;
	my $ceiling = 100000;

	for (1 .. $ceiling) {
		$prime_count++ if (check_prime($_, int(sqrt$_) + 1));
	}

	print "total $prime_count\n";
}

sub is_prime {
	my $number = shift;
	my $divisor = int (sqrt $number) + 1;
	while( $divisor > 1) {
		return 0 if ($number % $divisor-- ==0); 
	}
	return 1;
}

&abc();

__END__
__C__

int check_prime (int number, int divisor) {
	while(divisor > 1) {
		if (number % divisor-- == 0) {
			return 0;
		}
	}
	return 1;
}
