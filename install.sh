#!/usr/bin/bash

# gobot
go get -d -u gobot.io/x/gobot/...
if [ $? -ne "0" ]
then
    echo "Could not get Gobot"
fi

# gocv and opencv
go get -u -d gocv.io/x/gocv
if [ $? -ne "0" ]
then
    echo "Could not get GoCV"
fi

cd $GOPATH/src/gocv.io/x/gocv
if [ $? -ne "0" ]
then
    echo "GOPATH variable not set"
fi

make install
if [ $? -ne "0" ]
then
    echo "Could not install OpenCV"
fi

# observer
go get github.com/chutommy/observer
if [ $? -ne "0" ]
then
    echo "Could not get Observer"
fi
