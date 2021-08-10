package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/blackjack/webcam"
)

func main() {
	cam, err := webcam.Open("/dev/video0")
	if err != nil {
		panic(err.Error())
	}
	defer func(cam *webcam.Webcam) {
		_ = cam.Close()
	}(cam)
	for code, format := range cam.GetSupportedFormats() {
		if format == "Motion-JPEG" {
			_, _, _, _ = cam.SetImageFormat(code, 1280, 720)
		}
	}

	err = cam.StartStreaming()
	if err != nil {
		panic(err.Error())
	}

	for {
		err = cam.WaitForFrame(50000)

		frame, err := cam.ReadFrame()
		if len(frame) != 0 {
			print(".")

			var img image.Image

			yuyv := image.NewYCbCr(image.Rect(0, 0, 1280, 720), image.YCbCrSubsampleRatio422)
			for i := range yuyv.Cb {
				ii := i * 4
				yuyv.Y[i*2] = frame[ii]
				yuyv.Y[i*2+1] = frame[ii+2]
				yuyv.Cb[i] = frame[ii+1]
				yuyv.Cr[i] = frame[ii+3]

			}
			img = yuyv
			buf := &bytes.Buffer{}
			if err := jpeg.Encode(buf, img, nil); err != nil {
				log.Fatal(err)
			}

			im, _, _ := image.Decode(buf)
			out, err := os.Create("./QRImg.png")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = png.Encode(out, im)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		} else if err != nil {
			panic(err.Error())
		}
	}
}
