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

  if [ $? -ne 0 ]; then
        exit -1
        else
    cd   $bin_dir
   ./runapi_generator
   cd ..
   zip  OpenIMServer.zip OpenIM服务器API/* -r
  fi
