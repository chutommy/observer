# in docker
go  build -o app .

go run app -help


sudo docker build -t obsimage .

sudo docker run --device=/dev/video0 -it obsimage
