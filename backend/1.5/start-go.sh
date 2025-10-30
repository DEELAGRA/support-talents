#!/bin/bash

project_name="$1"

if [ -z "$project_name" ]; then
    echo "Использование: $0 имя_проекта"
    exit 1
fi

mkdir "$project_name"
cd "$project_name" || exit 1

go mod init github.com/DEELAGRA/"$project_name"

cat <<EOF > main.go
package main

import "fmt"

func main(){
    fmt.Println("Привет из $project_name!")
}
EOF


echo "✔ Проект '$project_name' создан"