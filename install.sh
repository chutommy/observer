#!/usr/bin/bash

# gobot
if ! go get -d -u gobot.io/x/gobot/...;
then
    echo "Could not get Gobot"
    exit 2
fi

# gocv and opencv
if ! go get -u -d gocv.io/x/gocv;
then
    echo "Could not get GoCV"
    exit 2
fi

if ! cd "$GOPATH/src/gocv.io/x/gocv";
then
    echo "GOPATH variable not set"
    exit 2
fi

if ! make install;
then
    echo "Could not install OpenCV"
    exit 2
fi

# observer
if ! go get github.com/chutommy/observer;
then
    echo "Could not get Observer"
    exit 2
fi
