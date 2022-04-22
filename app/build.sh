#!/bin/bash

pwd=`pwd`
ls=`ls`
model=$pwd"/app/model"
service=$pwd"/app/service"
docs=$pwd"/app/service/docs"
length=${#pwd}

buildModel(){
  for file in `ls $1`
    do
      if [ -d $1"/"$file ]
      then
          if [[ $file != '.' && $file != '..' ]]
          then
              buildModel $1"/"$file
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
          protoc --proto_path=./ --gofast_out=plugins=grpc:./ --gofast_opt=paths=source_relative $protofile
          protoc --proto_path=./ --xorm_out=./ --xorm_opt=paths=source_relative $protofile
          #protoc --proto_path=./ --xorm_out=./ --xorm_out=paths=source_relative $protofile
          set +ex
        fi     
      fi
  done
}

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
          protoc --proto_path=./ --gofast_out=plugins=grpc:./ --gofast_opt=paths=source_relative $protofile
          protoc --proto_path=./ --gin_out=./ --gin_opt=paths=source_relative $protofile
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
          mkdir -p $docs
          protoc --proto_path=./ --openapiv2_out=allow_delete_body=true,allow_merge=true,logtostderr=true:$docs $path/*.proto
          createSwagger $docs
          set +ex
        fi     
      fi
  done
}

buildAllSwagger(){
  for file in `ls $1`
    do
      if [ -d $1"/"$file ]
      then
          if [[ $file != '.' && $file != '..' ]]
          then
              buildAllSwagger $1"/"$file
          fi
      else
        if [[ $file == *"gin.pb.go" ]]
        then
          path="./"${1: $length+1}
          protos=$protos" "$path"/*.proto"
        fi     
      fi
  done

}

createSwagger(){
  cat > $1/swagger.go << EOF
package docs

import (
	_ "embed"
)

var (
	//go:embed *swagger.json
	swagger string
)

//get docs
func GetDocs() string {
	return swagger
}
EOF
}

# build
buildModel $model

buildProto $service

buildSwagger $service

buildAllSwagger $service

set -ex
mkdir -p ./app/service/docs
createSwagger ./app/service/docs

protoc --proto_path=. \
--proto_path=./ \
--openapiv2_out ./app/service/docs \
--openapiv2_opt allow_delete_body=true \
--openapiv2_opt allow_merge=true \
--openapiv2_opt logtostderr=true \
--openapiv2_opt json_names_for_fields=true \
$protos
set +ex