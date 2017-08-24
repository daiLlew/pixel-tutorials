#!/usr/bin/env bash

echo "Removing previous binary..."
rm -f ../lib/matrices_example

echo "Compiling..."
go build -o ../lib/matrices_example

echo "Running"
cd ../
./lib/matrices_example -imgPath="./img/hiking.png"