#!/usr/bin/env bash

echo "Removing previous binary..."
rm -f ../lib/sprite_example

echo "Compiling..."
go build -o ../lib/sprite_example

echo "Running"
cd ../
./lib/sprite_example -imgPath="./img/hiking.png"