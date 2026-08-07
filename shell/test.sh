#!/bin/bash
echo  '======NOTE BOOK======='
echo   'write your notes'
read note
echo $note >> note.txt
content=$(cat note.txt)
echo $content
