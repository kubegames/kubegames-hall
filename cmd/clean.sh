#!/bin/bash
NAME=$1
PUSH=$2
DEPLOY=$3

for file in ./*
do
    if test -d $file
    then
      name=${file##*./}
      if [ $NAME == "ALL" ] || [ $name == $NAME ];then
        cd $file
          kubectl delete -f ./deploy.yaml
        cd ../
      fi 
    fi
done
