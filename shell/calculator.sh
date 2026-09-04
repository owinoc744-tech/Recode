#!/bin/bash
while true ; do
echo "========≠=============="
echo "#==shell calculator===#"
echo "========≠=============="
  sleep 1
read -p "Enter first number : " n1
read -p "Enter operator : " op
read -p "Enter second number : " n2
  case $op in 
"+") echo addition $(( $n1 + $n2 )) ;;

 "-") echo substraction $(( $n1 - $n2 )) ;;
 "*") echo multiply $((n1*n2)) ;;
  "/") echo division $((n1/n2)) ;;
   "%") echo modulus $((n1%n2)) ;;
esac
done
