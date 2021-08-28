---
weight: 3
title: "Installation"
draft: false
---

## Installation

{{< notification type="info" title="Info">}}
Tested on Raspberry Pi 3 Model B+
{{< /notification >}}

With `go`, `opencv4`, `gocv` and `gobot` installed, run this command to install
executable file into your **PATH**:

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
