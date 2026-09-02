#!/bin/bash
n="linux"
b="linux"
echo "enter linux"
read -p "enter name above : " n 
 
if [ "$n" = "linux" ]
then 
echo "correct"
elif [ "$n" != "linux" ]
then
echo "incorrect"
fi
