#!/usr/bin/perl

use strict;
use warnings;

use JSON;
use JSON::Schema;
use Data::Dump;
use Term::ANSIColor qw(:constants);

my $JSON_DATA_PATH = "./JSON/";
my $JSON_SCHEMA_DATA_PATH = "./Schema/";
my $ERROR_COUNT = 0;

sub validate {
	my ($json_file_name) = @_;

	my $json_schema_file_name;
	if ($json_file_name =~ /^(.*)\.json$/) {
		$json_schema_file_name = $1."Schema.json";
		#print "$json_schema_file_name\n";
	}

	$json_file_name = $JSON_DATA_PATH.$json_file_name;
	$json_schema_file_name = $JSON_SCHEMA_DATA_PATH.$json_schema_file_name;

	my $json_string;
	open FILE, "<$json_file_name" or die "Can not open file : $json_file_name";
	while (my $line = <FILE>) {
		$json_string .= $line;
	}
	close FILE;

	#print "$json_string\n";

	my $json_schema_string;
	open FILE, "<$json_schema_file_name" or die "Can not open file : $json_schema_file_name";
	while (my $line = <FILE>) {
		$json_schema_string .= $line;
	}
	close FILE;

	#print "$json_schema_string\n";

	my $validator = JSON::Schema->new($json_schema_string);
	my $json      = from_json($json_string);
	my $result    = $validator->validate($json);

	if ($result) {
		print GREEN, "Pass ", RESET, "$json_file_name\n";
	} else {
		print RED, "Fail ", RESET, "$json_file_name\n";
		$ERROR_COUNT++;
	}

}


sub main {

	my @files = split(' ', `ls $JSON_DATA_PATH`);

	for my $file (@files) {
		if ($file !~ /.json$/) {
			next;
		}
		validate($file);
	}

	if ($ERROR_COUNT > 0) {
		print "There are errors. Please check execution log.\n";
	} else {
		print "All Clear!\n";
	}
}

&main();
