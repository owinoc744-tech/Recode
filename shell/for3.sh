#!/bin/bash
# prints the alphabet in order
echo -e "alphabet in ascending order\n"
for a in {a..z}
do 
echo  -n  " $a"
done
echo
echo   -e  "\nalpbet in descending order"
echo
for b in {z..a}
do 
echo  -n " $b "
done
echo -e "\n"
