# Observer

<p align="center">
  <img src="https://raw.githubusercontent.com/chutified/observer/master/img/00.jpg">
</p>
<br>

Observer is a software for a face recognition and face tracking. This software is a part of a six-month school project.

The application should be able to fully and reliably control the movement of the camera using two 180° servomotors in a half-space. For the recognition it uses the HaarCascades that are located in the `data` directory. All variables are editable (for more information, run the application with the `--help` flag).

The final work is available at https://drive.google.com/drive/folders/1of6aFjSCA9LWL8vPtI93ILUU1oUt7Fva?usp=sharing.

## Dependencies

The observer uses these dependencies to run properly:
  * Go v1.14 (https://golang.org/dl/)
  * Gobot v1.14.0 (https://gobot.io/)
  * GoCV v0.23.0 with OpenCV v4.3.0 (https://gocv.io/)
  * HaarCascade (https://github.com/opencv/opencv)
  
## Installation

On Raspberry Pi (CPU ARMv6+) with `go`, `opencv4`, `gocv` and `gobot` installed run this:

```bash
$ go get github.com/chutified/observer
```

Or get `install.sh` file and run this in the same location to install all depencies and the Observer software (can takes up to 15 minutes, depends on the performence):

```bash
$ ./install
```

Type `./observer -h` for the help.

*(tested on Raspbian OS with Raspberry Pi 3B+)*

## Samples

<p align="left">
  <img src="https://raw.githubusercontent.com/chutified/observer/master/img/05.gif">
</p>

<p align="left">
  <img src="https://raw.githubusercontent.com/chutified/observer/master/img/04.gif">
</p>

<p align="left">
  <img src="https://raw.githubusercontent.com/chutified/observer/master/img/03.gif">
</p>
<br>

To see more, click <a href="https://drive.google.com/drive/folders/1of6aFjSCA9LWL8vPtI93ILUU1oUt7Fva?usp=sharing" target="_blank">here</a>.
