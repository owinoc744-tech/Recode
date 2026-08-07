#!/bin/sh
count=10
while [ $count -gt 0 ]
do
    echo "Starting in $count..."
    sleep  0.5
    count=$((count - 1))
done
echo "Go!"
