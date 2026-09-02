#!/bin/bash
  c=5
if [ $c -eq 5 ] ;
 then
echo "okay"
fi
age=18
read -p "enter your age " age
if [ $age -ge 18 ] 
then 
echo "you can vote " 
else 
echo "you are a minor"
 fi 

