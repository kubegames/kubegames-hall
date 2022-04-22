#!/bin/bash

pwd=`pwd`
ls=`ls`
model=$pwd"/app/model"
service=$pwd"/app/service"
length=${#pwd}

buildProto(){
  for file in `ls $1`
    do
      if [ -d $1"/"$file ]
      then
          if [[ $file != '.' && $file != '..' ]]
          then
              buildProto $1"/"$file
          fi
      else
        if [ "${file##*.}"x = "proto"x ]
        then
          path=${1: $length+1}
          name=${file%.*}
          gofile="./"$path/$name"*.go"
          jsonfile="./"$path/$name"*.json"
          protofile="./"$path/$name".proto"
          set -ex
          rm -rf $gofile $jsonfile
          set +ex
        fi     
      fi
  done
}

buildSwagger(){
  for file in `ls $1`
    do
      if [ -d $1"/"$file ]
      then
          if [[ $file != '.' && $file != '..' ]]
          then
              buildSwagger $1"/"$file
          fi
      else
        if [[ $file == *"gin.pb.go" ]]
        then
          path="./"${1: $length+1}
          docs=$path"/docs"
          set -ex
          rm -rf $docs
          set +ex
        fi     
      fi
  done
}

# rm
rm -rf ./app/service/docs

buildProto $model

buildProto $service

buildSwagger $service


