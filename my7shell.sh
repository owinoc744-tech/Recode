#!/bin/bash
echo Enter your name
read name 
 echo Welcome $name  >> user.txt
echo  "Enter your id" 
read id
echo your id number is $id >> user.txt
echo enter your password
read -s   password
echo password recieved 
echo  $name password is $password >> user.txt
