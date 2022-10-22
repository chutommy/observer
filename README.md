# Observer

![eye logo](https://raw.githubusercontent.com/chutommy/observer/master/docs/project/logo.svg)

The Observer is a software for a face recognition and face tracking devices. The
application is able to fully and reliably control the movement of the camera
using two servomotors in any direction. The real-time recognition algorithm is
provided by the optimized computer vision library OpenCV.

The entire software was developed with a performance in mind and all decisions
were made to be perfectly compatible with low-end PCs (aiming to be runnable on
minicomputers like Raspberry Pis).

### Website

Check out the project website at [chutommy.com/observer](https://chutommy.com/observer/).

## Dependencies

The observer uses these technologies to run properly:

* [Go v1.14](https://golang.org/dl/)
* [Gobot v1.14.0](https://gobot.io/)
* [GoCV v0.23.0 ](https://gocv.io/)
* [OpenCV v4.3.0](https://opencv.org/)
* [HaarCascade](https://github.com/opencv/opencv/)

## Installation

Tested on Raspberry Pi 3 Model B+:

*With `go`, `opencv4`, `gocv` and `gobot` installed, run this command to install
executable file into your **PATH**:*

```bash
$ go get github.com/chutommy/observer-rpi
```

**OR**

Download the source and run the `install.sh` file as a sudo user in the project
root directory to install all dependencies and the Observer software (can take
up to 15 minutes).

```bash
$ sudo ./install
```

### Post-install steps

Type `./observer -h` or `--help` in a terminal to get more info about the
observer command.

## Project publication

The document which explains the goals of the Observer project. Everything
including the original inspiration and motive is thoroughly described and
discussed. All the questions/answers and critical decisions that were made
during the development process are included and analyzed as well.

The document also communicates my personal struggles, mistakes and doubts caused
by lack of knowledge, poor design choices and low budget for hardware
components. However, all obstacles and inaccuracies had been overcome and the
project was successfully finished.

**Czech
version:** [Dozorce - CS](https://docs.google.com/document/d/19FhBe4HvJeZMuhAFcyugnGlhiUhVPWTIfCu6MOH3Rxk/edit?usp=sharing)

## Gallery

**Full gallery:**
*See all photos and videos captured throughout the growth and development
process of the Observer project in
the [docs/](https://drive.google.com/drive/folders/1of6aFjSCA9LWL8vPtI93ILUU1oUt7Fva?usp=sharing) folder.*

### Samples

![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/project/gifs/1.gif)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/project/gifs/2.gif)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/project/gifs/3.gif)

### Skelet

#### Scheme

![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/skelet/schema/00.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/skelet/schema/02.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/skelet/schema/03.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/skelet/schema/07.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/skelet/schema/09.jpg)

#### Construction

![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/skelet/construction/01.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/skelet/construction/09.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/skelet/construction/11.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/skelet/construction/13.jpg)

### Cover

![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/cover/construction/01.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/cover/construction/04.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/cover/construction/09.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/cover/construction/11.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/cover/construction/12.jpg)

### Result

![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/result/images/00.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/result/images/01.jpg)
![](https://raw.githubusercontent.com/chutommy/observer/master/docs/compressed/result/images/02.jpg)

## License

The project is under the MIT open source software license.
