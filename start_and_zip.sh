#!/usr/bin/env bash


COLOR_SUFFIX="\033[0m"
RED_PREFIX="\033[31m"
GREEN_PREFIX="\033[32m"
YELLOW_PREFIX="\033[33m"

bin_dir="./bin"
#Automatically created when there is no bin, logs folder
if [ ! -d $bin_dir ]; then
  mkdir -p $bin_dir
fi

if [ -z "$1" ]; then
    echo "args is null you should ./start_and_zip your project name(the type must be runapi)"
            exit -1
fi

  if [ $? -ne 0 ]; then
        exit -1
        else
    cd   $bin_dir
    if [ "$2" = "x" ]; then
        # If the second argument is "x", pass -x to runapi_generator
        ./runapi_generator -p $1 -x
    else
        ./runapi_generator -p $1
    fi
   cd ..
   zip  $1.zip $1/* -r
  fi
