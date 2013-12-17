#!/usr/bin/perl

use strict;
use Data::Dumper;
use Benchmark qw(:all);

sub returning(&) {
    my ($code) = @_; 
    return $code->();
}

sub calc {
    my ($a, $b) = @_; 
    return $a * $a + $b * $b; 
}

sub main {

	my $h = +{
		1 => 'a',
		2 => 'bc',
		3 => 'def',
	};

	print Dumper @$h->{qw/1 2 3/};
}
sub test {
    my $r = timethese(
        10_000_000, 
        {   
            '2_flat' => q{
                my $a = int(rand(100));
                my $b = int(rand(100));
                my $c = $a * $a + $b * $b;
            },
            '3_sub' => q{
                my $a = int(rand(100));
                my $b = int(rand(100));
                my $c = calc($a, $b);
            },
            '1_closure' => q{
                my $a = int(rand(100));
                my $b = int(rand(100));
                my $c = returning(sub {
                    return $a * $a + $b * $b;
                });
            },
        }   
    );  
    cmpthese $r; 
}

main();

