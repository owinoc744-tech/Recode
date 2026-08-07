#!/bin/bash
  count=10
until [count=0]
do 
echo 'Tminus '
count=$((count -1))
sleep 1
done
echo Blastoff
