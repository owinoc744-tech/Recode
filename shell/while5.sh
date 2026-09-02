#!/bin/bash
nu=1
while   [ $nu -gt 0 ]  && [ $nu -lt 5 ]
do 
read  -p "enter number 1..5: " nu
if [ $nu == 1 ] ||[ $nu == 2 ] || [ $nu == 3 ]
then 
echo "good choice"
 elif [[ $nu == 4 || 5 ]]
 then
echo "bad choice" 
fi
done
