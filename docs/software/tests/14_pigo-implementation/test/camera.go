package main

import (
	"fmt"
	"image"
	"log"
	"os"
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

func getCamForm(fd map[webcam.PixelFormat]string) webcam.PixelFormat {
	var format webcam.PixelFormat

	fmt.Println("Available formats:")
	for _, s := range fd {
		_, _ = fmt.Fprintln(os.Stderr, s)
	}

	for f := range fd {
		if supportedFormats[f] {
			format = f
			break
		}
	}
	if format == 0 {
		log.Fatalln("No format found, exiting")
	}

	return format
}

func getCamSize(fr FrameSizes, fd map[webcam.PixelFormat]string, fm webcam.PixelFormat) *webcam.FrameSize {
	sort.Sort(fr)

	_, _ = fmt.Fprintln(os.Stderr, "Supported frame sizes for format", fd[fm])
	for _, f := range fr {
		_, _ = fmt.Fprintln(os.Stderr, f.GetString())
	}
	var size *webcam.FrameSize
	size = &fr[len(fr)-1]
	if size == nil {
		log.Fatalln("No matching frame size, exiting")
	}

	return size
}

func encodeToImage(back chan struct{}, fi chan []byte, li chan *image.Image, w, h uint32, format webcam.PixelFormat) {

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

		li <- &img
	}
}
