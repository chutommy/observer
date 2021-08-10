package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"runtime"
	"sort"

	"github.com/blackjack/webcam"
)

const (
	V4L2_PIX_FMT_PJPG = 0x47504A50
	V4L2_PIX_FMT_YUYV = 0x56595559
)

type FrameSizes []webcam.FrameSize

func (slice FrameSizes) Len() int {
	return len(slice)
}

func (slice FrameSizes) Less(i, j int) bool {
	ls := slice[i].MaxWidth * slice[i].MaxHeight
	rs := slice[j].MaxWidth * slice[j].MaxHeight
	return ls < rs
}

func (slice FrameSizes) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

var supportedFormats = map[webcam.PixelFormat]bool{
	V4L2_PIX_FMT_PJPG: true,
	V4L2_PIX_FMT_YUYV: true,
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println(1)

	cam, err := webcam.Open("/dev/video0")
	if err != nil {
		panic(err.Error())
	}
	defer func(cam *webcam.Webcam) {
		_ = cam.Close()
	}(cam)

	formatDesc := cam.GetSupportedFormats()

	fmt.Println("Available formats:")
	for _, s := range formatDesc {
		_, _ = fmt.Fprintln(os.Stderr, s)
	}

	var format webcam.PixelFormat
FMT:
	for f := range formatDesc {
		if supportedFormats[f] {
			format = f
			break FMT
		}
	}
	if format == 0 {
		log.Println("No format found, exiting")
		return
	}

	frames := FrameSizes(cam.GetSupportedFrameSizes(format))
	sort.Sort(frames)

	_, _ = fmt.Fprintln(os.Stderr, "Supported frame sizes for format", formatDesc[format])
	for _, f := range frames {
		_, _ = fmt.Fprintln(os.Stderr, f.GetString())
	}
	var size *webcam.FrameSize
	size = &frames[len(frames)-1]
	if size == nil {
		log.Println("No matching frame size, exiting")
		return
	}

	_, _ = fmt.Fprintln(os.Stderr, "Requesting", formatDesc[format], size.GetString())
	f, w, h, err := cam.SetImageFormat(format, size.MaxWidth, size.MaxHeight)
	if err != nil {
		log.Println("SetImageFormat return error", err)
		return

	}
	_, _ = fmt.Fprintf(os.Stderr, "Resulting image format: %s %dx%d\n", formatDesc[f], w, h)

	err = cam.StartStreaming()
	if err != nil {
		log.Println(err)
		return
	}

	var (
		li   = make(chan *bytes.Buffer)
		fi   = make(chan []byte)
		back = make(chan struct{})
	)
	go encodeToImage(back, fi, li, w, h, f)

	timeout := uint32(5)

	for {
		err = cam.WaitForFrame(timeout)
		if err != nil {
			log.Println(err)
			return
		}

		switch err.(type) {
		case nil:
		case *webcam.Timeout:
			log.Println(err)
			continue
		default:
			log.Println(err)
			return
		}

		frame, err := cam.ReadFrame()
		if err != nil {
			log.Println(err)
			return
		}
		if len(frame) != 0 {
			select {
			case fi <- frame:
				fmt.Println(2)
				<-back
				fmt.Println(3)
			default:
			}
		}

		im, _, _ := image.Decode(<-li)
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
		os.Exit(1)
	}
}

func encodeToImage(back chan struct{}, fi chan []byte, li chan *bytes.Buffer, w, h uint32, format webcam.PixelFormat) {

	var (
		frame []byte
		img   image.Image
	)
	for {
		bframe := <-fi

		if len(frame) < len(bframe) {
			frame = make([]byte, len(bframe))
		}
		copy(frame, bframe)
		back <- struct{}{}

		switch format {
		case V4L2_PIX_FMT_YUYV:
			yuyv := image.NewYCbCr(image.Rect(0, 0, int(w), int(h)), image.YCbCrSubsampleRatio422)
			for i := range yuyv.Cb {
				ii := i * 4
				yuyv.Y[i*2] = frame[ii]
				yuyv.Y[i*2+1] = frame[ii+2]
				yuyv.Cb[i] = frame[ii+1]
				yuyv.Cr[i] = frame[ii+3]

			}
			img = yuyv
		default:
			log.Fatal("invalid format ?")
		}

		buf := &bytes.Buffer{}
		if err := jpeg.Encode(buf, img, nil); err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(4)
		li <- buf
		fmt.Println(5)
	}
}
