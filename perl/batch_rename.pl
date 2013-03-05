#!/usr/bin/perl

use strict;
use warnings;
use Getopt::Long;
use Pod::Usage;
use Data::Dumper;

=head1 NAME

batch_rename.pl - Rename multiple files in a single run!

=head1 SYNOPSIS

batch_rename.pl OPTIONS [DIR] [SRC] [DST]

=head1 OPTIONS

=over

=item C<-r>

Run name subtitution recursively

=item C<-n>

No dry run mode

=item C<-v>

Verbose mode

=back

=cut

my %OPT;

my $DIR;
my $SRC;
my $DST;

sub runCmd {
	my ($cmd) = @_;
	print "Executing  $cmd ... \n";
	if (!$OPT{'no-dryrun'}) {
		`$cmd`;
	}
}

sub name_convertion {
	my ($file_name) = @_;

	my $new_file_name = $file_name;
	# here to implement name convertion rules
	#if (1) {
		$new_file_name =~ s/$SRC/$DST/g;
	#}

	return $new_file_name; 
}

sub getFullPathName {
	my ($file) = @_;
	my $cur_dir = `pwd`;
	chomp($cur_dir);
	return "$cur_dir/$file";
}

sub main {
	if (!GetOptions(\%OPT, qw(no-dryrun|n recursive|r verbose|v))) {
		pod2usage(0);
	}

	#print Dumper @ARGV;
	my $cur = `pwd`;
	chomp($cur);

	$DIR = shift(@ARGV);
	$SRC = shift(@ARGV);
	$DST = shift(@ARGV);

	my $dir_flag = +{};
	my @dirs;
	
	# if $DIR starts with '.' or '..'
	# we should call getFullPathName
	# otherwise we should use $DIR as global directory
	if ($DIR =~ /^\./) {
		$DIR = getFullPathName($DIR);
	}
	
	push(@dirs, $DIR);
	$dir_flag->{$DIR} = 1;

	# BFS
	while(@dirs) {
		my $dir = shift(@dirs);
		
		chdir "$dir" or die "Can not enter dir $dir, $!";
		
		if ($OPT{'verbose'}) {
			print STDERR "Searching $dir ... \n";
		}

		my @all_files = glob "*";

		for my $file (@all_files) {
			if (-d $file && $OPT{'recursive'}) {
				# get its global directory
				my $tgt = getFullPathName($file);
				
				# check visit flag
				if (!$dir_flag->{tgt}) {
					push(@dirs, $tgt);
					# set visit flag
					$dir_flag->{$tgt} = 1;
				}
			} elsif (-e $file && $file =~ /$SRC/) {
				my $new_name = name_convertion($file);
				runCmd("mv $file $new_name");
			} else {
				# do nothing
			}
		}
	}
}

&main();
