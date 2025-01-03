#!/bin/bash
SERVICE_NAME=$1
RELEASE_VERSION=$2
 
sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

if [ ! -d "golang" ]; then
    mkdir golang
fi
cd golang

if [ ! -d "${SERVICE_NAME}" ]; then
    mkdir ${SERVICE_NAME}
fi
cd ${SERVICE_NAME}
cd ../../

protoc --go_out=./golang --go_opt=paths=source_relative \
  --go-grpc_out=./golang --go-grpc_opt=paths=source_relative \
 ./${SERVICE_NAME}/*.proto

cd golang/${SERVICE_NAME}
go mod init \
  github.com/andrei-kozel/gladiator-arena-proto/golang/${SERVICE_NAME} ||true
go mod tidy
cd ../../ 
git config --global user.email "kozel.andrei.94@gmail.com"
git config --global user.name "Andrei Kozel"
git add . && git commit -am "proto update" || true
git tag -fa golang/${SERVICE_NAME}/${RELEASE_VERSION} \
  -m "golang/${SERVICE_NAME}/${RELEASE_VERSION}" 
git push origin refs/tags/golang/${SERVICE_NAME}/${RELEASE_VERSION}
