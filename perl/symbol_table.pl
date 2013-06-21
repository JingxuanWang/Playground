#!/usr/bin/perl

use Data::Dumper;

for my $symname (sort keys %main::) {
	local *sym = $main::{$symname};
	#print "*sym: ",*sym,"\tsymname: $symname","\n";
	print $sym if defined $sym;
	print Dumper \@sym if @sym;
	print Dumper \%sym if %sym;
}

sub populate {
	my %hash = (
		a => 1,
		b => 2,
	);
	return \%hash;
}


*units = populate();
print Dumper \%units;
