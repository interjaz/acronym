#!/bin/bash

#make sure all require directories exits
mkdir -p "bin"
mkdir -p "pkg"
mkdir -p "src"

#set new Go path
newGoPath=`echo $(pwd)`

#get required packages for developing in go in VS Code
export GOPATH=$newGoPath
export PATH=$newGoPath/bin:$PATH

#create export script to be run in the terminal
echo "# ----------------"
echo "# Run following script or set it in your .bash_profile"
echo "# ----------------"
echo "# Acronym.Server.Fetcher settings"
echo "export GOPATH=\"$newGoPath/\""
echo "export PATH=\"$newGoPath/bin:$PATH\""
