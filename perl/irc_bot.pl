#!/usr/bin/perl

use strict;
use Bot::BasicBot;

# with all known options
my $bot = wBot->new(

  server => "irc.admin.mbga.local",
  port   => "6667",
  channels => ["#wjx-test"],

  nick      => "test-bot",
  alt_nicks => ["test-bot"],
  username  => "test-bot",
  name      => "test-bot",
);
eval {
	$bot->run();
};
if ($@) {
	print $@;
}

package wBot;

use base qw/Bot::BasicBot/;
use Data::Dumper;

my $PARAMS = +{};

sub said {
	my $self = shift;
	my ($params) = @_;

	my $who = $params->{who};
	my $content = $params->{body};
	return if ($who eq 'NickServ');

	if ($content eq 'time') {
		print `date "+%Y-%m-%d %H:%M:%S"`;

		#$self->schedule_tick(1);
		#$self->notice(
		#	channel => "#wjx-test",
		#	body => "$content",,
		#);
		#$self->schedule_tick(1);
		$self->forkit(+{
			channel => "#wjx-test",
			run     => q{ date "+%Y-%m-%d %H:%M:%S" },
		});
	}
	elsif ($content =~ /^(\S*)\+\+$/) {
		my $var = $1;
		$PARAMS->{$var}++;
		$self->notice(
			channel => "#wjx-test",
			body => "$var : $PARAMS->{$var}",
		);
	}

	# debug
	$self->notice(
		channel => "#wjx-test",
		body => "Noticed $who said : $content",
	);

}

1;

