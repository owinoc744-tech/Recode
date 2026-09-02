#!/bin/bash
r=16
read -p "Enter range : " ra
if [ $ra -gt $r ]
then 
echo "out of range"
elif [ $ra -lt $r ]
then 
echo "range low"

fi
