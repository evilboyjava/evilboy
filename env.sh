#!/bin/bash
if [ "darwin" = "$1" ] ;then
    export GOPATH=$GOPATH:`pwd`
    export GOOS=darwin
elif [ "linux" = "$1" ] ;then  
	export CGO_ENABLED=0 
	export GOOS=linux
	export GOARCH=amd64 
    export GOPATH=$GOPATH:`pwd`
else 
    export GOPATH=$GOPATH:`pwd`
	export GOOS=darwin
fi  
echo $GOPATH

