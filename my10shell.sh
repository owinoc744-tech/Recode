#!/bin/bash
# another example of expr
   x=400
   y=5
    add=$(expr $y + $x )
     minus=$(expr $y - $x)
  echo $add
echo $minus
