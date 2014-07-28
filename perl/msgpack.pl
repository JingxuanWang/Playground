#!/usr/bin/perl

use strict;
use warnings;

use lib "~/perl5";

#use JSON;
use Data::MessagePack;
use Data::Dumper;
use Getopt::Long;

my $OUT_PATH = "./";

my $hash = +{
	a => +{
		f => -0.123,
		g => 1.2345,
		l => 1234567890
	},
	bList => [
		{ name => "Hello" }, 
		{ name => "World" }
	],
	count => 100,
	data => "hoge",
	stringList => ["This", "Is", "A", "Test"],
};
	
my $mp = Data::MessagePack->new();

my $SERIALIZE = 0;
my $DESERIALIZE = 0;
my $FILE;

GetOptions(
	's' => \$SERIALIZE,
	'd' => \$DESERIALIZE,
	'file=s' => \$FILE,
);

sub serialize {
	print $mp->pack($hash);
}

sub deserialize {
	my ($bytes) = @_;
	print Dumper $mp->unpack($bytes);
}

sub main {
	
	if ($SERIALIZE) {
		serialize();
		return;
	}
	if ($DESERIALIZE) {
		my $bytes;
		if ($FILE) {
			open (FILE, "<$FILE") or die "Can not open file $FILE";
			while(my $line = <FILE>) {
				chomp($line);
				$bytes .= $line;
			}
		}
		deserialize($bytes);
		return;
	}
	
}

&main();
