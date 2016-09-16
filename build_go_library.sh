#!/bin/sh

PROGRAM_NAME=sum.go

go build -buildmode=c-shared -o libsum.so $PROGRAM_NAME

go build $PROGRAM_NAME
