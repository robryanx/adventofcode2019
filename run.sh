#!/bin/bash

go build -o ~/adventofcode/2019/bin/$1 ~/adventofcode/2019/days/$1/*

if [ $? -eq 0 ]
then
    ~/adventofcode/2019/bin/$1;
fi
