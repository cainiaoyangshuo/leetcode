#!/bin/bash
# get all filename in specified path
path=$1
echo $path
files=$(ls $path)

for filename in $files
do
echo $filename
#if ["$filename" == *.php];then
#len=${filename}
#dirname=${filename: 0: $len}
#echo $dirname
#fi
done

