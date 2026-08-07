#!/bin/sh
echo "====CALCULATOR===="
echo enter first number
 read n1
 echo enter operator
op=+,-,*,/
  read op
echo eneter second number
  read n2
# result=$(( $n1 + $n2 ))
for result in ( 1)
do
case +
$((n1+n2))
case -
$((n1-n2))
case /
$((n1/n2))
case *
$((n1/*n2))
done
echo   results $result
#while 
#math=2
#do 
#echo 

#done
