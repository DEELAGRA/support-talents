#!/bin/bash
count_txt=0
count_sh=0
count_dir=0


if [ -z "$1" ]; then
    echo "Введите путь к директории при запуске скрипта"
    exit 1
fi
cd $1
for file in *.txt; do
    
    if [ -f "$file" ]; then
        echo "В файле $file"
        head -n 3 $file
        echo ""
        ((count_txt++))
        
    fi
done

for file in $1; do
    if [ -x $file ]; then
        ((count_sh++))
    fi
done

for item in */; do
    if [ -d "$item" ];then
        if [ -z "$(ls -A $item)" ]; then
            ((count_dir++))
        fi
    fi
done

echo "Найдено: $count_txt txt, $count_sh exe, $count_dir пустых папки"