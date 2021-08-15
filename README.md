# Observer

![eye logo](docs/project/logo.svg)

The Observer is a software for a face recognition and face tracking devices. The
application is able to fully and reliably control the movement of the camera
using two servomotors in any direction. The recognition algorithm is provided by
the real-time optimized computer vision library OpenCV.

The entire software was developed with a performance in mind and all decisions
were made to be perfectly compatible with low-end PCs (aiming to be runnable on
minicomputers like Raspberry Pis).

## Dependencies

The observer uses these technologies to run properly on all devices:

* [Go v1.14](https://golang.org/dl/)
* [Gobot v1.14.0](https://gobot.io/)
* [GoCV v0.23.0 ](https://gocv.io/)
* [OpenCV v4.3.0](https://opencv.org/)
* [HaarCascade](https://github.com/opencv/opencv/)

## Installation

Tested on Raspberry Pi 3 Model B+:

*With `go`, `opencv4`, `gocv` and `gobot` installed, run this command to install
executable file into your PATH:*

```bash
$ go get github.com/chutommy/observer-rpi
```

**OR**

You can also get `install.sh` file and run it in the project folder to install
all dependencies and the Observer software (can take up to 15 minutes).

```bash
$ ./install
```

Enter `./observer -h` to get more info about the execution command.

## Samples

![gif sample](docs/project/gifs/1.gif)
![gif sample](docs/project/gifs/2.gif)
![gif sample](docs/project/gifs/3.gif)
