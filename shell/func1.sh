#!/bin/bash
add(){
local a=12
local b=13
local c=$((a+b))
echo "addition:12+13: $c"
}
minus(){
local a=23
local b=8
local c=$((a-b))
echo "minus:23-8: $c"
}
multi(){
a=24
b=12
c=$(($a*$b))
echo "multiplication:24×12 $c"
}
devide(){
a=24
b=6
c=$((a/b))
echo "division:24/6: $c"
}
modulo(){
a=25
b=7
c=$((a%b))
echo "modulus:25%7: $c"
}
add
minus
multi
devide
modulo
