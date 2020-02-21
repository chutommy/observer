FROM ubuntu
RUN apt-get update -y
RUN apt-get install -y libv4l-dev

FROM denismakogon/gocv-alpine:4.0.1-buildstage
RUN go get -u -d gocv.io/x/gocv
RUN cd $GOPATH/src/gocv.io/x/gocv && go build -o $GOPATH/bin/gocv-version ./cmd/version/main.go
RUN go get -d -u gobot.io/x/gobot/...
RUN mkdir /app
ADD . /app
WORKDIR /app
