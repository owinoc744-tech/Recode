#!/bin/bash

read -p "enter range : " r
if [ $r -ge 30 ] # || [ $r ]
then 
echo "out of range"
elif [ $r -ge 16 ] || [ $r lt 30 ]
then 
echo "in range "
elif [ $r -lt 15 ]
then 
echo "range low"
fi

