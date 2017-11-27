#!/bin/bash

# This script make a lot of assumptions and has no error handling


BIN_DIR=`dirname "$0"`
cd $BIN_DIR/../..

BASE_DIR=`pwd`

echo "Base Dir: " $BASE_DIR

rm -rf $BASE_DIR/build
mkdir $BASE_DIR/build

###################
# Build and Package for OSX
###################
echo "Building for OSX"

mkdir $BASE_DIR/build/osx
export GOOS="darwin"

rm -f Issues2Markdown
go build
mv Issues2Markdown $BASE_DIR/build/osx


###################
# Build and Package CLI for Linux
###################


echo "Building for Linux"

mkdir $BASE_DIR/build/linux
export GOOS="linux"

rm -f Issues2Markdown
go build
mv Issues2Markdown $BASE_DIR/build/linux

##################
# Build and Package CLI for Windows
###################


echo "Building for Windows"

mkdir $BASE_DIR/build/windows
export GOOS="windows"

rm -f Issues2Markdown.exe
go build
mv Issues2Markdown.exe $BASE_DIR/build/windows

###################
# Done!!!
###################
export GOOS=""
echo "Done..."
echo ""
echo "See $BASE_DIR/build for binary files"
open $BASE_DIR/build



