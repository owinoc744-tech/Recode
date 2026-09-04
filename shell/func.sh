#!/bin/bash
function Add(){
local n=5
 local b=6
local r=$(( n + b ))
echo  "add : $r"
}
 function minus(){
local b=50
local n=15
local c=$(( b - n))
echo "minus : $c"
}
function multiply(){

local a=34
local b=23
local c=$(( a * b))
echo "multiplication : $c"
}
function devide(){
local a=35
local b=7
local c=$((a/b))
echo "division : $c"
}
function modulo(){
 local a=7
 local b=5
local c=$((a%b))
echo "modulus : $c"
}
Add
minus
multiply
devide
modulo

