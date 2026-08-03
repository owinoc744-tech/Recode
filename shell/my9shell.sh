#!/bin/bash
 a=3 #this is how to declare a variable in shell
  b=6
   c=3
    d=20
    sum=$((a+b)) #this is the new  way of doing calculations
    minus=$((a-b))
     multiply=$((c*d))
      devide=$((d/c))
       modulus=$((d%c))
#to echo is to print on the console
       echo 'sum=' $sum
      echo 'minus='$minus
     echo 'multiply='$multiply
    echo 'devide='$devide
   echo 'modulo='$modulus
# how to `expr`
 x=50
  y=20
   add=`expr $x + $y`
   minus=`expr $x - $y`
    multiply=`expr $x \* $y`
     division= `expr $y \/ $x`
   echo 'add:'$add
  echo 'minus:'$minus
 echo 'multiply:'$multiply
echo  'division:'$division
