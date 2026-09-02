#!/bin/bash
if [[ -f if7.sh ]] 
then
echo "file exists"
else
echo "file does not exist"
fi
if [[ -f  if4* ]]
then 
echo "file not a directory"
elif [[ -d if4* ]] then
echo "file is a directory"
fi
