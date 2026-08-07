#!/bin/sh
echo '====FILE READER===='
echo "Enter your notes"
read line
echo $line
for i in $(seq 145)
do 
echo $line >> file2.txt
done 
