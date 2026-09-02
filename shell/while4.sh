#!/bin/bash
math=1
while [ $math -gt 0 ] && [ $math -le 10 ]
  do
  echo "count $math"
math=$((math+1))
if [ $math == 7 ]
then 
break
fi
done
