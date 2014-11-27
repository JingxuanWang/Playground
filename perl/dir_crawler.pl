#!/usr/bin/perl

use strict;
use warnings;

my $START_PAGE = +[
	"http://www.google.com"
];

my @QUEUE;

sub parse {
	my ($url) = @_;

	my $html = `curl $url`;
	my @lines = split(/[\<\>]/, $html);

	for my $line (@lines) {
		if ($line =~ /a href=\"(.*)\"/) {
			my $sub_url = $1;
			my $full_url = $url.$sub_url;
				
			print "$full_url\t";

			# folder
			if ($sub_url =~ /^[^\/]\S*\/$/) {
				print "FOLDER\n";
				push(@QUEUE, $full_url);
			}
			# asset
			elsif ($sub_url =~ /^[^\/]\S*\.\S*$/) {
				print "ASSET\n";

				# create dir
				my @dirs = split(/\//, $full_url);

				# remove http
				shift @dirs;
				# remove ''
				shift @dirs;
				# remove filename
				pop @dirs;

				my $dir = join('/', @dirs);

				unless (-e $dir) {
					`mkdir -p $dir`;
				}

				# change dir
				chdir $dir;

				`curl -O $full_url`;

				# restore dir
				my $back_dir = join('/', map { ".." } @dirs);
				chdir $back_dir;
			}
			
			print "\n";
		}
	}
}

sub main {

	push @QUEUE, @$START_PAGE;

	while (scalar(@QUEUE) > 0) {
		my $elem = shift @QUEUE;
		parse($elem);
	}
}

&main();
