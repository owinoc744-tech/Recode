#!/bin/bash
#  count=5
until [[ count == 5 ]]
do 
count=$((count +1))
echo 'Tminus '
sleep 1
done
echo Blastoff
