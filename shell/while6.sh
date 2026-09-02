#!/bin/bash
x=0
while [ $x == 30 ] 
do 
echo "$conut $x"
x=$((x+1))
if [ $x -gt 6 ]
then 
continue
elif [ $x == 20]
then
break
fi
done
