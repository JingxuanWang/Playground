package utils;

use strict;

sub setUnionIntersectDiff {
	my ($setA, $setB) = @_;
	
	my %union;
	my %isect;
	foreach (@$setA, @$setB) {
		$union{$_}++ && $isect{$_}++;
	}
	
	my @union	= keys %union;  
	my @isect	= keys %isect;  
	my @diff	= grep { $union{$_} == 1; } @union; 

	return (\@union, \@isect,\@diff);
}

sub hashIntersect {
	my ($rBase, $rOther) = @_;

	my $ret = {};

	if (!$rBase || !$rOther) {
		return $ret;
	}

	for my $key (keys %$rBase) {
		if (defined $rOther->{$key}) {
			$ret->{$key} = $rBase->{$key};
		}
	}
	
	return $ret;
}

sub mergeHash {
    my ($rDstHash, $rSrcHash, $match) = @_; 
    
    if ($match) {
        for my $key (keys(%{$rSrcHash})) {
            if ($key =~ /$match/) {
                $rDstHash->{$key} = $rSrcHash->{$key};
            }
        }
    } else {
        for my $key (keys(%{$rSrcHash})) {
            $rDstHash->{$key} = $rSrcHash->{$key};
        }
    }   
}

# get date format (yyyymmdd) from datetime integer gotten by time() method
sub get_date {
    my ($base_time) = @_; 
    my ($sec, $min, $hour, $day, $month, $year) = localtime($base_time);
    $year += 1900;
    $month++;
    return ($year, $month, $day, $hour, $min, $sec);
}

# return start time (base_date 00:00:00) in select condition
sub convert_start_time {
    my ($base_time) = @_;   
    my ($sec, $min, $hour, $day, $month, $year) = localtime($base_time);
    return timelocal(0, 0, 0, $day, $month, $year); 
}

# return end time (base_date 23:59:59) in select condition
sub convert_end_time {
    my ($base_time) = @_; 
    my ($sec, $min, $hour, $day, $month, $year) = localtime($base_time);
    return timelocal(59, 59, 23, $day, $month, $year);  
}

# get table text from array reference, hash reference
# the array is two-dimentional, including [rows][columns]
sub get_table_text {
    my ($records, $col_size, $row_limit) = @_; 
    my $text;
    my $row_count = 0;
    if (ref($records) eq "ARRAY") {
        foreach my $value (@$records) {
            for (my $i = 0; $i < $col_size; $i++) {
                if (!$value->[$i]) {
                    $value->[$i] = 0;
                }   
            }   
            $text .= join("\t|", @$value)."\n";
            last if ($row_limit && $row_count >= $row_limit);
            $row_count++;
        }   
    } elsif (ref($records) eq "HASH") {
        foreach my $key (sort {$a <=> $b} keys(%$records)) {
            my $value = $records->{$key};
            if (ref($value) eq "ARRAY") {
                for (my $i = 0; $i < $col_size - 1; $i++) {
                    if (!$value->[$i]) {
                        $value->[$i] = 0;
                    }   
                }   
                $text .= join("\t|", $key, @$value)."\n";
            } else {
                $text .= join("\t|", $key, $value)."\n";
            }   
            last if ($row_limit && $row_count >= $row_limit);
            $row_count++;
        }   
    }   
    $text =~ s/\n*$//;
    return $text;
}
# get seconds from time format (hh:mm:ss)
sub get_seconds {
    my $time = $_[0];
    my ($hour, $minute, $second) = split(/:/, $time);   
    return $hour * 3600 + $minute * 60 + $second;
}
# get time format(hh:mm:ss) from seconds
sub get_time_format {
    my $sec = $_[0];
    my $hour = $sec / 3600;
    my $minute = ($sec % 3600) / 60; 
    my $seconds = (($sec % 3600) % 60);
    return sprintf("%02d:%02d:%02d", $hour, $minute, $seconds);
}
# calculate percentage rate
sub calc_rate {
    my ($molecule, $denominator) = @_; 
    if ($molecule == 0 || $denominator == 0) {
        return 0;
    }   
    return $molecule / $denominator
}

sub getDays {
    my $startDate = shift;
    my $endDate   = shift;
    my $daysDiff = Delta_Days( split( "/", $startDate ), split( "/", $endDate ) );
    my @result;
    for( 0 .. $daysDiff ){
        push( @result, sprintf( "%04d%02d%02d", Add_Delta_Days( split( "/", $startDate ), $_ ) ) );
    }   
    return @result;
}

sub unclassified {
	use POSIX qw(strftime); 
	my $time = strftime("%Y/%m/%d %H:%M:%S", localtime);

	use Term::ANSIColor qw (color :constants);
	print color 'bold red';
	print "$time\n";
	print color 'reset';
	print "$time\n";

	use Getopt::Long;
    my $data   = "file.dat";
	my $length = 24;
	my $verbose;
	my $result = GetOptions (
		"length|l=i" => \$length,    # numeric
		"file|f=s"   => \$data,      # string
		"verbose|v"  => \$verbose    # flag
	);
	print "\$length=$length - \$data=$data - \$verbose=$verbose\n";
}

1;
