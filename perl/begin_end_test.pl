#- BeginEndTest.pl
#- Copyright (c) 1995 by Dr. Herong Yang, http://www.herongyang.com/
#
$f = "Fortran";
print("Printing before the first BEGIN() code...\n");

sub BEGIN {
  print("In BEGIN: name space = ",__PACKAGE__,"\n");
  print("In BEGIN: \$f = $f...\n");
}   	

sub CHECK {
  print("In CHECK: name space = ",__PACKAGE__,"\n");
  print("In CHECK: \$f = $f...\n");
}   	

sub INIT {
  print("In INIT: name space = ",__PACKAGE__,"\n");
  print("In INIT: \$f = $f...\n");
}   	

sub END {
  print("In END: name space = ",__PACKAGE__,"\n");
  print("In END: \$f = $f...\n");
}   	

print("Printing after the first END() code...\n");

package MyPackage;
print("Printing before the second BEGIN() code...\n");

sub BEGIN {
  print("In BEGIN: name space = ",__PACKAGE__,"\n");
  print("In BEGIN: \$f = $f...\n");
}   	

sub CHECK {
  print("In CHECK: name space = ",__PACKAGE__,"\n");
  print("In CHECK: \$f = $f...\n");
}   	

sub INIT {
  print("In INIT: name space = ",__PACKAGE__,"\n");
  print("In INIT: \$f = $f...\n");
}   	

sub END {
  print("In END: name space = ",__PACKAGE__,"\n");
  print("In END: \$f = $f...\n");
}   	

print("Printing after the second END() code...\n");
exit;
