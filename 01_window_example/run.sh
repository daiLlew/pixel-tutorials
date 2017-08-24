#!/usr/bin/env bash

echo "Removing previous binary..."
rm -f ../lib/windows_example

echo "Compiling..."
go build -o ../lib/windows_example

echo "Running"
cd ../
./lib/windows_example