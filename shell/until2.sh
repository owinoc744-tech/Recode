#!/bin/bash
n=1
until [[ $n -gt 5 ]]
do
n=$((n+1))
echo " counting"
sleep 1
done
