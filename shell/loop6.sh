#!/bin/bash
for i in 1  2  3  4  5  6  7  8  9
do 
for j in 1 2 3 4 5 6 7 8 9 
do
multiply=$(( $i * $j ))
echo -n " $multiply"
done
   echo 
done
