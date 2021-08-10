package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"

	"github.com/blackjack/webcam"
	pigo "github.com/esimov/pigo/core"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

var raspiAdaptor = raspi.NewAdaptor()

func main() {

	work := func() {

		cam, err := webcam.Open("/dev/video0")
		if err != nil {
			panic(err.Error())
		}
		defer func(cam *webcam.Webcam) {
			_ = cam.Close()
		}(cam)

		formatDesc := cam.GetSupportedFormats()
		format := getCamForm(formatDesc)

		frames := FrameSizes(cam.GetSupportedFrameSizes(format))
		size := getCamSize(frames, formatDesc, format)

		_, _ = fmt.Fprintln(os.Stderr, "Requesting", formatDesc[format], size.GetString())
		f, w, h, err := cam.SetImageFormat(format, size.MaxWidth, size.MaxHeight)
		if err != nil {
			log.Println("SetImageFormat return error", err)
			return
		}
		_, _ = fmt.Fprintf(os.Stderr, "Resulting image format: %s %dx%d\n", formatDesc[f], w, h)

		err = cam.StartStreaming()
		if err != nil {
			log.Printf("unalbe to start streaming: %v\n", err)
			return
		}

		var (
			li   = make(chan *image.Image)
			fi   = make(chan []byte)
			back = make(chan struct{})
		)
		go encodeToImage(back, fi, li, w, h, f)

		cascadeFile, err := ioutil.ReadFile(cascade)
		if err != nil {
			log.Fatalf("Error reading the cascade file: %v", err)
		}
		pigi := pigo.NewPigo()
		classifier, err := pigi.Unpack(cascadeFile)
		if err != nil {
			log.Fatalf("Error reading the cascade file: %s", err)
		}

		for {
			_ = cam.WaitForFrame(camTimeout)

			frame, err := cam.ReadFrame()
			if err != nil {
				log.Fatalf("unable to read frame: %v", err)
				return
			}
			if len(frame) != 0 {
				select {
				case fi <- frame:
					<-back
				default:
				}
			}

			func() {
				src := pigo.ImgToNRGBA(*(<-li))

				pixels := pigo.RgbToGrayscale(src)
				cols, rows := src.Bounds().Max.X, src.Bounds().Max.Y

				cParams := pigo.CascadeParams{
					MinSize:     20,
					MaxSize:     1000,
					ShiftFactor: 0.1,
					ScaleFactor: 1.1,

					ImageParams: pigo.ImageParams{
						Pixels: pixels,
						Rows:   rows,
						Cols:   cols,
						Dim:    cols,
					},
				}

				dets := classifier.RunCascade(cParams, 0)
				dets = classifier.ClusterDetections(dets, 0.2)

				fmt.Println(len(dets))

			}()
		}
	}

	connections := []gobot.Connection{raspiAdaptor}
	var devices []gobot.Device

	robot := gobot.NewRobot(
		"testRobot",
		connections,
		devices,
		work,
	)

	_ = robot.Start()
}
