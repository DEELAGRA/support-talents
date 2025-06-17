#!/bin/bash

echo "Как тебя зовут?"
read name

if [ -d "$name" ]; then
    echo "Папка с именем $name уже существует"
else 
    echo "Папки с именем $name успешно создана"
    mkdir $name
    cd $name
    echo "Привет, $name. Ты создал папку с доком" > welocme.txt
fi