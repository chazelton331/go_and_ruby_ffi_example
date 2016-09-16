#!/bin/sh

PROGRAM_NAME=bon.go

go build -buildmode=c-shared -o libbongo.so $PROGRAM_NAME

go build $PROGRAM_NAME
