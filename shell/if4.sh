#!/bin/bash
name=liam
 echo $name
if [[ $name == liam ]]
then 
echo "name is truly liam"
fi
echo " next type admin as user and 1234 as password"
u=admin
p=1234
read -p "enter user : " us
if [[ $u == $us ]]
then 
echo "welcome $u"
elif [[ $u != $us ]]
then 
 echo "incorrect username"
fi
read -p "enter password" p
if [[ p = 1234 && 4321 ]]
then 
echo "correct password"
elif  [[ p != 1234 && 4321 ]]
then 
echo "incorrect password"
fi 
