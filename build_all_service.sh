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

#begin path
begin_path=$PWD

service_source_root=(
  ./cmd/
)

for ((i = 0; i < ${#service_source_root[*]}; i++)); do
  cd $begin_path
  service_path=${service_source_root[$i]}
  cd $service_path
  make install
  if [ $? -ne 0 ]; then
        echo -e "${RED_PREFIX}${service_names[$i]} build failed ${COLOR_SUFFIX}\n"
        exit -1
        else
         echo -e "${GREEN_PREFIX}${service_names[$i]} successfully be built ${COLOR_SUFFIX}\n"
  fi
done
echo -e ${YELLOW_PREFIX}"all services build success"${COLOR_SUFFIX}
