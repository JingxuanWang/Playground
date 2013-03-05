#!/usr/bin/perl

my $from = 'shift-jis';
my $to = 'utf-8';

sub main {
	my @FileList = @ARGV;

	for my $file (@FileList) {
		my $tFile = $file.".tmp";
		my $bFile = $file.".bak";
		#`iconv -f $from -t $to $file > $tFile`;
		`nkf -s -W -O $file $tFile`;
		`cp $file $bFile`;
		`mv $tFile $file`;
	}
}

&main();
